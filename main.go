package main

import (
	"context"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lahsivjar/grafonnet-playground/config"
	"github.com/lahsivjar/grafonnet-playground/grafana"
	"github.com/lahsivjar/grafonnet-playground/handlers"
)

func main() {
	cfg := config.Load()
	grafanaService := grafana.NewService(cfg)

	if cfg.AutoCleanup {
		err := grafanaService.SetupCleanerJob(context.Background())
		if err != nil {
			panic(err)
		}
	}

	router := setupGin()

	router.Use(
		static.Serve("/playground", static.LocalFile("./public", true)),
		static.Serve("/playground/dist", static.LocalFile("./dist", true)),
	)

	router.GET("/health", handlers.HealthCheckHandler)

	api := router.Group("/playground/api/v1")
	{
		api.POST("/run", handlers.RunHandler(cfg, grafanaService))
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
