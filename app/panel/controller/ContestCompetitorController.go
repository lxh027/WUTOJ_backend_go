package controller

// //自建
// import (
// 	"OnlineJudge/app/common/validate"
// 	"OnlineJudge/app/helper"
// 	"OnlineJudge/app/panel/model"
// 	"OnlineJudge/constants"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// //TODO:

// //AddContestCompetitors 添加
// func AddContestCompetitors(c *gin.Context) {
// 	contestCompetitorValidate := validate.ContestCompetitorValidate
// 	contestCompetitorModel := model.ContestCompetitor{}

// 	//contestCompetitorsJSON
// 	contestCompetitorsJSON := struct {
// 		ContestID int    `json:"contest_id" form:"contest_id"`
// 		Rids      string `json:"rids" form:"rids"`
// 	}{}

// 	if err := c.ShouldBind(&contestCompetitorsJSON); err != nil {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
// 		return
// 	}

// 	contestCompetitorsMap := helper.Struct2Map(contestCompetitorsJSON)
// 	if res, err := contestCompetitorValidate.ValidateMap(contestCompetitorsMap, "addGroup"); !res {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
// 		return
// 	}

// 	var rids []int
// 	_ = json.Unmarshal([]byte((contestCompetitorsJSON.Rids)), &rids)
// 	fmt.Println(rids)
// 	for _, rid := range rids {
// 		res := contestCompetitorModel.AddContestCompetitor(model.ContestCompetitor{ContestID: contestCompetitorsJSON.ContestID, Rid: rid})
// 		if res.Status != constants.CodeSuccess {
// 			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, "编号为"+string(rune(rid))+"的角色添加失败", res.Data))
// 			return
// 		}

// 	}
// 	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "添加成功", true))
// 	return
// }

// //DeleteContestCompetitors 删除
// func DeleteContestCompetitors(c *gin.Context) {
// 	contestCompetitorValidate := validate.ContestCompetitorValidate
// 	contestCompetitorModel := model.ContestCompetitor{}

// 	contestCompetitorsJSON := struct {
// 		ContestID int    `json:"contest_id" form:"contest_id"`
// 		Rids      string `json:"rids" form:"rids"`
// 	}{}

// 	if err := c.ShouldBind(&contestCompetitorsJSON); err != nil {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
// 		return
// 	}

// 	contestCompetitorsMap := helper.Struct2Map(contestCompetitorsJSON)
// 	if res, err := contestCompetitorValidate.ValidateMap(contestCompetitorsMap, "deleteGroup"); !res {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
// 		return
// 	}

// 	var rids []int
// 	_ = json.Unmarshal([]byte((contestCompetitorsJSON.Rids)), &rids)
// 	fmt.Println(rids)
// 	for _, rid := range rids {
// 		res := contestCompetitorModel.DeleteContestCompetitor(model.ContestCompetitor{ContestID: contestCompetitorsJSON.ContestID, Rid: rid})
// 		if res.Status != constants.CodeSuccess {
// 			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, "编号为"+string(rune(rid))+"的权限删除失败", res.Data))
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "删除成功", true))
// 	return

// }

// //GetContestCompetitorsList 获取
// func GetContestCompetitorsList(c *gin.Context) {
// 	contestCompetitorValidate := validate.ContestCompetitorValidate
// 	competitorRoleModel := model.CompetitorRole{}

// 	competitorRoleJSON := struct {
// 		ContestID int `json:"contest_id" form:"contest_id"`
// 	}{}

// 	if err := c.ShouldBind(&competitorRoleJSON); err != nil {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
// 		return
// 	}

// 	competitorRoleMap := helper.Struct2Map(competitorRoleJSON)
// 	if res, err := contestCompetitorValidate.ValidateMap(competitorRoleMap, "getContestCompetitor"); !res {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
// 		return
// 	}

// 	allCompetitorRoles := competitorRoleModel.GetCompetitorRoleNoRules()

// 	res := competitorRoleModel.GetContestCompetitor(competitorRoleJSON.ContestID)
// 	competitorRoles := res.Data.([]model.CompetitorRole)
// 	var val []int
// 	for _, competitorRole := range competitorRoles {
// 		val = append(val, competitorRole.Rid)
// 	}
// 	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, map[string]interface{}{
// 		"allCompetitorRoles": allCompetitorRoles.Data,
// 		"values":             val,
// 	}))
// 	return
// }
