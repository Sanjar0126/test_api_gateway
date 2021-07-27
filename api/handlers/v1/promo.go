package v1

import (
	"context"
	"net/http"

	pbc "genproto/content_service"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// @Router /v1/promo/{promo_id} [get]
// @Summary Get promo
// @Tags promo
// @Description API for getting promo
// @Accept  json
// @Produce json
// @Param promo_id path string true "promo_id"
// @Success 200 {object} models.GetPromoModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetPromo(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
	)

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	promo, err := h.grpcClient.PromoService().Get(
		context.Background(), &pbc.PromoId{
			Id: c.Param("promo_id"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting promo") {
		return
	}

	js, err := jspbMarshal.MarshalToString(promo)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/promo [get]
// @Summary Get All promos
// @Tags promo
// @Description API for getting promo
// @Accept  json
// @Produce json
// @Param Shipper header string true "Shipper"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetAllPromosModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllPromos(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
	)

	shipperID := c.Request.Header.Get("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := h.grpcClient.PromoService().GetAll(
		context.Background(), &pbc.GetAllPromosRequest{
			ShipperId: shipperID,
			Page:      int64(page),
			Limit:     int64(limit),
			Search:    c.Query("search"),
		})
	if handleGrpcErrWithMessage(c, h.log, err, "error while getting promo") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
