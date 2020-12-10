package app

import (
	"os"

	"github.com/PubApi/cmd/api/app/infrastructure/controller/middleware"
	"github.com/PubApi/pkg/logger"
	"github.com/PubApi/pkg/mlhandlers"
)

func StartApp() {
	router := mlhandlers.DefaultRouter()
	router.Use(middleware.ErrorHandler())

	MapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8080"
	}

	if err := router.Run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}
