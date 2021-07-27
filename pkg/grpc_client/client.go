package grpc_client

import (
	"fmt"
	pba "genproto/auth_service"
	pbc "genproto/catalog_service"
	pbcon "genproto/consumer_service"
	pbco "genproto/content_service"
	pbcrm "genproto/crm_service"
	pbf "genproto/fare_service"
	pbgeo "genproto/geo_service"
	pbo "genproto/order_service"
	pbp "genproto/payment_service"
	pbs "genproto/sms_service"
	pbu "genproto/user_service"

	"google.golang.org/grpc"

	"gitlab.udevs.io/delever/delever_customer_api_gateway/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	OrderService() pbo.OrderServiceClient
	SmsService() pbs.SmsServiceClient
	CustomerService() pbu.CustomerServiceClient
	BranchService() pbu.BranchServiceClient
	CategoryService() pbc.CategoryServiceClient
	ProductService() pbc.ProductServiceClient
	FareService() pbf.FareServiceClient
	AuthService() pba.AuthServiceClient
	ShipperService() pbu.ShipperServiceClient
	PromoService() pbco.PromoServiceClient
	BannerService() pbco.BannerServiceClient
	PushService() pbcon.PushServiceClient
	ReviewService() pbco.ReviewServiceClient
	UserReviewsService() pbco.UserReviewsServiceClient
	GeoService() pbgeo.GeoServiceClient
	CRMService() pbcrm.CRMServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connOrder, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("order service dial host: %s port:%d err: %s",
			cfg.OrderServiceHost, cfg.OrderServicePort, err)
	}

	connSms, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.SmsServiceHost, cfg.SmsServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("sms service dial host: %s port:%d err: %s",
			cfg.SmsServiceHost, cfg.SmsServicePort, err)
	}

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d err: %s",
			cfg.UserServiceHost, cfg.UserServicePort, err)
	}

	connCatalog, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CatalogServiceHost, cfg.CatalogServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("catalog service dial host: %s port:%d err: %s",
			cfg.CatalogServiceHost, cfg.CatalogServicePort, err)
	}

	connAuth, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.AuthServiceHost, cfg.AuthServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %s port:%d err: %s",
			cfg.AuthServiceHost, cfg.AuthServicePort, err)
	}

	connFare, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.FareServiceHost, cfg.FareServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("fare service dial host: %s port:%d err: %s",
			cfg.FareServiceHost, cfg.FareServicePort, err)
	}

	connPayment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PaymentServiceHost, cfg.PaymentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Payment service dial host: %s post:%d err: %s",
			cfg.PaymentServiceHost, cfg.PaymentServicePort, err)
	}

	connContent, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ContentServiceHost, cfg.ContentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Content service dial host: %s post:%d err: %s",
			cfg.ContentServiceHost, cfg.ContentServicePort, err)
	}

	connConsumer, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ConsumerServiceHost, cfg.ConsumerServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Consumer service dial host: %s post:%d err: %s",
			cfg.ContentServiceHost, cfg.ContentServicePort, err)
	}

	connGeo, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.GeoServiceHost, cfg.GeoServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Geor service dial host: %s post:%d err: %s",
			cfg.GeoServiceHost, cfg.GeoServicePort, err)
	}

	connCrm, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CRMServiceHost, cfg.CRMServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("CRM service dial host: %s post:%d err: %s",
			cfg.CRMServiceHost, cfg.CRMServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"order_service":        pbo.NewOrderServiceClient(connOrder),
			"sms_service":          pbs.NewSmsServiceClient(connSms),
			"customer_service":     pbu.NewCustomerServiceClient(connUser),
			"fare_service":         pbf.NewFareServiceClient(connFare),
			"category_service":     pbc.NewCategoryServiceClient(connCatalog),
			"branch_service":       pbu.NewBranchServiceClient(connUser),
			"product_service":      pbc.NewProductServiceClient(connCatalog),
			"auth_service":         pba.NewAuthServiceClient(connAuth),
			"shipper_service":      pbu.NewShipperServiceClient(connUser),
			"payment_service":      pbp.NewPaymentServiceClient(connPayment),
			"promo_service":        pbco.NewPromoServiceClient(connContent),
			"banner_service":       pbco.NewBannerServiceClient(connContent),
			"push_service":         pbcon.NewPushServiceClient(connConsumer),
			"review_service":       pbco.NewReviewServiceClient(connContent),
			"user_reviews_service": pbco.NewUserReviewsServiceClient(connContent),
			"geo_service":          pbgeo.NewGeoServiceClient(connGeo),
			"crm_service":          pbcrm.NewCRMServiceClient(connCrm),
		},
	}, nil
}

// CustomerService ...
func (g *GrpcClient) CustomerService() pbu.CustomerServiceClient {
	return g.connections["customer_service"].(pbu.CustomerServiceClient)
}

// OrderService ...
func (g *GrpcClient) OrderService() pbo.OrderServiceClient {
	return g.connections["order_service"].(pbo.OrderServiceClient)
}

// SmsService ...
func (g *GrpcClient) SmsService() pbs.SmsServiceClient {
	return g.connections["sms_service"].(pbs.SmsServiceClient)
}

func (g *GrpcClient) FareService() pbf.FareServiceClient {
	return g.connections["fare_service"].(pbf.FareServiceClient)
}

func (g *GrpcClient) BranchService() pbu.BranchServiceClient {
	return g.connections["branch_service"].(pbu.BranchServiceClient)
}

func (g *GrpcClient) CategoryService() pbc.CategoryServiceClient {
	return g.connections["category_service"].(pbc.CategoryServiceClient)
}

func (g *GrpcClient) ProductService() pbc.ProductServiceClient {
	return g.connections["product_service"].(pbc.ProductServiceClient)
}

func (g *GrpcClient) AuthService() pba.AuthServiceClient {
	return g.connections["auth_service"].(pba.AuthServiceClient)
}

func (g *GrpcClient) ShipperService() pbu.ShipperServiceClient {
	return g.connections["shipper_service"].(pbu.ShipperServiceClient)
}

func (g *GrpcClient) PaymentService() pbp.PaymentServiceClient {
	return g.connections["payment_service"].(pbp.PaymentServiceClient)
}

func (g *GrpcClient) PromoService() pbco.PromoServiceClient {
	return g.connections["promo_service"].(pbco.PromoServiceClient)
}

func (g *GrpcClient) BannerService() pbco.BannerServiceClient {
	return g.connections["banner_service"].(pbco.BannerServiceClient)
}

func (g *GrpcClient) PushService() pbcon.PushServiceClient {
	return g.connections["push_service"].(pbcon.PushServiceClient)
}

func (g *GrpcClient) ReviewService() pbco.ReviewServiceClient {
	return g.connections["review_service"].(pbco.ReviewServiceClient)
}

func (g *GrpcClient) UserReviewsService() pbco.UserReviewsServiceClient {
	return g.connections["user_reviews_service"].(pbco.UserReviewsServiceClient)
}

func (g *GrpcClient) GeoService() pbgeo.GeoServiceClient {
	return g.connections["geo_service"].(pbgeo.GeoServiceClient)
}

func (g *GrpcClient) CRMService() pbcrm.CRMServiceClient {
	return g.connections["crm_service"].(pbcrm.CRMServiceClient)
}
