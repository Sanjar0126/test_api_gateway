package v1

import (
	"context"
	"encoding/json"
	pba "genproto/auth_service"
	pbu "genproto/user_service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/helpers"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/config"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
)

// @Security ApiKeyAuth
// @Tags customer
// @Router /v1/customers/{customer_id} [get]
// @Summary Get Customer
// @Description API for getting customer info
// @Accept  json
// @Produce json
// @Param customer_id path string true "customer_id"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomer(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		userInfo    models.UserInfo
		customer    *pbu.Customer
		err         error
		model       models.GetCustomerModel
	)
	err = getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	customer, err = h.grpcClient.CustomerService().Get(
		context.Background(), &pbu.CustomerId{
			Id: c.Param("customer_id"),
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer") {
		return
	}

	js, err := jspbMarshal.MarshalToString(customer)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, model)
}

// @Tags customer
// @Router /v1/customers/phone [post]
// @Summary Get Customer By Phone
// @Description API for getting customer by phone
// @Accept  json
// @Produce json
// @Param Shipper header string true "Shipper"
// @Param login body models.CustomerLoginRequest true "login"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetCustomerByPhone(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		customer      *pbu.Customer
		model         models.GetCustomerModel
		customerPhone pbu.GetCustomerByPhoneRequest
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

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &customerPhone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while unmarshalling data", logger.Error(err))
		return
	}

	customer, err = h.grpcClient.CustomerService().GetByPhone(
		context.Background(), &pbu.GetCustomerByPhoneRequest{
			Phone:     customerPhone.Phone,
			ShipperId: shipperID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer") {
		return
	}

	js, err := jspbMarshal.MarshalToString(customer)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	err = json.Unmarshal([]byte(js), &model)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, model)
}

// @Security ApiKeyAuth
// @Router /v1/customers/{customer_id} [put]
// @Summary Update Customer
// @Description API for updating customer
// @Tags customer
// @Accept  json
// @Produce  json
// @Param customer_id path string true "customer_id"
// @Param customer body models.UpdateCustomerModel true "customer"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateCustomer(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		customer      pbu.Customer
		userInfo      models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while unmarshalling data", logger.Error(err))
		return
	}

	customer.ShipperId = userInfo.ShipperID
	customer.Id = c.Param("customer_id")
	result, _ := h.grpcClient.CustomerService().GetByPhone(
		context.Background(), &pbu.GetCustomerByPhoneRequest{
			ShipperId: userInfo.ShipperID,
			Phone:     customer.Phone,
		},
	)

	if result != nil && result.Id != customer.Id {
		c.JSON(http.StatusConflict, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: "Phone already exists",
			},
		})
		h.log.Error("Error while checking phone, Already exists", logger.Error(err))
		return
	}

	_, err = h.grpcClient.CustomerService().Update(
		context.Background(), &customer,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while updating customer") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully updated",
	})
}

// @Security ApiKeyAuth
// @Tags customer
// @Router /v1/customers/{customer_id} [delete]
// @Summary Delete Customer
// @Description API for deleting customer
// @Accept  json
// @Produce  json
// @Param customer_id path string true "customer_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteCustomer(c *gin.Context) {
	var (
		userInfo models.UserInfo
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	_, err = h.grpcClient.CustomerService().Delete(
		context.Background(),
		&pbu.DeleteCustomerRequest{
			ShipperId: userInfo.ShipperID,
			Id:        c.Param("customer_id"),
		},
	)

	if handleGrpcErrWithMessage(c, h.log, err, "error while deleting customer") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully updated",
	})
}

// @Router /v1/customers/login [POST]
// @Summary Customer Login
// @Description API that checks whether customer exists
// @Description and if exists sends sms to their number
// @Tags customer
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Param login body models.CustomerLoginRequest true "login"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CustomerLogin(c *gin.Context) {
	var (
		login     models.CustomerLoginRequest
		shipperID string
	)

	shipperID = c.Request.Header.Get("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&login)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	_, err = h.grpcClient.AuthService().CustomerLogin(context.Background(), &pba.OTPLoginRequest{
		Phone:     login.Phone,
		ShipperId: shipperID,
		Tag:       login.Tag,
	})
	if handleGrpcErrWithMessage(c, h.log, err, "error while logging in") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "Code has been sent",
	})
}

// @Router /v1/customers/confirm-login [POST]
// @Summary Confirm Customer Login
// @Description API that confirms sms code
// @Tags customer
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Param Platform header string true "Platform"
// @Param confirm_phone body models.ConfirmCustomerLoginRequest true "confirm login"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ConfirmCustomerLogin(c *gin.Context) {
	var (
		cm          models.ConfirmCustomerLoginRequest
		jspbMarshal jsonpb.Marshaler
	)

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	shipperID := c.Request.Header.Get("Shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	platformID := c.Request.Header.Get("Platform")
	if platformID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "platform not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&cm)
	if handleBadRequestErrWithMessage(c, h.log, err, "error while binding to json") {
		return
	}

	resp, err := h.grpcClient.AuthService().CustomerConfirmLogin(context.Background(), &pba.OTPConfirmRequest{
		ShipperId:   shipperID,
		PlatformId:  platformID,
		Phone:       cm.Phone,
		Code:        cm.Code,
		FcmToken:    cm.FcmToken,
		TgChatId:    cm.TgChatId,
		BotLanguage: cm.BotLanguage,
	})

	if handleGrpcErrWithMessage(c, h.log, err, "error while logging in") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Summary Register
// @Description Register - API for registering customers
// @Tags customer
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Param register body models.RegisterModel true "register"
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/customers/register [post]
func (h *handlerV1) CustomerRegister(c *gin.Context) {
	var (
		reg       models.RegisterModel
		shipperID string
	)

	shipperID = c.Request.Header.Get("Shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&reg)
	if handleBadRequestErrWithMessage(c, h.log, err, "Error while binding json") {
		return
	}

	err = helpers.ValidatePhone(reg.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	_, err = h.grpcClient.AuthService().CustomerRegister(context.Background(), &pba.OTPRegisterRequest{
		Phone:     reg.Phone,
		Name:      reg.Name,
		ShipperId: shipperID,
	})
	if handleGrpcErrWithMessage(c, h.log, err, "error while registering customer") {
		return
	}

	c.JSON(200, models.ResponseOK{
		Message: "Code has been sent",
	})
}

// @Router /v1/customers/register-confirm [post]
// @Summary Register confirm
// @Description Register - API for confirming and inserting user to DB
// @Tags customer
// @Accept  json
// @Produce  json
// @Param Shipper header string true "Shipper"
// @Param Platform header string true "Platform"
// @Param register_confirm body models.RegisterConfirmModel true "register_confirm"
// @Success 200 {object} models.GetCustomerModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CustomerRegisterConfirm(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		reg         models.RegisterConfirmModel
		shipperID   string
	)

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	shipperID = c.Request.Header.Get("shipper")
	if shipperID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "shipper not found",
			"code":    ErrorBadRequest,
		})
		return
	}

	platformID := c.Request.Header.Get("platform")
	if platformID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "platform not found in header",
			"code":    ErrorBadRequest,
		})
		return
	}

	err := c.ShouldBindJSON(&reg)
	if handleBadRequestErrWithMessage(c, h.log, err, "Error binding json") {
		return
	}

	resp, err := h.grpcClient.AuthService().CustomerConfirmRegister(context.Background(), &pba.OTPConfirmRequest{
		ShipperId:   shipperID,
		PlatformId:  platformID,
		Phone:       reg.Phone,
		Code:        reg.Code,
		FcmToken:    reg.FcmToken,
		TgChatId:    reg.TgChatId,
		BotLanguage: reg.BotLanguage,
	})
	if handleGrpcErrWithMessage(c, h.log, err, "error while logging in") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/customers/refresh-token [post]
// @Summary Customer Refresh Token
// @Description Customer refresh token
// @Tags customer
// @Accept  json
// @Produce  json
// @Param refresh_token body models.RefreshTokenRequest true "refresh-token"
// @Success 200 {object} models.RefreshTokenResponse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CustomerRefreshToken(c *gin.Context) {
	var (
		refreshTokenRequest models.RefreshTokenRequest
		jspbMarshal         jsonpb.Marshaler
	)

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err := c.ShouldBindJSON(&refreshTokenRequest)
	if err != nil {
		h.log.Error("error while binding refresh token request parameters", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "refresh token field is required ",
			"code":    ErrorBadRequest,
		})
		return
	}

	resp, err := h.grpcClient.AuthService().RefreshToken(context.Background(), &pba.RefreshTokenRequest{
		RefreshToken: refreshTokenRequest.RefreshToken,
		UserTypeId:   config.CustomerTypeID,
	})
	if handleGrpcErrWithMessage(c, h.log, err, "error while logging in") {
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/customer-profile [get]
// @Summary Customer Profile Information
// @Description Customer profile information
// @Tags customer
// @Accept  json
// @Produce  json
// @Success 200 {object} models.GetCustomerModel
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) ProfileInformation(c *gin.Context) {
	var (
		userInfo    models.UserInfo
		jspbMarshal jsonpb.Marshaler
		customer    *pbu.Customer
	)

	err := getUserInfo(h, c, &userInfo)
	if err != nil {
		return
	}

	customer, err = h.grpcClient.CustomerService().Get(
		context.Background(), &pbu.CustomerId{
			Id: userInfo.ID,
		})

	if handleGrpcErrWithMessage(c, h.log, err, "error while getting customer") {
		return
	}

	js, err := jspbMarshal.MarshalToString(customer)

	if handleGrpcErrWithMessage(c, h.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/customers/update-tg-chat-id [patch]
// @Summary Tg chat id
// @Description Tg chat id
// @Tags tg chat id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ResponseOK
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateTgChatId(c *gin.Context) {
	var (
		userInfo      models.UserInfo
		jspbUnMarshal jsonpb.Unmarshaler
		tgChatRequest pbu.UpdateCustomerTgChatIdRequest
	)

	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	err = jspbUnMarshal.Unmarshal(c.Request.Body, &tgChatRequest)
	if handleInternalWithMessage(c, h.log, err, "error while unmarshalling") {
		return
	}

	tgChatRequest.ShipperId = userInfo.ShipperID
	_, err = h.grpcClient.CustomerService().UpdateTgChatId(context.Background(), &tgChatRequest)
	if handleGrpcErrWithMessage(c, h.log, err, "error while updating tg chat id") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{Message: "Tg chat id started to update"})
}

// @Security ApiKeyAuth
// @Router /v1/customer/bot-language [put]
// @Summary AttachBotLanguage
// @Description API for updating customer
// @Tags customer
// @Accept  json
// @Produce  json
// @Param lang body models.BotLanguageModel true "lang"
// @Success 200 {object} models.GetCustomerModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) AttachBotLanguage(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		userInfo      models.UserInfo
		languageModel pbu.AttachBotLanguageRequest
	)
	err := getUserInfo(h, c, &userInfo)

	if err != nil {
		return
	}

	jspbMarshal.OrigName = true

	err = jspbUnmarshal.Unmarshal(c.Request.Body, &languageModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server error",
			},
		})
		h.log.Error("Error while unmarshalling data", logger.Error(err))
		return
	}

	_, err = h.grpcClient.CustomerService().AttachBotLanguage(
		context.Background(), &pbu.AttachBotLanguageRequest{
			Id:   userInfo.ID,
			Lang: languageModel.Lang,
		},
	)
	h.log.Error("Error while attaching language", logger.Error(err))
	return

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully attached language",
	})
}
