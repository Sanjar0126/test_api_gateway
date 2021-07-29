package grpc_client

import (
	"fmt"
	pbpost "genproto/post_service"
	"google.golang.org/grpc"

	"github.com/Sanjar0126/test_api_gateway/config"
)

// GrpcClientI ...
type GrpcClientI interface {
	PostService() pbpost.PostServiceClient
}

// GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

// New ...
func New(cfg config.Config) (*GrpcClient, error) {
	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("Post service dial host: %s post:%d err: %s",
			cfg.PostServiceHost, cfg.PostServicePort, err)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"post_service": pbpost.NewPostServiceClient(connPost),
		},
	}, nil
}

func (g *GrpcClient) PostService() pbpost.PostServiceClient {
	return g.connections["post_service"].(pbpost.PostServiceClient)
}
