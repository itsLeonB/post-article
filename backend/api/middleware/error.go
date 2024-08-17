package middleware

import (
	"net/http"
	"post-api/apperror"
	"post-api/dto"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0].Err.(*apperror.AppError).Err
			switch e := err.(type) {
			case *dto.ErrorResponse:
				abortWithError(ctx, e.Code, e)
			default:
				abortWithError(ctx, http.StatusInternalServerError, dto.InternalServerError())
			}
		}
	}
}

func abortWithError(ctx *gin.Context, statusCode int, e *dto.ErrorResponse) {
	ctx.AbortWithStatusJSON(statusCode, dto.NewErrorResponse(e))
}
