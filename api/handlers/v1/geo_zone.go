package v1

import (
	"context"
	pbgeo "genproto/geo_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/spf13/cast"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
)

// @Security ApiKeyAuth
// @Router /v1/geozones/{geozone_id} [get]
// @Summary Get Geozone
// @Description API for getting geozone
// @Tags geozone
// @Accept  json
// @Produce  json
// @Param geozone_id path string true "geozone_id"
// @Success 200 {object} models.GetGeozoneModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetGeozone(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	res, err := h.grpcClient.GeoService().Get(
		context.Background(), &pbgeo.GeozoneId{
			Id: c.Param("geozone_id"),
		},
	)

	if handleGRPCErr(c, h.log, err) {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)

	if handleInternalWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/geozones [get]
// @Summary Get Geozones
// @Description API for getting geozones
// @Tags geozone
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllGeozonesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetAllGeozones(c *gin.Context) {
	var (
		userInfo    models.UserInfo
		jspbMarshal jsonpb.Marshaler
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)
	if handleBadRequestErrWithMessage(c, h.log, err, "error parsing page query parameter") {
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if handleBadRequestErrWithMessage(c, h.log, err, "error parsing limit query parameter") {
		return
	}

	res, err := h.grpcClient.GeoService().GetAll(
		context.Background(),
		&pbgeo.GetAllGeozonesRequest{
			ShipperId: userInfo.ShipperID,
			Page:      uint64(page),
			Limit:     uint64(limit),
		},
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error getting geozones") {
		return
	}
	js, err := jspbMarshal.MarshalToString(res)

	if handleInternalWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/geozones [post]
// @Summary Create Gezone
// @Description API for creating geozone
// @Tags geozone
// @Accept  json
// @Produce  json
// @Param geozone body models.CreateGeozoneModel true "geozone"
// @Success 200 {object} models.GetGeozoneModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateGeozone(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		geozone       pbgeo.Geozone
		userInfo      models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &geozone)
	if handleBadRequestErrWithMessage(c, h.log, err, "Error while unmarshalling") {
		return
	}

	geozone.ShipperId = userInfo.ShipperID

	resp, err := h.grpcClient.GeoService().Create(
		context.Background(), &geozone,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while creating geozone") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleInternalWithMessage(c, h.log, err, "Error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/geozones/{geozone_id} [put]
// @Summary Update Geozone
// @Description API for updating geozone
// @Tags geozone
// @Accept  json
// @Produce  json
// @Param geozone_id path string true "geozone_id"
// @Param geozone body models.UpdateGeozoneModel true "geozone"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateGeozone(c *gin.Context) {

	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		geozone       pbgeo.Geozone
		userInfo      models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	geozone.Id = c.Param("geozone_id")
	err = jspbUnmarshal.Unmarshal(c.Request.Body, &geozone)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while unmarshalling data") {
		return
	}

	_, err = h.grpcClient.GeoService().Update(
		context.Background(),
		&geozone,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "Error while updating geozone") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/geozones/{geozone_id} [delete]
// @Summary Delete Geozone
// @Description API for deleting geozone
// @Tags geozone
// @Accept  json
// @Produce  json
// @Param geozone_id path string true "geozone_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteGeozone(c *gin.Context) {
	var userInfo models.UserInfo
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	_, err = h.grpcClient.GeoService().Delete(
		context.Background(),
		&pbgeo.GeozoneId{
			Id: c.Param("geozone_id"),
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "Error while deleting geozone") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": "success",
	})
}

// @Security ApiKeyAuth
// @Router /v1/check-point-geozone [get]
// @Summary Check if point inside geozone
// @Description API for checking if point inside geozone
// @Tags geozone
// @Accept  json
// @Produce  json
// @Param lat query float32 true "lat"
// @Param lon query float32 true "lon"
// @Success 200 {object} models.CheckPointInsideGeozoneResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CheckPointInsideGeozone(c *gin.Context) {

	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	resp, err := h.grpcClient.GeoService().CheckPointGeozone(
		context.Background(),
		&pbgeo.CheckPointGeozoneRequest{
			ShipperId: userInfo.ShipperID,
			Point: &pbgeo.Point{
				Lat: cast.ToFloat32(c.Query("lat")),
				Lon: cast.ToFloat32(c.Query("lon")),
			},
		},
	)

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true
	if handleInternalWithMessage(c, h.log, err, "Error while checking point") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)

	if handleInternalWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}
