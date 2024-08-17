package handler

import (
	"net/http"
	"post-api/dto"
	"post-api/service"

	"github.com/gin-gonic/gin"
)

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
			ctx.Error(err)
			return
		}

		ctx.JSON(http.StatusOK, dto.NewSuccessResponse(posts))
	}

}
