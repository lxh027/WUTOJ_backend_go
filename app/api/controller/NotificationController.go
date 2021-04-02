package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/db_server"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"log"
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

	UserNick := GetUserNickFromSession(c)
	keyValue := "User:" + UserNick + ":Notification"
	log.Print(keyValue)
	var LastID int

	LastNotification, err := redis.Int(db_server.GetFromRedis(keyValue))

	if err != nil {
		LastID = 0
	} else {
		LastID = LastNotification
	}

	res, UpdateNotificationID := notificationModel.GetNotification(notificationJson.ContestID, LastID)

	if UpdateNotificationID == LastID {
		res.Msg = "无最新通知"
	}

	_ = db_server.DeleteFromRedis(keyValue)
	_ = db_server.PutToRedis(keyValue, UpdateNotificationID, 86400)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}
