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
// @Router /v1/review/{review_id} [get]
// @Summary Get review
// @Tags review
// @Description API for getting review
// @Accept  json
// @Produce json
// @Param review_id path string true "review_id"
// @Success 200 {object} models.Response
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetReview(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	jspbMarshal.OrigName = true

	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	review, err := h.grpcClient.ReviewService().Get(context.Background(), &pbc.ReviewId{Id: c.Param("review_id")})

	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting specific review") {
		return
	}

	js, err := jspbMarshal.MarshalToString(review)
	if handleInternalWithMessage(c, h.log, err, "Error while marshaling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)

}

// @Security ApiKeyAuth
// @Router /v1/review [get]
// @Summary Get All reviews
// @Tags review
// @Description API for getting review
// @Accept  json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetAllReviewsModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllReviews(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
		shipperID   string
	)
	if c.GetHeader("Authorization") == "" && c.GetHeader("Shipper") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: Authorization or shipper id have to be on the header")
		return
	}

	if c.GetHeader("Authorization") != "" {
		err := getUserInfo(h, c, &userInfo)

		if err != nil {
			return
		}
		shipperID = userInfo.ShipperID
	} else if c.GetHeader("Shipper") != "" {
		shipperID = c.GetHeader("Shipper")
	}

	jspbMarshal.OrigName = true

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: ErrorBadRequest})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: ErrorBadRequest})
		return
	}

	search := c.Query("search")

	reviews, err := h.grpcClient.ReviewService().GetAll(context.Background(), &pbc.GetAllReviewsRequest{
		Limit:     int64(limit),
		Page:      int64(page),
		Search:    search,
		ShipperId: shipperID,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting all reviews") {
		return
	}

	js, err := jspbMarshal.MarshalToString(reviews)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
