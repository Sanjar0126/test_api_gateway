package v1

import (
	"context"
	pbp "genproto/payment_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
)

// @Security ApiKeyAuth
// @Router /v1/payment/payme-link [post]
// @Summary Generate payme link
// @Description API for generating payme link
// @Tags payment
// @Accept  json
// @Produce  json
// @Param generate-link body models.GenerateLinkRequestModel true "generate-link"
// @Success 200 {object} models.GenerateLinkResponseModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GeneratePaymeLink(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		userInfo      models.UserInfo
		linkRequest   pbp.GeneratePaymeLinkRequest
	)

	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &linkRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.PaymentService().GeneratePaymeLink(
		context.Background(),
		&pbp.GeneratePaymeLinkRequest{
			ShipperId: userInfo.ShipperID,
			Amount:    int64(linkRequest.Amount),
			ParamValue: linkRequest.ParamValue,
			ParamName:  "order_num",
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while generating payme link") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshaling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/payment/click-link [post]
// @Summary Generate click link
// @Description API for generating click link
// @Tags payment
// @Accept  json
// @Produce  json
// @Param generate-link body models.GenerateLinkRequestModel true "generate-link"
// @Success 200 {object} models.GenerateLinkResponseModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GenerateClickLink(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		userInfo      models.UserInfo
		linkRequest   pbp.GenerateClickLinkRequest
	)

	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &linkRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.PaymentService().GenerateClickLink(
		context.Background(),
		&pbp.GenerateClickLinkRequest{
			ShipperId: userInfo.ShipperID,
			Amount:    int64(linkRequest.Amount),
			Value:     linkRequest.Value,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while generating click link") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshaling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
