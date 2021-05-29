package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "strconv"
)

func GetAllContest(c *gin.Context) {
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

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	ContestID := c.Param("contest_id")
	fmt.Println(ContestID)
	contestModel := model.Contest{}

	res = contestModel.GetContestById(ContestID)
	if res.Status != common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		res = contestModel.GetContestByName(ContestID)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	}

}

func JoinContest(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

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

	res = contestUserModel.AddContestUser(contestUserJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))

}

func GetContestStatus(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	var contestModel = model.Contest{}
	var ContestID struct {
		ID int `form:"contest_id"`
	}
	if err := c.ShouldBind(&ContestID); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
	res = contestModel.GetContestStatus(ContestID.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserContest(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	UserID := GetUserIdFromSession(c)
	log.Print(UserID)
	contestUserModel := model.ContestUser{}
	if UserID == 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "请先登陆", ""))
		return
	}
	res = contestUserModel.GetUserContest(int(UserID))
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetContestProblems(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	var ContestJson model.Contest
	contestModel := model.Contest{}
	problemModel := model.Problem{}
	if err := c.ShouldBindUri(&ContestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", ""))
		return
	}
	contestValidate := validate.ContestValidate
	contestMap := helper.Struct2Map(ContestJson)

	if res, err := contestValidate.ValidateMap(contestMap, "getProblems"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}
	res = contestModel.GetContestProblems(ContestJson.ContestID)

	if res.Status != common.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	// 获取题目id
	type problemInfo struct {
		ID	uint 	`json:"id"`
		Info  map[string]interface{} `json:"info""`
	}

	problemIDsStr := res.Data.(string)
	var problemIDs []uint

	if err := json.Unmarshal([]byte(problemIDsStr), &problemIDs); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "题目设置格式出错，请联系管理员", nil))
		return
	}

	fields := []string{"problem_id", "title"}
	res = problemModel.GetProblemFieldsByIDList(problemIDs, fields)

	// init result
	result := make([]problemInfo, 0)
	if res.Status != common.CodeSuccess {
		// 查询题目数据失败
		for _, problemID := range problemIDs {
			item := problemInfo{ID: problemID, Info: make(map[string]interface{})}
			result =  append(result, item)
		}
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "获取题目信息失败", result))
		return
	}

	problemInfos := res.Data.([]model.Problem)
	for _, info := range problemInfos {
		// gen problem info
		extraMap := map[string]interface{} {
			"title": info.Title,
		}

		item := problemInfo{ID: info.ProblemID, Info: extraMap}
		result =  append(result, item)
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取比赛题目信息成功", result))

	return
}

func SearchContest(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == common.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

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

func CheckContest(c *gin.Context) {
	session := sessions.Default(c)
	UserID := int(session.Get("user_id").(uint))
	contestUserModel := model.ContestUser{}
	contestValidate := validate.ContestValidate

	var contestJson model.Contest
	if err := c.ShouldBindUri(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "join"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := contestUserModel.CheckUserContest(UserID, contestJson.ContestID)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}
