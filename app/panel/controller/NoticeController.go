package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetAllNotice(c *gin.Context)  {
	if res := haveAuth(c, "getAllNotice"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeModel := model.Notice{}

	noticeJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Title 	string 	`json:"title" form:"title"`
			BeginTime 	time.Time `json:"begintime" form:"begintime"`
		}
	}{}

	if c.ShouldBind(&noticeJson) == nil {
		noticeJson.Offset = (noticeJson.Offset-1)*noticeJson.Limit
		res := noticeModel.GetAllNotice(noticeJson.Offset, noticeJson.Limit, noticeJson.Where.Title, noticeJson.Where.BeginTime)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetNoticeByID(c *gin.Context) {
	if res := haveAuth(c, "getAllNotice"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeValidate := validate.NoticeValidate
	noticeModel := model.Notice{}

	var noticeJson model.Notice

	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	noticeMap := helper.Struct2Map(noticeJson)
	if res, err:= noticeValidate.ValidateMap(noticeMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := noticeModel.FindNoticeByID(noticeJson.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddNotice(c *gin.Context) {
	if res := haveAuth(c, "addNotice"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeValidate := validate.NoticeValidate
	noticeModel := model.Notice{}

	var noticeJson model.Notice
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	noticeMap := helper.Struct2Map(noticeJson)
	if res, err:= noticeValidate.ValidateMap(noticeMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := noticeModel.AddNotice(noticeJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteNotice(c *gin.Context) {
	if res := haveAuth(c, "deleteNotice"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeValidate := validate.NoticeValidate
	noticeModel := model.Notice{}

	var noticeJson model.Notice
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	noticeMap := helper.Struct2Map(noticeJson)
	if res, err:= noticeValidate.ValidateMap(noticeMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := noticeModel.DeleteNotice(noticeJson.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateNotice(c *gin.Context) {
	if res := haveAuth(c, "updateNotice"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	noticeValidate := validate.NoticeValidate
	noticeModel := model.Notice{}

	var noticeJson model.Notice
	if err := c.ShouldBind(&noticeJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	noticeMap := helper.Struct2Map(noticeJson)
	if res, err:= noticeValidate.ValidateMap(noticeMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := noticeModel.UpdateNotice(noticeJson.ID, noticeJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
