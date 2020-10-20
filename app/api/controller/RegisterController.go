package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context)  {
	var userModel = model.User{}
	var userValidate = validate.UserValidate

	if res, err := userValidate.Validate(c, "register"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	password, passwordCheck := c.PostForm("password"), c.PostForm("password_check")

	if password != passwordCheck {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "两次密码输入不一致", ""))
	}

	var userJson model.User
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	userJson.Password = common.GetMd5(userJson.Password)

	if userJson.Avatar == "" {
		userJson.Avatar = "../uploads/image/20200214/fc3d5f691e86c9f621621682c57de59b.jpg"
	}

	res := userModel.AddUser(userJson)

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}