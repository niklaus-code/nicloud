package vmerror

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Message string
}

func (err Error) Error() string {
	return err.Message
}

func REQUESTERROR(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"err": err.Error(),
	})
}

func SERVERERROR(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"err": Error{Message: err.Error()},
		})
	}
}

func SUCCESS(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"res": data,
		"err": nil,
	})
}
