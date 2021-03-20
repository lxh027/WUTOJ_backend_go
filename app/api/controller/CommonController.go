package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckContest(c *gin.Context) {

}

func CheckLogin(c *gin.Context) bool {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		return true
	}
	return false
}
