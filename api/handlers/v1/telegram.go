package v1

import (
	"context"
	pbc "genproto/consumer_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
)

// @Security ApiKeyAuth
// @Router /v1/telegram [post]
// @Summary Push Telegram shipper users
// @Description API for pushing to telegram users
// @Tags telegram
// @Accept  json
// @Produce  json
// @Param pushtelegram body models.PushTelegramModel true "pushtelegram"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) PushTelegramShipperUsers(c *gin.Context) {
	var (
		jspbUnMarshal jsonpb.Unmarshaler
		userInfo      models.UserInfo
		pushTelegram  pbc.PushTelegramShipperUsersRequest
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	err = jspbUnMarshal.Unmarshal(c.Request.Body, &pushTelegram)
	if handleInternalWithMessage(c, h.log, err, "error while unmarshalling") {
		return
	}

	pushTelegram.ShipperId = userInfo.ShipperID
	_, err = h.grpcClient.PushService().PushTelegramShipperUsers(context.Background(), &pushTelegram)
	if handleGrpcErrWithMessage(c, h.log, err, "error while pushing") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Push started successfully",
	})
}
