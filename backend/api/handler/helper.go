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

func QueryNumeric(ctx *gin.Context, key string, defaultVal int64) (int64, error) {
	query := ctx.Query(key)
	if query == "" {
		return defaultVal, nil
	}

	id, err := strconv.Atoi(query)
	if err != nil {
		return 0, apperror.NewError(
			err,
			"helper.go",
			fmt.Sprintf("QueryNumeric(%s)", key),
			fmt.Sprintf("strconv.Atoi(%s)", query),
		)
	}

	return int64(id), nil
}
