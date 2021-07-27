package v1

import (
	"context"
	pbc "genproto/content_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
)

// @Security ApiKeyAuth
// @Router /v1/user_reviews [post]
// @Summary Create UserReviews
// @Description API for creating UserReviews
// @Tags UserReviews
// @Accept  json
// @Produce  json
// @Param userReviews body models.CreateUserReviewsModel true "userReviews"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateUserReviews(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnMarshal jsonpb.Unmarshaler
		userReviews   pbc.UserReviews
		userInfo      models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	err = jspbUnMarshal.Unmarshal(c.Request.Body, &userReviews)

	if handleInternalWithMessage(c, h.log, err, "Error while unmarshaling request.body") {
		return
	}

	userReviews.ShipperId = userInfo.ShipperID
	res, err := h.grpcClient.UserReviewsService().Create(context.Background(), &userReviews)

	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating UserReviews") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleInternalWithMessage(c, h.log, err, "Error while marshaling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)

}
