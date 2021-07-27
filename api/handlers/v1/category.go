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

// @Security ApiKeyAuth
// @Router /v1/category [get]
// @Summary Get All Category
// @Description API for getting all category
// @Tags category
// @Accept  json
// @Produce  json
// @Param Shipper header string false "Shipper"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param parent_id query integer false "parent_id"
// @Param search query string false "search"
// @Param all query bool false "all"
// @Param with_products query bool false "with_products"
// @Success 200 {object} models.GetAllCategoriesModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllCategory(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		userInfo   models.UserInfo
		shipperID  string
		model      models.GetAllCategoriesModel
	)

	marshaller.OrigName = true
	marshaller.EmitDefaults = true

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

	marshaller.OrigName = true
	marshaller.EmitDefaults = true
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

	all, err := ParseAllQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing all query",
			},
		})
		h.log.Error("error while parsing all query", logger.Error(err))
		return
	}

	withProducts, err := ParseBoolQueryParam(c, "with_products")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: "error while parsing bool query",
			},
		})
		h.log.Error("error while parsing bool query", logger.Error(err))
		return
	}

	resp, err := h.grpcClient.CategoryService().GetAll(
		context.Background(),
		&pb.GetAllRequest{
			ShipperId:     shipperID,
			Page:          int64(page),
			Limit:         int64(limit),
			ParentId:      c.Query("parent_id"),
			All:           bool(all),
			Search:        c.Query("search"),
			WithProducts:  withProducts,
			GetOnlyActive: true,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all categories") {
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

	if handleInternalWithMessage(c, h.log, err, "Error while marshaling to string") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/category/{category_id} [get]
// @Summary Get Category
// @Description API for getting a category
// @Tags category
// @Accept  json
// @Produce  json
// @Param Shipper header string false "Shipper"
// @Param category_id path string true "category_id"
// @Success 200 {object} models.GetCategoryModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCategory(c *gin.Context) {
	var (
		marshaller jsonpb.Marshaler
		userInfo   models.UserInfo
		shipperID  string
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

	marshaller.OrigName = true
	marshaller.EmitDefaults = true

	resp, err := h.grpcClient.CategoryService().Get(
		context.Background(),
		&pb.GetRequest{
			ShipperId: shipperID,
			Id:        c.Param("category_id"),

			GetOnlyActive: true,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting a category") {
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
