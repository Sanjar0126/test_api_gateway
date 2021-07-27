package v1

import (
	"context"
	pbc "genproto/content_service"
	"net/http"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
)

// @Router /v1/banner/{banner_id} [get]
// @Summary Get banner
// @Tags banner
// @Description API for getting banner
// @Accept json
// @Produce json
// @Param banner_id path string true "banner_id"
// @Success 200 {object} models.GetBannerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetBanner(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	banner, err := h.grpcClient.BannerService().Get(context.Background(), &pbc.BannerId{Id: c.Param("banner_id")})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting banner") {
		return
	}

	js, err := jspbMarshal.MarshalToString(banner)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Router /v1/banner [get]
// @Summary Get All banners
// @Tags banner
// @Description API for getting all banners
// @Accept json
// @Produce json
// @Param Shipper header string true "Shipper"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} models.GetAllPromosModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllBanners(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

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

	res, err := h.grpcClient.BannerService().GetAll(context.Background(), &pbc.GetAllBannersRequest{
		ShipperId: shipperID,
		Page:      int64(page),
		Limit:     int64(limit),
		Search:    c.Query("search"),
	})
	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all banners") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
