package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetAllContest(c *gin.Context) {
	contestJson := struct {
		PageNumber int `form:"page_number" json:"page_number"`
	}{}
	contestModel := model.Contest{}

	if err := c.ShouldBind(&contestJson); err == nil {
		res := contestModel.GetAllContest(constants.PageLimit*(contestJson.PageNumber-1), constants.PageLimit)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

}

func GetContestByID(c *gin.Context) {
	userID := int(GetUserIdFromSession(c))

	ContestIDRaw := c.Param("contest_id")
	ContestID, err := strconv.Atoi(ContestIDRaw)
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "比赛ID错误", 0))
		return
	}
	fmt.Println(ContestID)
	contestModel := model.Contest{}

	res := contestModel.GetContestById(ContestID)
	if res.Status != constants.CodeError {
		contestUserModel := model.ContestUser{}
		if participation := contestUserModel.CheckUserContest(userID, ContestID); participation.Status != constants.CodeSuccess {
			c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "尚未参赛，请参赛", 0))
			return
		}
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, "数据查找失败", res.Msg))
	}

}
func JoinContest(c *gin.Context) {
	var contestUserModel = model.ContestUser{}
	var contestUserJson model.ContestUser

	if err := c.ShouldBindQuery(&contestUserJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	if err := c.ShouldBindUri(&contestUserJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
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
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
	res := contestModel.GetContestStatus(ContestID.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetUserContest(c *gin.Context) {
	UserID := GetUserIdFromSession(c)
	log.Print(UserID)
	contestUserModel := model.ContestUser{}
	if UserID == 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "请先登陆", ""))
		return
	}
	res := contestUserModel.GetUserContest(int(UserID))
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func GetContestProblems(c *gin.Context) {
	userID := int(GetUserIdFromSession(c))

	var ContestJson model.Contest
	contestModel := model.Contest{}
	problemModel := model.Problem{}
	if err := c.ShouldBindUri(&ContestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", ""))
		return
	}

	contestID := ContestJson.ContestID
	contestUserModel := model.ContestUser{}
	if participation := contestUserModel.CheckUserContest(userID, contestID); participation.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "尚未参赛，请参赛", 0))
		return
	}

	contestValidate := validate.ContestValidate
	contestMap := helper.Struct2Map(ContestJson)

	if res, err := contestValidate.ValidateMap(contestMap, "getProblems"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	res := contestModel.GetContestProblems(ContestJson.ContestID)

	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	// 获取题目id
	type problemInfo struct {
		ID   uint                   `json:"id"`
		Info map[string]interface{} `json:"info""`
	}

	problemIDsStr := res.Data.(string)
	var problemIDs []uint

	if err := json.Unmarshal([]byte(problemIDsStr), &problemIDs); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "题目设置格式出错，请联系管理员", nil))
		return
	}

	fields := []string{"problem_id", "title"}
	res = problemModel.GetProblemFieldsByIDList(problemIDs, fields)

	// init result
	result := make([]problemInfo, 0)
	if res.Status != constants.CodeSuccess {
		// 查询题目数据失败
		for _, problemID := range problemIDs {
			item := problemInfo{ID: problemID, Info: make(map[string]interface{})}
			result = append(result, item)
		}
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "获取题目信息失败", result))
		return
	}

	problemInfos := res.Data.([]model.Problem)
	for _, info := range problemInfos {
		// gen problem info
		extraMap := map[string]interface{}{
			"title": info.Title,
		}

		item := problemInfo{ID: info.ProblemID, Info: extraMap}
		result = append(result, item)
	}
	c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeSuccess, "获取比赛题目信息成功", result))

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
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}
}

// MARK: 检测用户是否参赛
func CheckContest(c *gin.Context) {
	session := sessions.Default(c)
	UserID := int(session.Get("user_id").(uint))
	contestUserModel := model.ContestUser{}
	contestValidate := validate.ContestValidate

	var contestJson model.Contest
	if err := c.ShouldBindUri(&contestJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "数据模型绑定错误", err.Error()))
		return
	}

	contestMap := helper.Struct2Map(contestJson)
	if res, err := contestValidate.ValidateMap(contestMap, "join"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := contestUserModel.CheckUserContest(UserID, contestJson.ContestID)

	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return

}
