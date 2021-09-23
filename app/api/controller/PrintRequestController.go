package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/config"
	"OnlineJudge/constants"
	"OnlineJudge/core/database"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

func PrintRequest(c *gin.Context) {
	PrintLog := model.PrintLog{}
	PrintLogModel := model.PrintLog{}
	userModel := model.User{}
	PrintLogValidate := validate.PrintLogValidate
	userID := GetUserIdFromSession(c)

	now := time.Now().Unix()
	interval := config.GetWutOjConfig()["print_interval_time"].(int)
	redisStr := "user_last_print_request" + strconv.Itoa(int(userID))
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

	userInfo := userModel.FindUserByID(userID).Data.(model.User)

	PrintLog.UserNick = userInfo.Nick
	PrintLog.Status = 0
	printLogMap := helper.Struct2Map(PrintLog)

	if res, err := PrintLogValidate.ValidateMap(printLogMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	PrintLog.RequestAt = time.Now()

	res := PrintLogModel.AddPrintLog(PrintLog)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
