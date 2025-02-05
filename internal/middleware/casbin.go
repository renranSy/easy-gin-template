package middleware

import (
	"easy-gin-template/internal/config"
	"easy-gin-template/internal/enum"
	"easy-gin-template/internal/mods/admin/schema"
	"easy-gin-template/pkg/errors"
	"easy-gin-template/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func (a *Middleware) Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user       = c.MustGet(enum.UserInfoContext).(*schema.Admin)
			roleIdList = make([]int, 0)
		)

		a.DB.Model(&schema.AdminRole{}).Select("role_id").Where("admin_id = ?", user.ID).Find(&roleIdList)

		// 获取请求的路径和方法
		method := c.Request.Method
		path := c.Request.URL.Path
		path = strings.Replace(path, config.CFG.Server.Prefix, "", 1)

		re := regexp.MustCompile(`/\d+`)

		path = re.ReplaceAllString(path, "/:id")

		// 校验用户是否有权限
		hasPermission := false
		for _, role := range roleIdList {
			ok, err := a.Enforcer.Enforce(strconv.Itoa(role), path, method)
			if err != nil {
				log.Printf("Casbin enforce failed: %v", err)
				util.ResError(c, errors.InternalServerError("权限校验失败"))
				c.Abort()
				return
			}
			if ok {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			util.ResError(c, errors.Unauthorized("无访问权限"))
			c.Abort()
			return
		}

		c.Next()
	}
}
