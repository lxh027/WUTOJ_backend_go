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

func GetAllContest(c *gin.Context)  {
	if res := haveAuth(c, "getAllContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	contestModel := model.Contest{}

	contestJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			ContestName 	string 	`json:"contest_name" form:"contest_name"`
			BeginTime 	time.Time `json:"begin_time" form:"begin_time"`
		}
	}{}

	if c.ShouldBind(&contestJson) == nil {
		contestJson.Offset = (contestJson.Offset-1)*contestJson.Limit
		res := contestModel.GetAllContest(contestJson.Offset, contestJson.Limit, contestJson.Where.ContestName, contestJson.Where.BeginTime)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetContestByID(c *gin.Context) {
	if res := haveAuth(c, "getAllContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

	var contestJson model.Contest

	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err:= contestValidate.ValidateMap(contestMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.FindContestByID(contestJson.ContestID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddContest(c *gin.Context) {
	if res := haveAuth(c, "addContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err:= contestValidate.ValidateMap(contestMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.AddContest(contestJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteContest(c *gin.Context) {
	if res := haveAuth(c, "deleteContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err:= contestValidate.ValidateMap(contestMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.DeleteContest(contestJson.ContestID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateContest(c *gin.Context) {
	if res := haveAuth(c, "updateContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err:= contestValidate.ValidateMap(contestMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.UpdateContest(contestJson.ContestID, contestJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeContestStatus(c *gin.Context) {
	if res := haveAuth(c, "updateContest"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	contestValidate := validate.ContestValidate
	contestModel := model.Contest{}

	var contestJson model.Contest
	if err := c.ShouldBind(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err:= contestValidate.ValidateMap(contestMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := contestModel.ChangeContestStatus(contestJson.ContestID, contestJson.Status)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
