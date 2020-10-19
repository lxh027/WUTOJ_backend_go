package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Register(c *gin.Context)  {
	var userModel = model.User{}
	var userValidate = validate.UserValidate

	if res, err := userValidate.Validate(c, "register"); !res {
		log.Println(err.Error())
		return
	}

	password, passwordCheck := c.PostForm("password"), c.PostForm("password_check")

	if password != passwordCheck {
		c.JSON(http.StatusOK, gin.H{
			"msg": "两次密码输入不一致",
		})
	}

	var userJson model.User
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": err.Error(),
		})
		return
	}

	userJson.Password = common.GetMd5(userJson.Password)

	if userJson.Avatar == "" {
		userJson.Avatar = "../uploads/image/20200214/fc3d5f691e86c9f621621682c57de59b.jpg"
	}

	res, err := userModel.AddUser(userJson)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	} else if !res {
		c.JSON(http.StatusOK, gin.H{
			"msg": "插入失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "插入成功",
		})
	}
	return
}