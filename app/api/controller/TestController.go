package controller

import (
	"OnlineJudge/app/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Test(c *gin.Context)  {

	var authModel = model.Authority{}

	auths, err := authModel.GetAllAuthority()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ERROR",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
		"data": auths,
	})
}