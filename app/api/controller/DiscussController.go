package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/config"
	"OnlineJudge/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllDiscuss(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	discussModel := model.Discuss{}

	discussJson := struct {
		//discuss model.Discuss
		ContestID  int `json:"contest_id" form:"contest_id"`
		ProblemID  int `json:"problem_id" form:"problem_id"`
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	ConfigMap := config.GetWutOjConfig()
	Limit := ConfigMap["page_limit"].(int)

	if err := c.ShouldBindUri(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据绑定失败", ""))
		return
	}

	if c.ShouldBindQuery(&discussJson) == nil {
		res := discussModel.GetAllDiscuss(discussJson.ContestID, discussJson.ProblemID, Limit*(discussJson.PageNumber-1), Limit)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return

}

func GetDiscussionByID(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	discussModel := model.Discuss{}
	discussValidate := validate.DiscussValidate

	discussJson := struct {
		DiscussID  int `json:"discuss_id" form:"discuss_id"`
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	if err := c.ShouldBind(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "", err.Error()))
		return
	}

	discussMap := helper.Struct2Map(discussJson)
	if res, err := discussValidate.ValidateMap(discussMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res = discussModel.GetDiscussionByID(discussJson.DiscussID, discussJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetContestDiscussion(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	discussModel := model.Discuss{}
	discussValidate := validate.DiscussValidate

	discussJson := struct {
		ContestId  int `json:"contest_id" form:"contest_id"`
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	if err := c.ShouldBind(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "", err.Error()))
		return
	}

	discussMap := helper.Struct2Map(discussJson)
	if res, err := discussValidate.ValidateMap(discussMap, "findByContestID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
	}

	res = discussModel.GetDiscussionByID(discussJson.ContestId, discussJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetProblemDiscussion(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	discussModel := model.Discuss{}
	discussValidate := validate.DiscussValidate

	discussJson := struct {
		ProblemId  int `json:"problem_id" form:"problem_id"`
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	if err := c.ShouldBind(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "", err.Error()))
		return
	}

	discussMap := helper.Struct2Map(discussJson)
	if res, err := discussValidate.ValidateMap(discussMap, "findByProblemID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
	}

	res = discussModel.GetDiscussionByID(discussJson.ProblemId, discussJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddDiscussion(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	discussModel := model.Discuss{}
	discussValidate := validate.DiscussValidate

	var discussJson model.Discuss
	//discussJson.Time = time.Now()

	//log.Print(discussJson)
	if err := c.ShouldBind(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	discussMap := helper.Struct2Map(discussJson)
	if res, err := discussValidate.ValidateMap(discussMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res = discussModel.AddDiscussion(discussJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddReply(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	replyModel := model.Reply{}
	replyValidate := validate.ReplyValidate

	replyJson := model.Reply{}

	if err := c.ShouldBind(&replyJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	replyJson.UserID = int(middleware.GetUserIdFromSession(c))
	replyMap := helper.Struct2Map(replyJson)
	if res, err := replyValidate.ValidateMap(replyMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据校验失败", err.Error()))
		return
	}

	res = replyModel.AddReply(replyJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
