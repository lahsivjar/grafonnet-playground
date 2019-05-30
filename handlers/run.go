package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/go-jsonnet"
	"github.com/lahsivjar/grafonnet-playground/config"
)

type runRequest struct {
	Code string `json:"code" binding:"required"`
}

type runResponse struct {
	URL string `json:"url"`
}

type grafanaReq struct {
	Dashboard interface{} `json:"dashboard"`
	Overwrite bool        `json:"overwrite"`
	FolderID  int         `json:"folderId,omitempty"`
}

type grafanaRes struct {
	ID      int    `json:"id"`
	UID     string `json:"uid"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	Version int    `json:"version"`
}

func grafanaReqWithDefaults(folderID int) *grafanaReq {
	return &grafanaReq{
		Overwrite: true,
		FolderID:  folderID,
	}
}

// RunHandler handles the run endpoint which converts jsonnet to json and
// creates a grafana snapshot, returning it to the client
func RunHandler(cfg *config.Config) func(*gin.Context) {
	return func(c *gin.Context) {
		var rReq runRequest
		if err := c.ShouldBindJSON(&rReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		j, err := getJsonnetVM(cfg.GrafonnetLibDir).
			EvaluateSnippet("grafonnet-playground", rReq.Code)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}
		gReq := grafanaReqWithDefaults(cfg.GrafonnetPlaygroundFolderID)
		if err := json.Unmarshal([]byte(j), &gReq.Dashboard); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		gRes, err := createDashboard(gReq, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		c.JSON(http.StatusOK, runResponse{
			URL: cfg.GrafanaURL + gRes.URL,
		})
	}
}

func getJsonnetVM(jPath string) *jsonnet.VM {
	vm := jsonnet.MakeVM()
	fmt.Println(jPath)
	i := &jsonnet.FileImporter{
		JPaths: []string{jPath},
	}
	vm.Importer(i)

	return vm
}

func createDashboard(g *grafanaReq, cfg *config.Config) (*grafanaRes, error) {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")

	reqBody, err := getRequestBody(g)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", cfg.GrafanaURL+"/api/dashboards/db", reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+cfg.GrafanaAPIKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		var gRes grafanaRes
		err = json.NewDecoder(resp.Body).Decode(&gRes)
		if err != nil {
			return nil, err
		}

		return &gRes, nil
	}

	errorMsg, err := ioutil.ReadAll(resp.Body)
	return nil, fmt.Errorf("Error occurred while creating graph: %s", errorMsg)
}

func getRequestBody(g *grafanaReq) (io.Reader, error) {
	b, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(b)), nil
}
