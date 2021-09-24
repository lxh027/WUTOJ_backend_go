package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "WUT OnlineJudge",
	})
}
