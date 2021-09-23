package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PrintRequest(c *gin.Context) {
	printLogModel := model.PrintLog{}

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

	requestInfo.Status = 1
	requestInfo.PrintAt = time.Now()

	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, gin.H{
		"print_id":    requestInfo.ID,
		"user_nick":   requestInfo.UserNick,
		"source_code": requestInfo.Code,
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
