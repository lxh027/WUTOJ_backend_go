package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/config"
	"OnlineJudge/judger"
	"github.com/gin-gonic/gin"
	_ "io"
	"net/http"
)

func Submit(c *gin.Context) {
	submitModel := model.Submit{}
	var submitJson model.Submit
	submitValidate := validate.SubmitValidate
	//session := sessions.Default(c)
	//UserID := session.Get("user_id")
	//
	//if UserID == nil {
	//	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "用户未登录", ""))
	//	return
	//}

	//submitJson.UserID = uint(UserID.(int))
	submitJson.UserID = 1
	if err := c.ShouldBind(&submitJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据绑定模型错误", err.Error()))
		return
	}

	submitMap := helper.Struct2Map(submitJson)

	if res, err := submitValidate.ValidateMap(submitMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), ""))
		return
	}

	res := submitModel.AddSubmit(submitJson)
	go func(submit model.Submit) {
		judge(submit)
	}(submitJson)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}

func judge(submit model.Submit) {
	submitData := judger.SubmitData{
		Id:           uint64(submit.ID),
		Pid:          uint64(submit.ProblemID),
		Language:     helper.LanguageType(submit.Language),
		Code:         submit.SourceCode,
		BuildScript:  "",
		RootfsConfig: nil,
	}

	callback := func(id uint64, result judger.JudgeResult) {
		// Put Result To DB
		if result.IsFinished {
			data := map[string]interface{}{
				"status": result.Status,
				"time":   result.Time,
				"memory": result.Memory,
				"msg":    result.Msg,
			}
			submitModel := model.Submit{}
			submitModel.UpdateStatusAfterSubmit(int(id), data)
		}
	}

	instance := judger.GetInstance()

	go instance.Submit(submitData, callback)
}

func GetSubmitInfo(c *gin.Context) {

	submitModel := model.Submit{}

	submitJson := struct {
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	ConfigMap := config.GetWutOjConfig()
	Limit := ConfigMap["page_limit"].(int)

	if c.ShouldBind(&submitJson) == nil {
		res := submitModel.GetAllSubmit(Limit*(submitJson.PageNumber-1), Limit)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))

}
