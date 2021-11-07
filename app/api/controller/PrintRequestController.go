package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/config"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"fmt"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func PrintRequest(c *gin.Context) {
	PrintLog := model.PrintLog{}
	PrintLogModel := model.PrintLog{}
	PrintLogValidate := validate.PrintLogValidate
	userID := GetUserIdFromSession(c)
	if userID == 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "用户未登录", ""))
		return
	}

	now := time.Now().Unix()
	interval := config.GetWutOjConfig()["print_interval_time"].(int)
	redisStr := redis_key.LastPrintRequest(int(userID))
	if value, err := database.GetFromRedis(redisStr); err == nil {
		last, _ := redis.Int64(value, err)
		fmt.Printf("now: %v, last: %v\n", now, last)

		if now-last <= int64(interval) {
			c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "打印申请间隔过短，请10秒后再试", ""))
			return
		}
	}
	_ = database.PutToRedis(redisStr, now, 3600)

	if err := c.ShouldBindJSON(&PrintLog); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	printLogMap := helper.Struct2Map(PrintLog)

	if res, err := PrintLogValidate.ValidateMap(printLogMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	SubmitLogModel := model.Submit{}
	res := SubmitLogModel.GetSubmitInPrintRequest(uint(PrintLog.SubmitID), userID)

	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, res.Msg, 0))
		return
	}

	submit := res.Data.(model.Submit)

	if submit.UserID != userID {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "同学别打印了，这题你把握不住", 0))
		return
	}

	res = PrintLogModel.AddPrintLog(PrintLog)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
