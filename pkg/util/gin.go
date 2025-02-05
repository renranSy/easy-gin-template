package util

import (
	"easy-gin-template/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResJSON(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	})
	c.Abort()
}

func ResOK(c *gin.Context, data interface{}, msg string) {
	ResJSON(c, http.StatusOK, data, msg)
}

func ResError(c *gin.Context, err error) {
	var reqErr *errors.Error
	if e, ok := errors.As(err); ok {
		reqErr = e
	} else {
		reqErr = errors.InternalServerError("")
	}
	ResJSON(c, reqErr.Status, nil, reqErr.Message)
}
