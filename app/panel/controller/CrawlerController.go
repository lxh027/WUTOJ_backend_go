package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"OnlineJudge/core/nsqueue"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO:搞了一些

//GetUserSubmit 发送爬取请求
func GetUserSubmit(c *gin.Context) {

	ojWebDataValidate := validate.OJWebDataValidate
	// var ojWebUserData model.OJWebUserData
	ojWebUserConfigModel := model.OJWebUserConfig{}

	// ojWebUserModel := model.OJWebUserData{}
	ojUserJSON := struct {
		UserID int    `json:"user_id" form:"user_id"`
		OJName string `json:"oj_name" form:"oj_name"`
		Status int    `json:"status" form:"status"`
	}{}
	if err := c.ShouldBind(&ojUserJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	ojWebUserDataMap := helper.Struct2Map(ojUserJSON)
	if res, err := ojWebDataValidate.ValidateMap(ojWebUserDataMap, "getAll"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	re := ojWebUserConfigModel.GetUserOJwebUserConfig(ojUserJSON.UserID, ojUserJSON.OJName)
	if re.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
		return
	}

	ojWebUserConfigRes := re.Data.([]model.OJWebUserConfig)
	// var userDatas []model.OJWebUserData

	// 定义客户端发送的request
	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

		spiderReq := &nsqueue.Request{
			Targets: []*nsqueue.TargetInfo{
				{
					UserInfo: &nsqueue.UserInfo{
						Name: fmt.Sprint("%d", ojUserJSON.UserID),
					},
					Oj: []*nsqueue.OJAccount{
						{
							OjName: ojUserJSON.OJName,
							Id:     []string{ojWebUserConfigRe.OJUserName},
						},
					},
				},
			},
			Status: ojUserJSON.Status,
		}
		nsqueue.RequestProducer.Publish(*spiderReq)
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "爬取请求发送成功", true))
	return
}

// //GetUserLastWeekSubmit 获取用户上周提交记录
// func GetUserLastWeekSubmit(c *gin.Context) {

// 	ojWebDataValidate := validate.OJWebDataValidate
// 	var ojWebUserData model.OJWebUserData
// 	ojWebUserConfigModel := model.OJWebUserConfig{}

// 	ojWebUserModel := model.OJWebUserData{}
// 	ojUserJSON := struct {
// 		UserID int    `json:"user_id" form:"user_id"`
// 		OJName string `json:"oj_name" form:"oj_name"`
// 	}{}
// 	if err := c.ShouldBind(&ojUserJSON); err != nil {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
// 		return
// 	}
// 	ojWebUserDataMap := helper.Struct2Map(ojUserJSON)
// 	if res, err := ojWebDataValidate.ValidateMap(ojWebUserDataMap, "getLastWeek"); !res {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
// 		return
// 	}

// 	re := ojWebUserConfigModel.GetUserOJwebUserConfig(ojUserJSON.UserID)
// 	if re.Status != constants.CodeSuccess {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
// 		return
// 	}
// 	cPd := pd.NewApiClient(rpcconn.RPCConn)

// 	ojWebUserConfigRes := re.Data.([]model.OJWebUserConfig)
// 	var userDatas []model.OJWebUserData

// 	// 定义客户端发送的request
// 	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

// 		spiderReq := &pd.Request{
// 			Targets: []*pd.TargetInfo{
// 				{
// 					UserInfo: &pd.UserInfo{
// 						Name: fmt.Sprint("%d", ojUserJSON.UserID),
// 					},
// 					Oj: []*pd.OJAccount{
// 						{
// 							OjName: ojUserJSON.OJName,
// 							Id:     []string{ojWebUserConfigRe.OJUserName},
// 						},
// 					},
// 				},
// 			},
// 		}
// 		res, err := cPd.GetLastWeek(context.Background(), spiderReq)
// 		if err != nil {
// 			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "启动爬虫服务失败", err.Error()))
// 			return
// 		}

// 		for _, data := range res.GetData() {
// 			for key, value := range data.GetData() {
// 				ojWebUserData.OJName = key
// 				ojWebUserData.UserID = ojUserJSON.UserID
// 				solvedData := value.GetData()
// 				for i := range solvedData.Problems {
// 					submitTime, err := time.Parse("2006-01-02 15:04:05 Z0700 MST", solvedData.Problems[i].GetSubmitTime())
// 					if err != nil {
// 						continue
// 					}
// 					ojWebUserData.SubmitTime = submitTime
// 					ojWebUserData.Status = solvedData.Problems[i].GetStatusWord()
// 					ojWebUserData.ProblemID = solvedData.Problems[i].GetProblemTitle()
// 					userDatas = append(userDatas, ojWebUserData)
// 				}
// 			}
// 		}

// 	}
// 	//操作db
// 	re = ojWebUserModel.AddOJWebUserDatas(userDatas)
// 	if re.Status != constants.CodeSuccess {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "爬取成功", true))
// 	return
// }

// //GetUserWeekSubmit 获取用户本周提交记录
// func GetUserWeekSubmit(c *gin.Context) {
// 	ojWebDataValidate := validate.OJWebDataValidate
// 	var ojWebUserData model.OJWebUserData
// 	ojWebUserConfigModel := model.OJWebUserConfig{}

// 	ojWebUserModel := model.OJWebUserData{}
// 	ojUserJSON := struct {
// 		UserID int    `json:"user_id" form:"user_id"`
// 		OJName string `json:"oj_name" form:"oj_name"`
// 	}{}
// 	if err := c.ShouldBind(&ojUserJSON); err != nil {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
// 		return
// 	}
// 	ojWebUserDataMap := helper.Struct2Map(ojUserJSON)
// 	if res, err := ojWebDataValidate.ValidateMap(ojWebUserDataMap, "getWeek"); !res {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
// 		return
// 	}

// 	re := ojWebUserConfigModel.GetUserOJwebUserConfig(ojUserJSON.UserID)
// 	if re.Status != constants.CodeSuccess {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
// 		return
// 	}
// 	cPd := pd.NewApiClient(rpcconn.RPCConn)

// 	ojWebUserConfigRes := re.Data.([]model.OJWebUserConfig)
// 	var userDatas []model.OJWebUserData

// 	// 定义客户端发送的request
// 	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

// 		spiderReq := &pd.Request{
// 			Targets: []*pd.TargetInfo{
// 				{
// 					UserInfo: &pd.UserInfo{
// 						Name: fmt.Sprint("%d", ojUserJSON.UserID),
// 					},
// 					Oj: []*pd.OJAccount{
// 						{
// 							OjName: ojUserJSON.OJName,
// 							Id:     []string{ojWebUserConfigRe.OJUserName},
// 						},
// 					},
// 				},
// 			},
// 		}
// 		res, err := cPd.GetWeek(context.Background(), spiderReq)
// 		if err != nil {
// 			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "启动爬虫服务失败", err.Error()))
// 			return
// 		}

// 		for _, data := range res.GetData() {
// 			for key, value := range data.GetData() {
// 				ojWebUserData.OJName = key
// 				ojWebUserData.UserID = ojUserJSON.UserID
// 				solvedData := value.GetData()
// 				for i := range solvedData.Problems {
// 					submitTime, err := time.Parse("2006-01-02 15:04:05 Z0700 MST", solvedData.Problems[i].GetSubmitTime())
// 					if err != nil {
// 						continue
// 					}
// 					ojWebUserData.SubmitTime = submitTime
// 					ojWebUserData.Status = solvedData.Problems[i].GetStatusWord()
// 					ojWebUserData.ProblemID = solvedData.Problems[i].GetProblemTitle()
// 					userDatas = append(userDatas, ojWebUserData)
// 				}
// 			}
// 		}

// 	}
// 	//操作db
// 	re = ojWebUserModel.AddOJWebUserDatas(userDatas)
// 	if re.Status != constants.CodeSuccess {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "爬取成功", true))
// 	return
// }

// //GetUserBetweenSubmit 获取用户间隔提交记录
// func GetUserBetweenSubmit(c *gin.Context) {

// 	ojWebDataValidate := validate.OJWebDataValidate
// 	var ojWebUserData model.OJWebUserData
// 	ojWebUserConfigModel := model.OJWebUserConfig{}

// 	ojWebUserModel := model.OJWebUserData{}
// 	ojUserJSON := struct {
// 		UserID    int    `json:"user_id" form:"user_id"`
// 		OJName    string `json:"oj_name" form:"oj_name"`
// 		StartTime string `json:"start_time" form:"start_time"`
// 		EndTime   string `json:"end_time" form:"end_time"`
// 	}{}
// 	if err := c.ShouldBind(&ojUserJSON); err != nil {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
// 		return
// 	}
// 	ojWebUserDataMap := helper.Struct2Map(ojUserJSON)
// 	if res, err := ojWebDataValidate.ValidateMap(ojWebUserDataMap, "getBetween"); !res {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
// 		return
// 	}

// 	re := ojWebUserConfigModel.GetUserOJwebUserConfig(ojUserJSON.UserID)
// 	if re.Status != constants.CodeSuccess {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
// 		return
// 	}
// 	cPd := pd.NewApiClient(rpcconn.RPCConn)

// 	ojWebUserConfigRes := re.Data.([]model.OJWebUserConfig)
// 	var userDatas []model.OJWebUserData

// 	// 定义客户端发送的request
// 	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

// 		spiderReq := &pd.GetBetweenRequest{
// 			Targets: []*pd.TargetInfo{
// 				{
// 					UserInfo: &pd.UserInfo{
// 						Name: fmt.Sprint("%d", ojUserJSON.UserID),
// 					},
// 					Oj: []*pd.OJAccount{
// 						{
// 							OjName: ojUserJSON.OJName,
// 							Id:     []string{ojWebUserConfigRe.OJUserName},
// 						},
// 					},
// 				},
// 			},
// 			Start: ojUserJSON.StartTime,
// 			End:   ojUserJSON.EndTime,
// 		}
// 		res, err := cPd.GetBetween(context.Background(), spiderReq)
// 		if err != nil {
// 			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "启动爬虫服务失败", err.Error()))
// 			return
// 		}

// 		for _, data := range res.GetData() {
// 			for key, value := range data.GetData() {
// 				ojWebUserData.OJName = key
// 				ojWebUserData.UserID = ojUserJSON.UserID
// 				solvedData := value.GetData()
// 				for i := range solvedData.Problems {
// 					submitTime, err := time.Parse("2006-01-02 15:04:05 Z0700 MST", solvedData.Problems[i].GetSubmitTime())
// 					if err != nil {
// 						continue
// 					}
// 					ojWebUserData.SubmitTime = submitTime
// 					ojWebUserData.Status = solvedData.Problems[i].GetStatusWord()
// 					ojWebUserData.ProblemID = solvedData.Problems[i].GetProblemTitle()
// 					userDatas = append(userDatas, ojWebUserData)
// 				}
// 			}
// 		}

// 	}
// 	//操作db
// 	re = ojWebUserModel.AddOJWebUserDatas(userDatas)
// 	if re.Status != constants.CodeSuccess {
// 		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "爬取成功", true))
// 	return
// }
