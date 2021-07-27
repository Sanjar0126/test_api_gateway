package v1

import (
	"context"
	pbu "genproto/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// @Router /v1/shippers/{shipper_id} [get]
// @Summary Get Shipper
// @Description API for getting shipper info
// @Tags shipper
// @Accept  json
// @Produce json
// @Param shipper_id path string true "shipper_id"
// @Success 200 {object} models.GetShipperModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetShipper(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.ShipperService().Get(
		context.Background(), &pbu.ShipperId{
			Id: c.Param("shipper_id"),
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while getting shipper") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
