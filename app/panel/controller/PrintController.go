package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func PrintRequest(c *gin.Context) {
	printLogModel := model.PrintLog{}
	submitModel := model.Submit{}
	userModel := model.User{}

	PrintRequestJson := struct {
		PrintID string `json:"print_id" form:"print_id"`
	}{}

	if err := c.ShouldBind(&PrintRequestJson); err != nil {
		log.Print(PrintRequestJson)
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "数据模型绑定失败", err.Error()))
		return
	}
	log.Print(PrintRequestJson)
	res := printLogModel.GetPrintLogByID(PrintRequestJson.PrintID)
	log.Print(res)
	requestInfo := res.Data.(model.PrintLog)

	res = submitModel.FindSubmitByID(requestInfo.SubmitID)
	submitInfo := res.Data.(model.Submit)

	res = userModel.GetUserByID(int(submitInfo.UserID))
	userInfo := res.Data.(model.User)

	requestInfo.Status = 1
	requestInfo.PrintAt = time.Now()

	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, gin.H{
		"print_id":    printLogModel.ID,
		"user_nick":   userInfo.Nick,
		"source_code": submitInfo.SourceCode,
		"print_at":    requestInfo.PrintAt,
	}))
	printLogModel.UpdateStatusAfterPrint(requestInfo.ID, requestInfo)
	return
}

func GetAllPrintRequest(c *gin.Context) {
	printRequestModel := model.PrintLog{}

	PrintRequestJson := struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	}{}

	if err := c.ShouldBind(&PrintRequestJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "数据模型绑定失败", err.Error()))
		return
	}

	PrintRequestJson.Offset = (PrintRequestJson.Offset - 1) * PrintRequestJson.Limit

	res := printRequestModel.GetAllPrintLog(PrintRequestJson.Offset, PrintRequestJson.Limit)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}
