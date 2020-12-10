package app

import (
	"fmt"

	"github.com/PubApi/cmd/api/app/infrastructure/config"
	"github.com/PubApi/cmd/api/app/infrastructure/container"
	controller "github.com/PubApi/cmd/api/app/infrastructure/controller/health"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	prefixScope := config.GetRoutePrefix()
	router.GET("/ping", controller.PingController.Ping)
	prefix := fmt.Sprintf("%s/PubApi/", prefixScope)

	baseUrl := router.Group(prefix)

	beer := baseUrl.Group("/Beer")

	beer.POST("", container.GetCreateBeerController().MakeCreateBeer)
}
