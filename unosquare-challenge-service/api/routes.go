package api

import (
	"unosquare-challenge/api/handlers"
	"unosquare-challenge/gateways"

	"github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) *gin.RouterGroup {
	subgroup := engine.Group("cluster")

	kubernetesGateway := gateways.NewKubernetesGateway()
	kubHandler := handlers.NewKubernetesHandler(kubernetesGateway)

	subgroup.GET("services", kubHandler.GetServices)

	return subgroup

}
