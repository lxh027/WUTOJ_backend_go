package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSamplesByProblemID(c *gin.Context) {
	if res := haveAuth(c, "getAllProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	sampleModel := model.Sample{}

	var sampleJson model.Sample

	if err := c.ShouldBind(&sampleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res := sampleModel.FindSamplesByProblemID(sampleJson.ProblemID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddSample(c *gin.Context) {
	if res := haveAuth(c, "addProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	sampleModel := model.Sample{}

	var sampleJson model.Sample
	if err := c.ShouldBind(&sampleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res := sampleModel.AddSample(sampleJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteSample(c *gin.Context) {
	if res := haveAuth(c, "deleteProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	sampleModel := model.Sample{}

	var sampleJson model.Sample
	if err := c.ShouldBind(&sampleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res := sampleModel.DeleteSample(sampleJson.SampleID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateSample(c *gin.Context) {
	if res := haveAuth(c, "updateProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	sampleModel := model.Sample{}

	var sampleJson model.Sample
	if err := c.ShouldBind(&sampleJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	res := sampleModel.UpdateSample(sampleJson.SampleID, sampleJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}
