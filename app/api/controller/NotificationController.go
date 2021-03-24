package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotification(c *gin.Context) {

	notificationJson := struct {
		ContestID int `json:"contest_id" form:"contest_id" uri:"contest_id"`
	}{}
	notificationModel := model.Notification{}
	notificationValidate := validate.NotificationValidate

	if err := c.ShouldBindQuery(&notificationJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	NotificationMap := helper.Struct2Map(notificationJson)

	if res, err := notificationValidate.ValidateMap(NotificationMap, "get"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res := notificationModel.GetNotification(notificationJson.ContestID)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}
