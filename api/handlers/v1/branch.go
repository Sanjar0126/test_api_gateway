package v1

import (
	"context"
	pbu "genproto/user_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
)

// @Security ApiKeyAuth
// @Tags branch
// @Router /v1/branches/{branch_id} [get]
// @Summary Get Branch
// @Description API for getting branch info
// @Accept  json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetBranch(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.BranchService().Get(
		context.Background(), &pbu.BranchId{
			Id: c.Param("branch_id"),
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while getting branch") {
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
// @Router /v1/branches [get]
// @Summary Get All Branches
// @Description API for getting branches
// @Tags branch
// @Accept  json
// @Produce  json
// @Param Shipper header string false "Shipper"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllBranchesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllBranches(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
		shipperID   string
	)

	if c.GetHeader("Authorization") == "" && c.GetHeader("Shipper") == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("shipper_id not found")
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

	res, err := h.grpcClient.BranchService().GetAll(
		context.Background(),
		&pbu.GetAllBranchesRequest{
			ShipperId: shipperID,
			Page:      uint64(page),
			Limit:     uint64(limit),
		},
	)
	if handleGRPCErr(c, h.log, err) {
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
// @Tags branch
// @Router /v1/nearest-branch [get]
// @Summary Get Nearest Branch
// @Description API for getting branch info
// @Accept  json
// @Produce json
// @Param long query string false "long"
// @Param lat query string false "lat"
// @Success 200 {object} models.GetBranchModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetNearestBranch(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		location    pbu.Location
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	longString, _ := c.GetQuery("long")
	latString, _ := c.GetQuery("lat")

	long, _ := strconv.ParseFloat(longString, 64)
	lat, _ := strconv.ParseFloat(latString, 64)

	location.Long = long
	location.Lat = lat

	res, err := h.grpcClient.BranchService().GetNearestBranch(
		context.Background(),
		&pbu.GetNearestBranchRequest{
			ShipperId: userInfo.ShipperID,
			Location:  &location,
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while getting branches") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
