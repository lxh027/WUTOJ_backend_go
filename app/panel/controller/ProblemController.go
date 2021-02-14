package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProblem(c *gin.Context)  {
	if res := haveAuth(c, "getAllProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemModel := model.Problem{}

	problemJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Title 	string 	`json:"title" form:"title"`
		}
	}{}

	if c.ShouldBind(&problemJson) == nil {
		problemJson.Offset = (problemJson.Offset-1)*problemJson.Limit
		res := problemModel.GetAllProblem(problemJson.Offset, problemJson.Limit, problemJson.Where.Title)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetProblemByID(c *gin.Context) {
	if res := haveAuth(c, "getAllProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem

	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.FindProblemByID(problemJson.ProblemID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddProblem(c *gin.Context) {
	if res := haveAuth(c, "addProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.AddProblem(problemJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteProblem(c *gin.Context) {
	if res := haveAuth(c, "deleteProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.DeleteProblem(problemJson.ProblemID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateProblem(c *gin.Context) {
	if res := haveAuth(c, "updateProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.UpdateProblem(problemJson.ProblemID, problemJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeProblemStatus(c *gin.Context) {
	if res := haveAuth(c, "updateProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.ChangeProblemStatus(problemJson.ProblemID, problemJson.Status)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}