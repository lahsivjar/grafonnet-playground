package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lahsivjar/grafonnet-playground/config"
	"github.com/lahsivjar/grafonnet-playground/handlers"
)

func main() {
	cfg := config.Load()

	router := setupGin()

	router.Use(
		static.Serve("/playground", static.LocalFile("./public", true)),
		static.Serve("/playground/dist", static.LocalFile("./dist", true)),
	)

	router.GET("/health", handlers.HealthCheckHandler)

	api := router.Group("/playground/api/v1")
	{
		api.POST("/run", handlers.RunHandler(cfg))
	}

	router.Run(":8080")
}

func setupGin() *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	return r
}
