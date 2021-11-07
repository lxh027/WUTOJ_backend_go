package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DoLogin(c *gin.Context) {

	session := sessions.Default(c)

	if session.Get("user_id") != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "已登陆", data))
		return
	}

	var userModel = model.User{}
	var userValidate = validate.UserValidate

	var userJson model.User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据绑定模型错误", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "login"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "输入信息不完整或有误", err.Error()))
		return
	}

	userJson.Password = helper.GetMd5(userJson.Password)

	res := userModel.LoginCheck(userJson)

	if res.Status == constants.CodeSuccess {
		userInfo := res.Data.(map[string]interface{})["userInfo"].(model.User)
		submitLog := res.Data.(map[string]interface{})["submitLog"].(model.UserSubmitLog)
		returnData := map[string]interface{}{
			"userId":     userInfo.UserID,
			"nick":       userInfo.Nick,
			"desc":       userInfo.Desc,
			"avatar":     userInfo.Avatar,
			"submit_log": submitLog,
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

func DoLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "注销成功", session.Get("user_id")))
}

func ForgetPassword(c *gin.Context) {
	var userJson model.User
	userValidate := validate.UserValidate

	if err := c.ShouldBindQuery(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据绑定模型错误", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "forget"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "输入信息不完整或有误", err.Error()))
		return
	}

	res, err := helper.SendMail(userJson.Mail)

	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdatePassword(c *gin.Context) {
	userJson := struct {
		Mail          string `json:"mail" form:"mail"`
		Password      string `json:"password" form:"password"`
		VerifyCode    string `json:"verify_code" form:"verify_code"`
		PasswordCheck string `json:"password_check" form:"password_check"`
	}{}

	userModel := model.User{}
	userValidate := validate.UserValidate
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "update_password"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "输入信息不完整或有误", err.Error()))
		return
	}

	KeyValue := redis_key.VerifyCode(userJson.Mail)

	VerifyCodeCorrect, err := redis.String(database.GetFromRedis(KeyValue))
	database.DeleteFromRedis(KeyValue)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "验证码已过期，请重新发送验证码", err.Error()))
		return
	}
	log.Print(VerifyCodeCorrect)
	if userJson.VerifyCode != VerifyCodeCorrect {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "验证码输入错误", ""))
		return
	}

	user := model.User{}
	user.Mail = userJson.Mail
	user.Password = helper.GetMd5(userJson.Password)

	res := userModel.UpdatePassword(user)
	c.JSON(constants.CodeSuccess, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
