package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DoLogin(c *gin.Context) {

	session := sessions.Default(c)

	if session.Get("user_id") != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, helper.ApiReturn(common.CODE_ERROE, "已登陆", data))
		return
	}

	var userModel = model.User{}
	var userValidate = validate.UserValidate

	var userJson model.User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "login"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	userJson.Password = helper.GetMd5(userJson.Password)

	res := userModel.LoginCheck(userJson)


	if res.Status == common.CODE_SUCCESS {
		userInfo := res.Data.(map[string]interface{})["userInfo"].(model.User)
		allProblem := res.Data.(map[string]interface{})["allProblem"].([]model.Submit)
		returnData := map[string]interface{} {
			"userId"	: userInfo.UserID,
			"nick"		: userInfo.Nick,
			"desc"		: userInfo.Desc,
			"avatar"	: userInfo.Avatar,
			"all_problem": allProblem,
		}
		jsonData, _ := json.Marshal(returnData)
		session.Set("user_id", userInfo.UserID)
		session.Set("nick", userInfo.Nick)
		session.Set("identity", userInfo.Identity)
		session.Set("data", string(jsonData))
		session.Save()
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, returnData))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DoLogout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, helper.ApiReturn(common.CODE_SUCCESS, "注销成功", session.Get("user_id")))
}
