package handler

import (
	"context"
	"net/http"
	"post-api/appcontext"
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
		limit, err := QueryNumeric(ctx, "limit", 0)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		offset, err := QueryNumeric(ctx, "offset", 0)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		statusID, err := QueryNumeric(ctx, "status_id", 0)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		pageCtx := context.WithValue(ctx, appcontext.KeyLimit, limit)
		pageCtx = context.WithValue(pageCtx, appcontext.KeyOffset, offset)
		pageCtx = context.WithValue(pageCtx, appcontext.KeyStatusID, statusID)

		posts, err := h.postSvc.GetAll(pageCtx)
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

		updatedPost, err := h.postSvc.Update(ctx, updatePost)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(updatedPost))
	}
}

func (h *PostHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := GetPathID(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		err = h.postSvc.Delete(ctx, id)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h *PostHandler) GetStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		statuses, err := h.postSvc.GetStatus(ctx)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(statuses))
	}
}
