package v1

import (
	"context"
	"encoding/json"
	pb "genproto/catalog_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
)

// @Router /v1/product [get]
// @Summary Get All Product
// @Description API for getting all product
// @Tags product
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param category_id query string false "category_id"
// @Param search query string false "search"
// @Success 200 {object} models.GetAllProductsModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllProducts(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		model      models.GetAllProductsModel
	)

	shipperID := c.Request.Header.Get("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}
	marshaller.OrigName = true

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing page",
			},
		})
		h.log.Error("error while parsing page", logger.Error(err))
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing limit",
			},
		})
		h.log.Error("error while parsing limit", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.ProductService().GetAll(
		context.Background(),
		&pb.GetAllProductsRequest{
			ShipperId:     shipperID,
			Page:          int64(page),
			Limit:         int64(limit),
			CategoryId:    c.Query("category_id"),
			Search:        c.Query("search"),
			GetOnlyActive: true,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all products") {
		return
	}

	js, _ := marshaller.MarshalToString(resp)

	err = json.Unmarshal([]byte(js), &model)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "error while parsing proto to struct",
			},
		})
		h.log.Error("error while parsing proto to struct", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Router /v1/product/{product_id} [get]
// @Summary Get Product
// @Description API for getting a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Param product_id path string true "product_id"
// @Success 200 {object} models.GetProductModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetProduct(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
	)

	shipperID := c.Request.Header.Get("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}

	marshaller.OrigName = true
	marshaller.EmitDefaults = true

	resp, err := h.grpcClient.ProductService().Get(
		context.Background(),
		&pb.GetRequest{
			ShipperId:     shipperID,
			Id:            c.Param("product_id"),
			GetOnlyActive: true,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting a product") {
		return
	}

	js, err := marshaller.MarshalToString(resp)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "err while marshaling",
			},
		})
		h.log.Error("error while parsing proto to struct", logger.Error(err))
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
