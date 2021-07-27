package v1

import (
	"context"
	"net/http"

	pb "genproto/fare_service"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// @Router /v1/delivery-price [get]
// @Summary Get Delivery Price
// @Description API for getting delivery price for shipper
// @Tags fare
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Success 200 {object} models.DeliveryPriceModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetDeliveryPrice(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
	)

	jspbMarshal.OrigName = true

	shipperID := c.Request.Header.Get("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.FareService().GetDeliveryPrice(
		context.Background(), &pb.GetDeliveryPriceRequest{
			ShipperId: shipperID,
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while delivery price") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/fares/compute-price [patch]
// @Summary Compute Delivery Price
// @Description API for computing delivery price
// @Tags fare
// @Accept  json
// @Produce  json
// @Param price body models.ComputeDeliveryPriceRequestModel true "price"
// @Success 200 {object} models.ComputeDeliveryPriceResponseModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ComputeDeliveryPrice(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		price         pb.ComputeDeliveryPriceRequest
		userInfo      models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &price)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while computing delivery price") {
		return
	}

	res, err := h.grpcClient.FareService().ComputeDeliveryPrice(
		context.Background(), &price,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while delivery price") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
