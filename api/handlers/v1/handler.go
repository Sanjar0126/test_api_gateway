package v1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	pba "genproto/auth_service"
	pbc "genproto/crm_service"
	pbo "genproto/order_service"
	pbp "genproto/payment_service"
	pbu "genproto/user_service"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/api/models"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/config"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/grpc_client"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
)

type handlerV1 struct {
	log        logger.Logger
	grpcClient *grpc_client.GrpcClient
	cfg        config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

const (
	// ErrorCodeInvalidURL ...
	ErrorCodeInvalidURL = "INVALID_URL"
	// ErrorCodeInvalidJSON ...
	ErrorCodeInvalidJSON = "INVALID_JSON"
	// ErrorCodeInternal ...
	ErrorCodeInternal = "INTERNAL"
	// ErrorCodeUnauthorized ...
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	// ErrorCodeAlreadyExists ...
	ErrorCodeAlreadyExists = "ALREADY_EXISTS"
	// ErrorCodeNotFound ...
	ErrorCodeNotFound = "NOT_FOUND"
	// ErrorCodeInvalidCode ...
	ErrorCodeInvalidCode = "INVALID_CODE"
	// ErrorBadRequest ...
	ErrorBadRequest = "BAD_REQUEST"
	// ErrorCodeForbidden ...
	ErrorCodeForbidden = "FORBIDDEN"
	// ErrorCodeNotApproved ...
	ErrorCodeNotApproved = "NOT_APPROVED"
	// ErrorCodeWrongClub ...
	ErrorCodeWrongClub = "WRONG_CLUB"
	// ErrorCodePasswordsNotEqual ...
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
)

var (
	signingKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
	SigningKey = []byte("FfLbN7pIEYe8@!EqrttOLiwa(H8)7Ddo")
)

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:        c.Logger,
		grpcClient: c.GrpcClient,
		cfg:        c.Cfg,
	}
}

// ParsePageQueryParam ...
func ParsePageQueryParam(c *gin.Context) (uint64, error) {
	page, err := strconv.ParseUint(c.DefaultQuery("page", "1"), 10, 10)
	if err != nil {
		return 0, err
	}
	if page < 0 {
		return 0, errors.New("page must be an positive integer")
	}
	if page == 0 {
		return 1, nil
	}
	return page, nil
}

// ParsePageSizeQueryParam ...
func ParsePageSizeQueryParam(c *gin.Context) (uint64, error) {
	pageSize, err := strconv.ParseUint(c.DefaultQuery("page_size", "10"), 10, 10)
	if err != nil {
		return 0, err
	}
	if pageSize < 0 {
		return 0, errors.New("page_size must be an positive integer")
	}
	return pageSize, nil
}

// ParseLimitQueryParam ...
func ParseLimitQueryParam(c *gin.Context) (uint64, error) {
	limit, err := strconv.ParseUint(c.DefaultQuery("limit", "10"), 10, 10)
	if err != nil {
		return 0, err
	}
	if limit < 0 {
		return 0, errors.New("page_size must be an positive integer")
	}
	if limit == 0 {
		return 10, nil
	}
	return limit, nil
}

// ParseAllQueryParam ...
func ParseAllQueryParam(c *gin.Context) (bool, error) {
	a, err := strconv.ParseBool(c.DefaultQuery("all", "false"))
	if err != nil {
		return false, err
	}
	return a, nil
}

func ParseBoolQueryParam(c *gin.Context, name string) (bool, error) {
	a, err := strconv.ParseBool(c.DefaultQuery(name, "false"))
	if err != nil {
		return false, err
	}
	return a, nil
}

func handleGRPCErr(c *gin.Context, l logger.Logger, err error) bool {
	if err == nil {
		return false
	}
	st, ok := status.FromError(err)
	var errI interface{} = models.InternalServerError{
		Code:    ErrorCodeInternal,
		Message: "Internal Server Error",
	}
	httpCode := http.StatusInternalServerError
	if ok && st.Code() == codes.InvalidArgument {
		httpCode = http.StatusBadRequest
		errI = ErrorBadRequest
	}
	c.JSON(httpCode, models.ResponseError{
		Error: errI,
	})
	if ok {
		l.Error(fmt.Sprintf("code=%d message=%s", st.Code(), st.Message()), logger.Error(err))
	}
	return true
}

func writeMessageAsJSON(c *gin.Context, l logger.Logger, msg proto.Message) {
	if msg == nil {
		c.String(http.StatusOK, "")
		return
	}
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	js, err := jspbMarshal.MarshalToString(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error("Error while marshaling", logger.Error(err))
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

func handleGrpcErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	st, ok := status.FromError(err)
	if !ok || st.Code() == codes.Internal {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: st.Message(),
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeNotFound,
				Message: st.Message(),
			},
		})
		l.Error(message+", not found", logger.Error(err))
		return true
	} else if st.Code() == codes.Unavailable {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message+", service unavailable", logger.Error(err))
		return true
	} else if st.Code() == codes.AlreadyExists {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeAlreadyExists,
				Message: st.Message(),
			},
		})
		l.Error(message+", already exists", logger.Error(err))
		return true
	} else if st.Code() == codes.InvalidArgument {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: st.Message(),
			},
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Code() == codes.Code(20) {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: st.Message(),
			},
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	} else if st.Err() != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorBadRequest,
				Message: st.Message(),
			},
		})
		l.Error(message+", invalid field", logger.Error(err))
		return true
	}
	return false
}

func handleInternalWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInternal,
				Message: "Internal Server Error",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func handleBadRequestErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: models.InternalServerError{
				Code:    ErrorCodeInvalidJSON,
				Message: "Invalid Json",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}

func getUserInfo(h *handlerV1, c *gin.Context, info *models.UserInfo) error {
	var (
		ErrUnauthorized = errors.New("Unauthorized")
		accessToken     string
	)

	accessToken = c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: ", logger.Error(ErrUnauthorized))
		return errors.New("Unauthorized request")
	}

	resp, err := h.grpcClient.AuthService().CheckPermission(context.Background(), &pba.CheckPermissionRequest{
		AccessToken: accessToken,
		UserTypeId:  config.CustomerTypeID,
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{
			Error: ErrorCodeUnauthorized,
		})
		h.log.Error("Unauthorized request: ", logger.Error(err))
		return errors.New("Unauthorized request")
	}

	if !resp.HasPermission {
		c.JSON(http.StatusForbidden, models.ResponseError{
			Error: ErrorCodeForbidden,
		})
		h.log.Error("Request Forbidden")
		return errors.New("Request Forbidden")
	}
	info.ID = resp.UserId
	info.UserID = resp.UserId
	info.ShipperID = resp.ShipperId

	return nil
}

func Contains(arr []interface{}, value interface{}) bool {
	if len(arr) == 0 {
		return false
	}

	if fmt.Sprintf("%T", arr[0]) != fmt.Sprintf("%T", value) {
		return false
	}

	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func (h *handlerV1) SendSms(model models.SendSmsPaymentsModel, userInfo models.UserInfo) (link string, err error) {
	var jspbMarshal jsonpb.Marshaler

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	if model.PaymentType != "click" && model.PaymentType != "payme" {
		return "", errors.New("wrong payment type")
	}

	shipper, err := h.grpcClient.ShipperService().Get(
		context.Background(),
		&pbu.ShipperId{
			Id: userInfo.ShipperID,
		})
	if err != nil {
		return "", err
	}

	order, err := h.grpcClient.OrderService().Get(
		context.Background(), &pbo.GetRequest{
			ShipperId: userInfo.ShipperID,
			Id:        model.OrderID,
		},
	)
	if err != nil {
		return "", err
	}

	orderNum := strconv.Itoa(int(order.ExternalOrderId))

	if shipper.Crm == "jowi" && order.JowiId.GetValue() != "" {

		res, err := h.grpcClient.CRMService().GetJowiBill(
			context.Background(),
			&pbc.GetJowiBillRequest{
				ShipperId:   userInfo.ShipperID,
				OrderJowiId: order.JowiId.Value,
			})
		if err != nil {
			h.log.Error("Error while getting jowi bill ID", logger.Error(err))
			return "", err
		}

		if model.PaymentType == "payme" {
			resPayme, err := h.grpcClient.PaymentService().GeneratePaymeLink(
				context.Background(),
				&pbp.GeneratePaymeLinkRequest{
					ShipperId:  userInfo.ShipperID,
					BranchId:   order.Steps[0].BranchId.GetValue(),
					Amount:     int64(order.DeliveryPrice + order.OrderAmount),
					ParamValue: res.BillId,
					ParamName:  "uid",
				})
			if err != nil {
				h.log.Error("Error while generating payme link", logger.Error(err))
				return "", err
			}

			link = resPayme.Link
		} else {
			resClick, err := h.grpcClient.PaymentService().GenerateClickLink(
				context.Background(),
				&pbp.GenerateClickLinkRequest{
					ShipperId: userInfo.ShipperID,
					BranchId:  order.Steps[0].BranchId.GetValue(),
					Amount:    int64(order.DeliveryPrice + order.OrderAmount),
					Value:     res.BillId,
				})
			if err != nil {
				h.log.Error("Error while generating click link", logger.Error(err))
				return "", err
			}

			link = resClick.Link
		}
	} else if model.PaymentType == "payme" {
		resPayme, err := h.grpcClient.PaymentService().GeneratePaymeLink(
			context.Background(),
			&pbp.GeneratePaymeLinkRequest{
				ShipperId:  userInfo.ShipperID,
				Amount:     int64(order.OrderAmount + order.CoDeliveryPrice),
				ParamName:  "order_num",
				ParamValue: orderNum,
				BranchId:   order.Steps[0].BranchId.GetValue(),
			},
		)
		if err != nil {
			h.log.Error("Error while generating payme link", logger.Error(err))
			return "", err
		}

		link = resPayme.Link
	} else {
		resClick, err := h.grpcClient.PaymentService().GenerateClickLink(
			context.Background(),
			&pbp.GenerateClickLinkRequest{
				ShipperId: userInfo.ShipperID,
				BranchId:  order.Steps[0].BranchId.GetValue(),
				Amount:    int64(order.OrderAmount + order.CoDeliveryPrice),
				Value:     orderNum,
			})
		if err != nil {
			h.log.Error("Error while generating click link", logger.Error(err))
			return "", err
		}

		link = resClick.Link
	}

	return link, nil
}
