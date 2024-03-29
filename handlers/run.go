package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-jsonnet"
	"github.com/lahsivjar/grafonnet-playground/config"
	"github.com/lahsivjar/grafonnet-playground/grafana"
	log "github.com/sirupsen/logrus"
)

type runRequest struct {
	Code string `json:"code" binding:"required"`
}

type runResponse struct {
	URL string `json:"url"`
}

// RunHandler handles the run endpoint which converts jsonnet to json and
// creates a grafana snapshot, returning it to the client
func RunHandler(cfg *config.Config, gSvc grafana.Service) func(*gin.Context) {
	return func(c *gin.Context) {
		var rReq runRequest
		if err := c.ShouldBindJSON(&rReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		j, err := getJsonnetVM(cfg.GrafonnetLibDirs).
			EvaluateSnippet("grafonnet-playground", rReq.Code)
		if err != nil {
			log.Error("Failed to evaluate jsonnet", err)
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		gReq := grafana.NewCreateRequest(cfg.GrafonnetPlaygroundFolderID)
		if err := json.Unmarshal([]byte(j), &gReq.Dashboard); err != nil {
			log.Error("Failed to create POST dashboard request", err)
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		gRes, err := gSvc.CreateDashboard(gReq)
		if err != nil {
			log.Error("Failed to create dashboard", err)
			c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
			return
		}

		c.JSON(http.StatusOK, runResponse{
			URL: cfg.GrafanaGetURL + gRes.URL,
		})
	}
}

func getJsonnetVM(jPaths []string) *jsonnet.VM {
	vm := jsonnet.MakeVM()
	i := &jsonnet.FileImporter{
		JPaths: jPaths,
	}
	vm.Importer(i)

	return vm
}
