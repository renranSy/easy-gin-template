package router

import (
	v1 "easy-gin-template/internal/api/v1"
	"easy-gin-template/internal/db"
	"easy-gin-template/internal/enum"
	"easy-gin-template/internal/middleware"
	"easy-gin-template/internal/model"
	"easy-gin-template/internal/util"
	"easy-gin-template/pkg/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func setupUserRouter() {
	group.POST("/user/login", loginHandler)
	group.GET("/user/get", middleware.Auth(), getUserHandler)
}

func getUserHandler(c *gin.Context) {
	var (
		user = c.MustGet(enum.UserInfoContext).(model.User)
	)

	c.JSON(http.StatusOK, v1.Response{
		Code:    200,
		Data:    util.BeanCopy(&user, &v1.UserInfo{}),
		Message: "ok",
	})
}

func loginHandler(c *gin.Context) {
	var (
		p    v1.LoginReq
		user model.User
	)

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, v1.Response{
			Code:    400,
			Message: fmt.Sprintf("请求参数错误[%s]", err.Error()),
			Data:    nil,
		})
		return
	}

	if err := db.DB.Select("id, user, password").Where("username = ? ADN password = ?", p.Username, p.Password).Take(&user).Error; err != nil {
		c.JSON(http.StatusOK, v1.Response{
			Code:    400,
			Message: "用户名或密码错误",
			Data:    nil,
		})
		return
	}

	token, err := jwt.Jwt.GenToken(strconv.FormatInt(user.ID, 10), time.Now().Add(2*time.Hour))

	if err != nil {
		c.JSON(http.StatusOK, v1.Response{
			Code:    500,
			Message: "系统错误",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, v1.Response{
		Code:    200,
		Message: "登录成功",
		Data: v1.LoginRes{
			Token: token,
			UserInfo: v1.UserInfo{
				ID:       user.ID,
				Username: user.Username,
			},
		},
	})
}
