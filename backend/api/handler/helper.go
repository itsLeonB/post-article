package handler

import (
	"fmt"
	"post-api/apperror"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPathID(ctx *gin.Context) (int64, error) {
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, apperror.NewError(
			err,
			"helper.go",
			"GetPathID()",
			fmt.Sprintf("strconv.Atoi(%s)", param),
		)
	}

	return int64(id), nil
}
