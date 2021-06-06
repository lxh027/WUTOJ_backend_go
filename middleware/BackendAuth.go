package middleware

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/controller"
	"OnlineJudge/constants"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BackendAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.FullPath())
		requestPath := c.FullPath()

		auth := constants.GetBackendRouterAuth(requestPath)
		fmt.Println(auth)
		if auth != constants.Pass {
			if res := haveAuth(c, auth); res != common.Authed {
				c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
				c.Abort()
			}
		}
	}
}

func haveAuth(c *gin.Context, authQuery string) int {
	session := sessions.Default(c)
	id := session.Get("userId")
	if id == nil {
		return common.UnLoggedIn
	} else if session.Get("identity").(uint) == 0 {
		return common.UnAuthed
	}
	_, auths, err := controller.GetUserAllAuth(id.(int))
	if err != nil {
		return common.AuthError
	} else {
		for _, auth := range auths {
			if auth == authQuery {
				return common.Authed
			}
		}
		return common.UnAuthed
	}
}