package config

import (
	"os"

	"github.com/spf13/cast"
)

const (
	CustomerTypeID        = "9b31286d-dc21-4572-b20c-9b6cb7d29d89"
	TelegramBotPlatformId = "e8f87c9e-fbb4-4572-be4d-c959e0cc2c2f"

	NewStatusId               = "986a0d09-7b4d-4ca9-8567-aa1c6d770505"
	OperatorAcceptedStatusId  = "ccb62ffb-f0e1-472e-bf32-d130bea90617"
	OperatorCancelledStatusId = "b5d1aa93-bccd-40bb-ae29-ea5a85a2b1d1"
	VendorAcceptedStatusId    = "1b6dc9a3-64aa-4f68-b54f-71ffe8164cd3"
	VendorReadyStatusId       = "b0cb7c69-5e3d-47c7-9813-b0a7cc3d81fd"
	CourierAcceptedStatusId   = "8781af8e-f74d-4fb6-ae23-fd997f4a2ee0"
	CourierPickedUpStatusId   = "84be5a2f-3a92-4469-8283-220ca34a0de4"
	DeliveredStatusId         = "79413606-a56f-45ed-97c3-f3f18e645972"
	FinishedStatusId          = "e665273d-5415-4243-a329-aee410e39465"
	ServerCancelledStatusId   = "d39cb255-6cf5-4602-896d-9c559d40cbbe"
)

var PaymentTypes = []interface{}{"cash", "payme", "click", "transfer"}

var OrderSources = []interface{}{"admin_panel", "ios", "android", "website", "bot"}

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost string
	UserServicePort int

	OrderServiceHost string
	OrderServicePort int

	SmsServiceHost string
	SmsServicePort int

	CatalogServiceHost string
	CatalogServicePort int

	FareServiceHost string
	FareServicePort int

	AuthServiceHost string
	AuthServicePort int

	PaymentServiceHost string
	PaymentServicePort int

	ContentServiceHost string
	ContentServicePort int

	ConsumerServiceHost string
	ConsumerServicePort int

	LogLevel string
	HTTPPort string

	MinioEndpoint       string
	MinioAccessKeyID    string
	MinioSecretAccesKey string

	GeoServiceHost string
	GeoServicePort int

	CRMServiceHost string
	CRMServicePort int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "prod"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":1239"))

	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "user-service"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 80))

	c.OrderServiceHost = cast.ToString(getOrReturnDefault("ORDER_SERVICE_HOST", "order-service"))
	c.OrderServicePort = cast.ToInt(getOrReturnDefault("ORDER_SERVICE_PORT", 80))

	c.SmsServiceHost = cast.ToString(getOrReturnDefault("SMS_SERVICE_HOST", "sms-service"))
	c.SmsServicePort = cast.ToInt(getOrReturnDefault("SMS_SERVICE_PORT", 80))

	c.FareServiceHost = cast.ToString(getOrReturnDefault("FARE_SERVICE_HOST", "fare-service"))
	c.FareServicePort = cast.ToInt(getOrReturnDefault("FARE_SERVICE_PORT", 80))

	c.CatalogServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "catalog-service"))
	c.CatalogServicePort = cast.ToInt(getOrReturnDefault("CATALOG_SERVICE_PORT", 80))

	c.AuthServiceHost = cast.ToString(getOrReturnDefault("AUTH_SERVICE_HOST", "auth-service"))
	c.AuthServicePort = cast.ToInt(getOrReturnDefault("AUTH_SERVICE_PORT", 80))

	c.PaymentServiceHost = cast.ToString(getOrReturnDefault("PAYMENT_SERVICE_HOST", "payment-service"))
	c.PaymentServicePort = cast.ToInt(getOrReturnDefault("PAYMENT_SERVICE_PORT", 80))

	c.ContentServiceHost = cast.ToString(getOrReturnDefault("CONTENT_SERVICE_HOST", "content-service"))
	c.ContentServicePort = cast.ToInt(getOrReturnDefault("CONTENT_SERVICE_PORT", 80))

	c.ConsumerServiceHost = cast.ToString(getOrReturnDefault("CONSUMER_SERVICE_HOST", "consumer-service"))
	c.ConsumerServicePort = cast.ToInt(getOrReturnDefault("CONSUMER_SERVICE_PORT", 80))

	c.GeoServiceHost = cast.ToString(getOrReturnDefault("GEO_SERVICE_HOST", "geo-service"))
	c.GeoServicePort = cast.ToInt(getOrReturnDefault("GEO_SERVICE_PORT", 80))

	c.MinioEndpoint = cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "cdn.test.delever.uz"))
	c.MinioAccessKeyID = cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "e963c4649b849ed4fa6180a1a80beb456f91d0bc964357dc46a7ffe948554e62"))
	c.MinioSecretAccesKey = cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY_ID", "c8b28792d5ddb42aa04fe50e4349fc822c3a8b2401647952f20517efa19c7c16"))

	c.GeoServiceHost = cast.ToString(getOrReturnDefault("GEO_SERVICE_HOST", "geo-service"))
	c.GeoServicePort = cast.ToInt(getOrReturnDefault("GEO_SERVICE_PORT", 80))

	c.CRMServiceHost = cast.ToString(getOrReturnDefault("CRM_SERVICE_HOST", "crm-service"))
	c.CRMServicePort = cast.ToInt(getOrReturnDefault("CRM_SERVICE_PORT", 80))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
