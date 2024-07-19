package middleware

import (
	"easy-gin-template/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GinLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		log.Logger.Info(fmt.Sprintf("| %d | \t%s | \t\t\t%s | %s \t\"%s\" %s\n",
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		))
		return ""
	})
}
