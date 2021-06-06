package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "比赛期间不开放注册", 0))
	return
	var userModel = model.User{}
	var userValidate = validate.UserValidate
	var userSubmitLogModel = model.UserSubmitLog{}

	var userJson struct {
		model.User
		PasswordCheck string `json:"password_check" form:"password_check"`
	}
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据绑定模型错误", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "register"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "输入信息不完整或有误", err.Error()))
		return
	}

	if userJson.Password != userJson.PasswordCheck {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "两次密码输入不一致", ""))
	}

	userJson.Password = helper.GetMd5(userJson.Password)

	if userJson.Avatar == "" {
		userJson.Avatar = "../uploads/image/20200214/fc3d5f691e86c9f621621682c57de59b.jpg"
	}

	res := userModel.AddUser(userJson.User)
	res = userSubmitLogModel.CreatUserSubmitLog(userJson.Nick)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
