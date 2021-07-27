package v1

import (
	"context"
	"encoding/json"
	"fmt"
	pbo "genproto/order_service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/config"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
)

// @Security ApiKeyAuth
// @Router /v1/customer-addresses/{phone} [get]
// @Summary Get Customer Order Addresses
// @Description API for taking all order addresses
// @Tags customer
// @Accept  json
// @Produce  json
// @Param phone path string true "phone"
// @Param search query string false "search"
// @Success 200 {object} models.CustomerAddressesModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomerAddresses(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	phone := c.Param("phone")

	res, err := h.grpcClient.OrderService().GetCustomerAddresses(
		context.Background(),
		&pbo.GetCustomerAddressesRequest{
			ShipperId: userInfo.ShipperID,
			Phone:     phone,
			Search:    c.Query("search"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer addresses") {
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
// @Router /v1/customers/{customer_id}/orders [get]
// @Summary Get Customer Orders
// @Description API for getting customer orders
// @Tags customer
// @Accept json
// @Produce json
// @Param customer_id path string true "customer_id"
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param status_ids query []string false "status_ids"
// @Param start_date query string false "start_date"
// @Param end_date query string false "end_date"
// @Param review_seen query string false "review_seen"
// @Success 200 {object} models.GetAllOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomerOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		order       *pbo.OrdersResponse
		model       models.GetAllOrderModel
		userInfo    models.UserInfo
		statusIDs   []string
		reviewSeen  bool
	)
	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	customerID := c.Param("customer_id")
	if c.Query("review_seen") == "false" {
		reviewSeen = true
	}

	jspbMarshal.OrigName = true

	page, err := ParsePageQueryParam(c)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	if len(c.QueryArray("status_ids")) > 0 {
		statusIDs = strings.Split(c.QueryArray("status_ids")[0], ",")
	}

	order, err = h.grpcClient.OrderService().GetAll(context.Background(),
		&pbo.OrdersRequest{
			ShipperId:  userInfo.ShipperID,
			CustomerId: customerID,
			Page:       page,
			Limit:      limit,
			StatusIds:  statusIDs,
			StartDate:  c.Query("start_date"),
			EndDate:    c.Query("end_date"),
			ReviewSeen: reviewSeen,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting all customer order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id}/review [patch]
// @Summary Create Review For An Order
// @Description API for creating review for order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Param order body models.OrderReview true "order_review"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateReview(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
		userInfo      models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	orderID := c.Param("order_id")

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &order)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	order.Id = orderID
	_, err = h.grpcClient.OrderService().CreateReview(context.Background(), &order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order review") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "review created successfully",
	})
}

// @Security ApiKeyAuth
// @Router /v1/ondemand-order [post]
// @Summary Create On Demand Order
// @Description API for creating on demand order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body models.CreateOnDemandOrderModel true "order"
// @Success 200 {object} models.CreateOnDemandOrderModelResponse
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateOnDemandOrder(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		order         pbo.Order
		userInfo      models.UserInfo
		response      models.CreateOnDemandOrderModelResponse
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &order)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	if order.ToLocation == nil || order.ToLocation.Lat == 0 || order.ToLocation.Long == 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("Location is not valid", logger.Error(err))
		return
	}

	if !Contains(config.PaymentTypes, order.PaymentType) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("payment type is not valid", logger.Error(err))
		return
	}

	if order.Source != "admin_panel" && order.Source != "website" &&
		order.Source != "bot" && order.Source != "android" && order.Source != "ios" {

		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("source type is not valid", logger.Error(err))
		return
	}

	if order.DeliveryType == "" {
		order.DeliveryType = "delivery"
	} else if order.DeliveryType != "delivery" && order.DeliveryType != "self-pickup" {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		h.log.Error("delivery type is not valid", logger.Error(err))
		return
	}

	if order.StatusId != config.OperatorAcceptedStatusId {
		order.StatusId = config.NewStatusId
	}

	order.DeliveryPrice = order.CoDeliveryPrice
	order.ShipperId = userInfo.ShipperID
	order.FareId = "b35436da-a347-4794-a9dd-1dcbf918b35d"
	order.CreatorTypeId = config.CustomerTypeID

	resp, err := h.grpcClient.OrderService().Create(context.Background(), &order)
	if handleGrpcErrWithMessage(c, h.log, err, "error while creating order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshaling") {
		return
	}

	err = json.Unmarshal([]byte(js), &response)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while unmarshalling response") {
		return
	}

	if order.PaymentType == "payme" || order.PaymentType == "click" {
		link, err := h.SendSms(models.SendSmsPaymentsModel{
			Phone:       order.ClientPhoneNumber,
			OrderID:     resp.OrderId,
			PaymentType: order.PaymentType,
		}, userInfo)

		if err == nil {
			response.PaymentLink = link
		}
	}

	// c.Header("Content-Type", "application/json")
	// c.String(http.StatusOK, js)
	c.JSON(http.StatusOK, response)
}

// @Security ApiKeyAuth
// @Router /v1/order/{order_id} [get]
// @Summary Get Order
// @Description API for getting order
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Success 200 {object} models.GetOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetOrder(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	orderID := c.Param("order_id")

	order, err := h.grpcClient.OrderService().Get(context.Background(), &pbo.GetRequest{
		ShipperId: userInfo.ShipperID,
		Id:        orderID,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/customers/{customer_id}/last-order [get]
// @Summary Get Customer Last Order
// @Description API for getting customer last order
// @Tags customer
// @Accept  json
// @Produce  json
// @Param customer_id path string true "customer_id"
// @Success 200 {object} models.GetOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomerLastOrder(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	order, err := h.grpcClient.OrderService().GetCustomerLastOrder(context.Background(), &pbo.GetCustomerLastOrderRequest{
		ShipperId:  userInfo.ShipperID,
		CustomerId: c.Param("customer_id"),
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer last order") {
		return
	}

	js, err := jspbMarshal.MarshalToString(order)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/order-statuses [get]
// @Summary Get All Possible Order Statuses
// @Description API for getting order statuses
// @Tags order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetStatuses
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetStatuses(c *gin.Context) {
	var ()
	m := make(map[string]string)
	m["new"] = config.NewStatusId
	m["operator_accepted"] = config.OperatorAcceptedStatusId
	m["operator_cancelled"] = config.OperatorCancelledStatusId
	m["vendor_accepted"] = config.VendorAcceptedStatusId
	m["vendor_ready"] = config.VendorReadyStatusId
	m["courier_accepted"] = config.CourierAcceptedStatusId
	m["courier_picked_up"] = config.CourierPickedUpStatusId
	m["delivered"] = config.DeliveredStatusId
	m["finished"] = config.FinishedStatusId
	m["server_cancelled"] = config.ServerCancelledStatusId

	c.JSON(http.StatusOK, m)
}

// @Security ApiKeyAuth
// @Router /v1/order [get]
// @Summary Get Customer Orders
// @Description API for getting customer orders
// @Tags order
// @Accept json
// @Produce json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Param status_ids query []string false "status_ids"
// @Param start_date query string false "start_date"
// @Param end_date query string false "end_date"
// @Param review_seen query string false "review_seen"
// @Success 200 {object} models.GetAllOrderModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCusomerOrders(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		model       models.GetAllOrderModel
		userInfo    models.UserInfo
		statusIDs   []string
		reviewSeen  bool
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	if c.Query("review_seen") == "false" {
		reviewSeen = true
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing page") {
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while parsing limit") {
		return
	}

	if len(c.QueryArray("status_ids")) > 0 {
		statusIDs = strings.Split(c.QueryArray("status_ids")[0], ",")
	}

	h.log.Info(fmt.Sprintf("user: %v", userInfo), logger.Any("user", userInfo))

	orders, err := h.grpcClient.OrderService().GetAll(
		context.Background(),
		&pbo.OrdersRequest{
			ShipperId:  userInfo.ShipperID,
			CustomerId: userInfo.UserID,
			StatusIds:  statusIDs,
			Page:       page,
			Limit:      limit,
			StartDate:  c.Query("start_date"),
			EndDate:    c.Query("end_date"),
			SortBy:     "-created_at",
			ReviewSeen: reviewSeen,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer orders") {
		return
	}

	js, err := jspbMarshal.MarshalToString(orders)
	if handleInternalWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)
	if handleInternalWithMessage(c, h.log, err, "error while unmarshal to json") {
		return
	}

	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/order-review/{order_id} [put]
// @Summary Update Order Review Seen
// @Description API for updating order review seen
// @Tags order
// @Accept  json
// @Produce  json
// @Param order_id path string true "order_id"
// @Param order_review body models.UpdateOrderReview true "order_review"
// @Success 200 {object} models.UpdateOrderReview
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateOrderReviewSeen(c *gin.Context) {
	var (
		seen     models.UpdateOrderReview
		userInfo models.UserInfo
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	orderID := c.Param("order_id")

	err = c.BindJSON(&seen)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Error while binding json! ",
		})
		return
	}

	_, err = h.grpcClient.OrderService().UpdateOrderReviewSeen(
		context.Background(),
		&pbo.UpdateReviewSeen{
			OrderId:    orderID,
			ShipperId:  userInfo.ShipperID,
			CustomerId: seen.UserID,
		},
	)

	if handleInternalWithMessage(c, h.log, err, "error while updating order review") {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Successfully Updated",
	})
}
