package middleware

import (
	"easy-gin-template/internal/enum"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/errors"
	"easy-gin-template/pkg/jwt"
	"easy-gin-template/pkg/util"
	"github.com/gin-gonic/gin"
)

func (a *Middleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user  *schema.Admin
			token = c.GetHeader("Authorization")
		)

		if token == "" {
			util.ResError(c, errors.Forbidden("请先登录"))
			return
		}

		claims, err := jwt.Jwt.ParseToken(token)
		if err != nil {
			util.ResError(c, errors.Forbidden("登录已过期"))
			return
		}

		if err := a.DB.Select("*").Where("id = ?", claims.UserId).Take(&user).Error; err != nil {
			util.ResError(c, errors.Forbidden("用户不存在"))
			return
		}
		c.Set(enum.UserInfoContext, user)

		c.Next()
	}
}
