package middleware

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/controller"
	"OnlineJudge/constants"
	"OnlineJudge/constants/backend_auth"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func BackendAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.FullPath())
		requestPath := c.FullPath()

		auth := backend_auth.GetBackendRouterAuth(requestPath)
		fmt.Println(auth)
		if auth != backend_auth.AuthPass {
			if res := haveAuth(c, auth); res != constants.Authed {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "权限不足", res))
				c.Abort()
			}
		}
	}
}

func haveAuth(c *gin.Context, authQuery string) int {
	session := sessions.Default(c)
	id := session.Get("userId")
	if id == nil {
		return constants.UnLoggedIn
	} else if session.Get("identity").(uint) == 0 {
		return constants.UnAuthed
	}
	_, auths, err := controller.GetUserAllAuth(id.(int))
	if err != nil {
		return constants.AuthError
	} else {
		for _, auth := range auths {
			if auth == authQuery {
				return constants.Authed
			}
		}
		return constants.UnAuthed
	}
}
