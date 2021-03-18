package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProblems(c *gin.Context) {

	problemModel := model.Problem{}
	res := problemModel.GetAllProblems()
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}

func GetProblemByID(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem

	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.GetProblemByID(int(problemJson.ProblemID))
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func SearchProblem(c *gin.Context) {
	// to do
}
