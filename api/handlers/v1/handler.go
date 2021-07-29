package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Sanjar0126/test_api_gateway/api/models"
	"github.com/Sanjar0126/test_api_gateway/config"
	"github.com/Sanjar0126/test_api_gateway/pkg/grpc_client"
	"github.com/Sanjar0126/test_api_gateway/pkg/logger"
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
