package handler

import (
	"net/http"
	"post-api/apperror"
	"post-api/dto"
	"post-api/service"

	"github.com/gin-gonic/gin"
)

const postHandlerFile = "post_handler.go"

type PostHandler struct {
	postSvc service.PostService
}

func NewPostHandler(ps service.PostService) *PostHandler {
	return &PostHandler{ps}
}

func (h *PostHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, err := h.postSvc.GetAll(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(posts))
	}

}

func (h *PostHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := GetPathID(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		post, err := h.postSvc.GetByID(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(post))
	}
}

func (h *PostHandler) Insert() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		newPost := new(dto.NewPostRequest)
		err := ctx.ShouldBindJSON(newPost)
		if err != nil {
			_ = ctx.Error(apperror.NewError(
				err,
				postHandlerFile,
				"PostHandler.Insert()",
				"ctx.ShouldBindJSON()",
			))
			return
		}

		createdPost, err := h.postSvc.Insert(ctx, newPost)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusCreated, dto.NewSuccessResponse(createdPost))
	}
}

func (h *PostHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		updatePost := new(dto.UpdatePostRequest)
		err := ctx.ShouldBindJSON(updatePost)
		if err != nil {
			_ = ctx.Error(apperror.NewError(
				err,
				postHandlerFile,
				"PostHandler.Update()",
				"ctx.ShouldBindJSON()",
			))
			return
		}

		id, err := GetPathID(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		updatePost.ID = id

		_, err = h.postSvc.GetByID(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		updatedPost, err := h.postSvc.Update(ctx, updatePost)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(updatedPost))
	}
}
