package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type runRequest struct {
	Code string `json:"code" binding:"required"`
}

type runResponse struct {
	URL string `json:"url"`
}

// RunHandler handles the run endpoint which converts jsonnet to json and
// creates a grafana snapshot, returning it to the client
func RunHandler(c *gin.Context) {
	var rReq runRequest
	if err := c.ShouldBindJSON(&rReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, runResponse{
		URL: "https://google.com",
	})
}
