package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserInfo(c *gin.Context) {
	res := checkLogin(c)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	var userJson model.User
	userModel := model.User{}
	userValidate := validate.UserValidate

	if err := c.ShouldBindUri(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userJson.UserID = GetUserIdFromSession(c)

	log.Print(userJson)
	userMap := helper.Struct2Map(userJson)

	if res, err := userValidate.ValidateMap(userMap, "update_info"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res = userModel.EditUserByID(userJson.UserID, userJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func SearchUser(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	userJson := struct {
		Param string `json:"param" form:"param" uri:"param"`
	}{}
	userModel := model.User{}

	if err := c.ShouldBindUri(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res = userModel.SearchUser(userJson.Param)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetUserByID(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	userJson := struct {
		UserID uint `json:"user_id" form:"user_id" uri:"user_id"`
	}{}
	userModel := model.User{}

	if err := c.ShouldBindUri(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	log.Print("-----")
	log.Print(userJson)
	res = userModel.FindUserByID(userJson.UserID)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}
