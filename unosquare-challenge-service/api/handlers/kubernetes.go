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

func NewKubernetesHandler(kb gateways.KubernetesGateway) *KubernetesHandler {
	return &KubernetesHandler{
		kubernetesGateWay: kb,
	}
}

// swagger:operation GET /cluster/services services list-services
//
// list services running on kubernetes
//
// ---
// produces:
// - application/json
//
//
// responses:
//  '200':
//    description: list of services
//    schema:
//     type: array
//     items:
//      "$ref": "#/definitions/service"
//  '500':
//    description: internal error
//    schema:
//     "$ref": "#/definitions/errorResponse"
func (kh *KubernetesHandler) GetServices(ctx *gin.Context) {

	services, err := kh.kubernetesGateWay.GetServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, models.ServiceGatewayToServiceAPI(services))

}
