package v1

import (
	"context"
	"encoding/json"
	pb "genproto/post_service"
	"github.com/Sanjar0126/test_api_gateway/api/models"
	"github.com/Sanjar0126/test_api_gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"net/http"
	"strconv"
)

// @Tags post
// @Router /v1/posts/ [post]
// @Summary Create Post
// @Description API for create post
// @Accept  json
// @Produce json
// @Param post body models.CreatePostModel true "Create Post"
// @Success 200 {object} models.GetPostModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (handler *handlerV1) CreatePost(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		response models.GetPostModel
		post pb.Post
		)
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &post)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		handler.log.Error("error while unmarshal", logger.Error(err))
		return
	}

	resp, err := handler.grpcClient.PostService().Create(context.Background(), &post)
	if handleGrpcErrWithMessage(c, handler.log, err, "error while creating post"){
		return
	}

	js, err := jspbMarshal.MarshalToString(resp)
	if handleGrpcErrWithMessage(c, handler.log, err, "error while marshaling") {
		return
	}

	err = json.Unmarshal([]byte(js), &response)
	if handleBadRequestErrWithMessage(c, handler.log, err, "error while unmarshalling response") {
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Tags post
// @Router /v1/posts/{post_id} [get]
// @Summary Get Post
// @Description API for getting post info
// @Accept  json
// @Produce json
// @Param post_id path uint32 true "post_id"
// @Success 200 {object} models.GetPostModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (handler *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal jsonpb.Marshaler
	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	postId, _ := strconv.ParseUint(c.Param("post_id"), 10, 32)
	res, err := handler.grpcClient.PostService().Get(
		context.Background(), &pb.PostId{
			Id: uint32(postId),
		},
	)
	if handleGrpcErrWithMessage(c, handler.log, err, "error while getting post") {
		return
	}

	js, err := jspbMarshal.MarshalToString(res)
	if handleGrpcErrWithMessage(c, handler.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Security ApiKeyAuth
// @Router /v1/posts [get]
// @Summary Get All Posts
// @Description API for getting posts
// @Tags post
// @Accept  json
// @Produce  json
// @Param page query integer false "page"
// @Param limit query integer false "limit"
// @Success 200 {object} models.GetAllPostsModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (handler *handlerV1) GetAllPosts(c *gin.Context) {
	var (
		jspbMarshal jsonpb.Marshaler
	)

	jspbMarshal.OrigName = true
	jspbMarshal.EmitDefaults = true

	page, err := ParsePageQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: ErrorBadRequest,
		})
		return
	}

	res, err := handler.grpcClient.PostService().GetAll(
		context.Background(),
		&pb.GetAllPostsRequest{
			Page:      uint64(page),
			Limit:     uint64(limit),
		},
	)
	if handleGRPCErr(c, handler.log, err) {
		return
	}
	js, err := jspbMarshal.MarshalToString(res)

	if handleGrpcErrWithMessage(c, handler.log, err, "error while marshalling") {
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, js)
}

// @Tags post
// @Router /v1/posts/{post_id} [delete]
// @Summary Delete Post
// @Description API for deleting post object
// @Accept  json
// @Produce  json
// @Param post_id path uint32 true "post_id"
// @Success 200 {object} models.ResponseOK
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (handler *handlerV1) DeletePost(c *gin.Context) {

	postId, _ := strconv.ParseUint(c.Param("post_id"), 10, 32)
	_, err := handler.grpcClient.PostService().Delete(
		context.Background(),
		&pb.DeletePostRequest{
			Id: uint32(postId),
		},
	)

	if handleGrpcErrWithMessage(c, handler.log, err, "error while deleting post") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully updated",
	})
}


// UpdatePost
// @Router /v1/posts/{post_id} [put]
// @Summary Update Post
// @Description API for updating post
// @Tags post
// @Accept  json
// @Produce  json
// @Param post_id path uint32 true "post_id"
// @Param post body models.UpdatePostModel true "post"
// @Success 200 {object} models.GetPostModel
// @Failure 404 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		jspbMarshal   jsonpb.Marshaler
		jspbUnmarshal jsonpb.Unmarshaler
		post      pb.Post
	)
	jspbMarshal.OrigName = true

	err := jspbUnmarshal.Unmarshal(c.Request.Body, &post)
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

	postId, _ := strconv.ParseUint(c.Param("post_id"), 10, 32)
	post.Id = uint32(postId)

	_, err = h.grpcClient.PostService().Update(
		context.Background(), &post,
	)
	if handleGrpcErrWithMessage(c, h.log, err, "error while updating post") {
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Successfully updated",
	})
}
