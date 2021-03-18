package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllDiscuss(c *gin.Context) {
	discussModel := model.Discuss{}

	res := discussModel.GetAllDiscuss()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

func GetDiscussionByID(c *gin.Context) {
	discussModel := model.Discuss{}
	discussValidate := validate.DiscussValidate

	discussJson := struct {
		id         int `json:"id" form:"id"`
		PageNumber int `json:"page_number" form:"page_number"`
	}{}

	if err := c.ShouldBind(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "", err.Error()))
		return
	}

	discussMap := helper.Struct2Map(discussJson)
	if res, err := discussValidate.ValidateMap(discussMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
	}

	res := discussModel.GetDiscussionByID(discussJson.id, discussJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetContestDiscussion(c *gin.Context) {
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

	res := discussModel.GetDiscussionByID(discussJson.ContestId, discussJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetProblemDiscussion(c *gin.Context) {
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

	res := discussModel.GetDiscussionByID(discussJson.ProblemId, discussJson.PageNumber)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddDiscussion(c *gin.Context) {
	discussModel := model.Discuss{}
	discussValidate := validate.DiscussValidate

	discussJson := struct {
		discuss model.Discuss
	}{}

	if err := c.ShouldBind(&discussJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	discussMap := helper.Struct2Map(discussJson)
	if res, err := discussValidate.ValidateMap(discussMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := discussModel.AddDiscussion(discussJson.discuss)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddReply(c *gin.Context) {
	replyModel := model.Reply{}
	replyValidate := validate.ReplyValidate

	replyJson := struct {
		reply model.Reply
	}{}

	if err := c.ShouldBind(&replyJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	replyMap := helper.Struct2Map(replyJson)
	if res, err := replyValidate.ValidateMap(replyMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := replyModel.AddReply(replyJson.reply)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
