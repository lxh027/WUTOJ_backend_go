package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(c *gin.Context)  {

	testJson := struct {
		Name string 	`uri:"name"`
		ID 	string  	`form:"id"`
	}{}

	if err1, err2 := c.ShouldBindUri(&testJson), c.ShouldBind(&testJson); err1 != nil && err2 != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", ""))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", testJson))
	return
	c.JSON(http.StatusOK, gin.H{
		"msg": "WUT OnlineJudge",
	})
}