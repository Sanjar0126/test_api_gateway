package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Sanjar0126/test_api_gateway/api/docs" //for swagger
	v1 "github.com/Sanjar0126/test_api_gateway/api/handlers/v1"
	"github.com/Sanjar0126/test_api_gateway/config"
	"github.com/Sanjar0126/test_api_gateway/pkg/grpc_client"

	"github.com/Sanjar0126/test_api_gateway/pkg/logger"
)

//Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// New
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

	//POST ENDPOINTS
	r.GET("/v1/posts/:post_id", handlerV1.GetPost)
	r.DELETE("/v1/posts/:post_id", handlerV1.DeletePost)
	r.PUT("/v1/posts/:post_id", handlerV1.UpdatePost)
	r.GET("/v1/posts/", handlerV1.GetAllPosts)
	r.POST("/v1/posts/", handlerV1.CreatePost)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
