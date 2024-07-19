package middleware

import (
	"easy-gin-template/internal/db"
	"easy-gin-template/internal/enum"
	"easy-gin-template/internal/model"
	"easy-gin-template/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user  model.User
			token = c.GetHeader("Authorization")
		)

		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "请先登录"})
			c.Abort()
			return
		}

		claims, err := jwt.Jwt.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "登录已过期"})
			c.Abort()
			return
		}

		if err := db.DB.Select("*").Where("id = ?", claims.UserId).Take(&user).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "message": "用户不存在"})
			c.Abort()
			return
		}
		c.Set(enum.UserInfoContext, user)

		c.Next()
	}
}
