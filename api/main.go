package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gitlab.udevs.io/delever/delever_customer_api_gateway/api/docs" //for swagger
	v1 "gitlab.udevs.io/delever/delever_customer_api_gateway/api/handlers/v1"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/config"
	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/grpc_client"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/pkg/logger"
)

//Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	r.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	// Customer endpoints
	r.PUT("/v1/customers/:customer_id", handlerV1.UpdateCustomer)
	r.DELETE("/v1/customers/:customer_id", handlerV1.DeleteCustomer)
	r.GET("/v1/customer-profile", handlerV1.ProfileInformation)
	r.GET("/v1/customers/:customer_id", handlerV1.GetCustomer)
	r.POST("/v1/customers/phone", handlerV1.GetCustomerByPhone)
	r.POST("/v1/customers/login", handlerV1.CustomerLogin)
	r.POST("/v1/customers/confirm-login", handlerV1.ConfirmCustomerLogin)
	r.POST("/v1/customers/register", handlerV1.CustomerRegister)
	r.PATCH("/v1/customers/update-tg-chat-id", handlerV1.UpdateTgChatId)
	r.POST("/v1/customers/register-confirm", handlerV1.CustomerRegisterConfirm)
	r.POST("/v1/customers/refresh-token", handlerV1.CustomerRefreshToken)
	r.GET("/v1/customers/:customer_id/orders", handlerV1.GetCustomerOrders)
	r.GET("/v1/customers/:customer_id/last-order", handlerV1.GetCustomerLastOrder)
	r.GET("/v1/customer-addresses/:phone", handlerV1.GetCustomerAddresses)
	r.PUT("/v1/customer/bot-language", handlerV1.AttachBotLanguage)

	// Product endpoints
	r.GET("/v1/product", handlerV1.GetAllProducts)
	r.GET("/v1/product/:product_id", handlerV1.GetProduct)

	// Branch endpoints
	r.GET("/v1/branches/:branch_id", handlerV1.GetBranch)
	r.GET("/v1/branches", handlerV1.GetAllBranches)
	r.GET("/v1/nearest-branch", handlerV1.GetNearestBranch)

	// Category endpoints
	r.GET("/v1/category/:category_id", handlerV1.GetCategory)
	r.GET("/v1/category", handlerV1.GetAllCategory)

	// Fare endpoints
	r.GET("/v1/delivery-price", handlerV1.GetDeliveryPrice)
	r.PATCH("/v1/fares/compute-price", handlerV1.ComputeDeliveryPrice)

	// Order enpoints
	r.PATCH("/v1/order/:order_id/review", handlerV1.CreateReview)
	r.POST("/v1/ondemand-order", handlerV1.CreateOnDemandOrder)
	r.GET("/v1/order/:order_id", handlerV1.GetOrder)
	r.GET("/v1/order-statuses", handlerV1.GetStatuses)
	r.GET("/v1/order", handlerV1.GetCusomerOrders)
	r.PUT("/v1/order-review/:order_id", handlerV1.UpdateOrderReviewSeen)

	// Shipper endpoints
	r.GET("/v1/shippers/:shipper_id", handlerV1.GetShipper)

	// Payment endpoints
	r.POST("/v1/payment/payme-link", handlerV1.GeneratePaymeLink)
	r.POST("/v1/payment/click-link", handlerV1.GenerateClickLink)

	// Promo endpoints
	r.GET("/v1/promo/:promo_id", handlerV1.GetPromo)
	r.GET("/v1/promo", handlerV1.GetAllPromos)

	// User reviews endpoints
	r.POST("/v1/user_reviews", handlerV1.CreateUserReviews)

	// Banner endpoints
	r.GET("/v1/banner/:banner_id", handlerV1.GetBanner)
	r.GET("/v1/banner", handlerV1.GetAllBanners)

	// TelegramPush endpoints
	r.POST("/v1/telegram", handlerV1.PushTelegramShipperUsers)

	// Review endpoints
	r.GET("v1/review/:review_id", handlerV1.GetReview)
	r.GET("v1/review", handlerV1.GetAllReviews)

	// Geozone endpoints
	r.POST("/v1/geozones", handlerV1.CreateGeozone)
	r.GET("/v1/check-point-geozone", handlerV1.CheckPointInsideGeozone)
	r.PUT("/v1/geozones/:geozone_id", handlerV1.UpdateGeozone)
	r.DELETE("/v1/geozones/:geozone_id", handlerV1.DeleteGeozone)
	r.GET("/v1/geozones/:geozone_id", handlerV1.GetGeozone)
	r.GET("/v1/geozones", handlerV1.GetAllGeozones)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
