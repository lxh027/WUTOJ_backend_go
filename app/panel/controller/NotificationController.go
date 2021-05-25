package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetAllNotification(c *gin.Context) {
	if res := haveAuth(c, "getAllContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeModel := model.Notification{}

	noticeJson := struct {
		ContestID 	int 	`json:"contest_id" form:"contest_id"`
	}{}

	if err := c.ShouldBind(&noticeJson); err == nil {
		res := noticeModel.GetAllNotification(noticeJson.ContestID)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", nil))
	return
}

func GetNotificationByID(c *gin.Context) {
	if res := haveAuth(c, "getAllContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeModel := model.Notification{}

	var noticeJson model.Notification

	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}


	res := noticeModel.GetNotificationByID(noticeJson.ID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddNotification(c *gin.Context) {
	session := sessions.Default(c)

	if res := haveAuth(c, "updateContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeModel := model.Notification{}

	var noticeJson model.Notification
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	noticeJson.SubmitTime = time.Now()
	noticeJson.ModifyTime = time.Now()
	noticeJson.UserID = session.Get("userId").(int)

	res := noticeModel.AddNotification(noticeJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteNotification(c *gin.Context) {
	if res := haveAuth(c, "updateContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeModel := model.Notification{}

	var noticeJson model.Notification
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res := noticeModel.DeleteNotification(noticeJson.ID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateNotification(c *gin.Context) {
	if res := haveAuth(c, "updateContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}

	noticeModel := model.Notification{}

	var noticeJson model.Notification
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	noticeJson.ModifyTime = time.Now()

	res := noticeModel.UpdateNotification(noticeJson.ID, noticeJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeNotificationStatus(c *gin.Context) {
	if res := haveAuth(c, "updateContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}

	noticeModel := model.Notification{}

	var noticeJson model.Notification
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res := noticeModel.ChangeNotificationStatus(noticeJson.ID, noticeJson.Status)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}
