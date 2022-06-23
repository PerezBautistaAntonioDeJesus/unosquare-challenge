// Unosquare challenge API specification
//
// The nex documentation provides you the information for the API endpoints
//
//
// Terms Of Service:
//
// the use of the API is the responsability of the user
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//     License: MIT no
//     Contact: Antonio Perez<antonio.perba@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
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
