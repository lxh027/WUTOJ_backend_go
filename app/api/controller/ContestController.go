package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "strconv"
)

func GetAllContest(c *gin.Context) {

	LoginStatus := CheckLogin(c)
	if !LoginStatus {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "请先登录", ""))
		return
	}

	contestJson := struct {
		PageNumber int `form:"page_number" json:"page_number"`
	}{}
	contestModel := model.Contest{}

	if err := c.ShouldBind(&contestJson); err == nil {
		res := contestModel.GetAllContest(common.PageLimit*(contestJson.PageNumber-1), common.PageLimit)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

}

func GetContest(c *gin.Context) {
	ContestID := c.Param("contest_id")
	fmt.Println(ContestID)
	contestModel := model.Contest{}
	var res helper.ReturnType
	res = contestModel.GetContestById(ContestID)
	if res.Status != common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		res = contestModel.GetContestByName(ContestID)
		if res.Status != common.CodeError {
			c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
			return
		} else {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "查找失败", ""))
			return
		}

	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func JoinContest(c *gin.Context) {
	var contestUserModel = model.ContestUser{}
	var contestUserJson model.ContestUser

	if err := c.ShouldBindQuery(&contestUserJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	if err := c.ShouldBindUri(&contestUserJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	res := contestUserModel.AddContestUser(contestUserJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

func GetContestStatus(c *gin.Context) {
	var contestModel = model.Contest{}
	var ContestID struct {
		ID int `form:"contest_id"`
	}
	if err := c.ShouldBind(&ContestID); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
	res := contestModel.GetContestStatus(ContestID.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserContest(c *gin.Context) {
	var queryInfo struct {
		UserID int `form:"user_id"`
	}
	contestUserModel := model.ContestUser{}
	if err := c.ShouldBind(&queryInfo); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", ""))
		return
	}
	res := contestUserModel.GetUserContest(queryInfo.UserID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetContestProblems(c *gin.Context) {
	var ContestJson model.Contest
	contestModel := model.Contest{}
	if err := c.ShouldBind(&ContestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", ""))
		return
	}
	contestValidte := validate.ContestValidate
	contestMap := helper.Struct2Map(ContestJson)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, string(ContestJson.ContestID), 0))
	if res, err := contestValidte.ValidateMap(contestMap, "getProblems"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}
	res := contestModel.GetContestProblems(ContestJson.ContestID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func SearchContest(c *gin.Context) {
	contestJson := struct {
		Param string `uri:"param" json:"param"`
	}{}

	contestJson.Param = c.Param("param")
	contestModel := model.Contest{}

	if err := c.ShouldBind(&contestJson); err == nil {
		res := contestModel.GetContest(contestJson.Param)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
}
