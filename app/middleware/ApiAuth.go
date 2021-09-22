package middleware

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var id interface{}
		if id = session.Get("user_id"); id == nil {
			c.JSON(http.StatusBadRequest, helper.ApiReturn(constants.CodeError, "未登录，请先登录", -1))
			c.Abort()
		}
		c.Next()
	}
}

func ContestAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 比赛开始鉴权
		c.Next()
	}
}
