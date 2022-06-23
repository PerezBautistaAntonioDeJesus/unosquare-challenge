package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"unosquare-challenge/api/models"
	"unosquare-challenge/gateways"
)

type KubernetesHandler struct {
	kubernetesGateWay gateways.KubernetesGateway
}

func NewKubernetesHandler( kb gateways.KubernetesGateway) *KubernetesHandler {
	return &KubernetesHandler{
		kubernetesGateWay: kb,
	}
}

func (kh *KubernetesHandler) GetServices(ctx *gin.Context) {


	services, err := kh.kubernetesGateWay.GetServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ServiceGatewayToServiceAPI(services))

}
