package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"fmt"
	"time"

	pd "OnlineJudge/core/grpc/rpc"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

//TODO:搞了一些

const PORT = ":50051"

// var spider map[int]func(pd.ApiClient)

// func init() {
// 	spider = map[int]func(pd.ApiClient){
// 		1:  ,
// 	}
// }

//GetUserAllSubmit 获取用户全部提交记录
func GetUserAllSubmit(c *gin.Context) {
	//客户端连接服务端
	conn, err := grpc.Dial("127.0.0.1"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Panic("network error", err)
	}

	//网络延迟关闭
	defer conn.Close()
	ojWebDataValidate := validate.OJWebDataValidate
	var ojWebUserData model.OJWebUserData
	ojWebUserConfigModel := model.OJWebUserConfig{}

	ojWebUserModel := model.OJWebUserData{}
	ojUserJSON := struct {
		UserID int    `json:"user_id" form:"user_id"`
		OJName string `json:"oj_name" form:"oj_name"`
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

	re := ojWebUserConfigModel.GetUserOJwebUserConfig(ojUserJSON.UserID)
	if re.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
		return
	}
	cPd := pd.NewApiClient(conn)

	ojWebUserConfigRes := re.Data.([]model.OJWebUserConfig)
	var userDatas []model.OJWebUserData

	// 定义客户端发送的request
	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

		spiderReq := &pd.Request{
			Targets: []*pd.TargetInfo{
				{
					UserInfo: &pd.UserInfo{
						Name: fmt.Sprint("%d", ojUserJSON.UserID),
					},
					Oj: []*pd.OJAccount{
						{
							OjName: ojUserJSON.OJName,
							Id:     []string{ojWebUserConfigRe.OJUserName},
						},
					},
				},
			},
		}
		res, err := cPd.GetAll(context.Background(), spiderReq)
		if err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "启动爬虫服务失败", err.Error()))
			return
		}

		for _, data := range res.GetData() {
			for key, value := range data.GetData() {
				ojWebUserData.OJName = key
				ojWebUserData.UserID = ojUserJSON.UserID
				solvedData := value.GetData()
				for i := range solvedData.Problems {
					submitTime, err := time.Parse("2006-01-02 15:04:05 Z0700 MST", solvedData.Problems[i].GetSubmitTime())
					if err != nil {
						continue
					}
					ojWebUserData.SubmitTime = submitTime
					ojWebUserData.Status = solvedData.Problems[i].GetStatusWord()
					ojWebUserData.ProblemID = solvedData.Problems[i].GetProblemTitle()
					userDatas = append(userDatas, ojWebUserData)
				}
			}
		}

	}
	//操作db
	re = ojWebUserModel.AddOJWebUserDatas(userDatas)
	if re.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "爬取成功", true))
	return
}

//GetUserLastWeekSubmit 获取用户全部提交记录
func GetUserLastWeekSubmit(c *gin.Context) {
	//客户端连接服务端
	conn, err := grpc.Dial("127.0.0.1"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Panic("network error", err)
	}

	//网络延迟关闭
	defer conn.Close()
	ojWebDataValidate := validate.OJWebDataValidate
	var ojWebUserData model.OJWebUserData
	ojWebUserConfigModel := model.OJWebUserConfig{}

	ojWebUserModel := model.OJWebUserData{}
	ojUserJSON := struct {
		UserID int    `json:"user_id" form:"user_id"`
		OJName string `json:"oj_name" form:"oj_name"`
	}{}
	if err := c.ShouldBind(&ojUserJSON); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	ojWebUserDataMap := helper.Struct2Map(ojUserJSON)
	if res, err := ojWebDataValidate.ValidateMap(ojWebUserDataMap, "getLastWeek"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	re := ojWebUserConfigModel.GetUserOJwebUserConfig(ojUserJSON.UserID)
	if re.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
		return
	}
	cPd := pd.NewApiClient(conn)

	ojWebUserConfigRes := re.Data.([]model.OJWebUserConfig)
	var userDatas []model.OJWebUserData

	// 定义客户端发送的request
	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

		spiderReq := &pd.Request{
			Targets: []*pd.TargetInfo{
				{
					UserInfo: &pd.UserInfo{
						Name: fmt.Sprint("%d", ojUserJSON.UserID),
					},
					Oj: []*pd.OJAccount{
						{
							OjName: ojUserJSON.OJName,
							Id:     []string{ojWebUserConfigRe.OJUserName},
						},
					},
				},
			},
		}
		res, err := cPd.GetLastWeek(context.Background(), spiderReq)
		if err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "启动爬虫服务失败", err.Error()))
			return
		}

		for _, data := range res.GetData() {
			for key, value := range data.GetData() {
				ojWebUserData.OJName = key
				ojWebUserData.UserID = ojUserJSON.UserID
				solvedData := value.GetData()
				for i := range solvedData.Problems {
					submitTime, err := time.Parse("2006-01-02 15:04:05 Z0700 MST", solvedData.Problems[i].GetSubmitTime())
					if err != nil {
						continue
					}
					ojWebUserData.SubmitTime = submitTime
					ojWebUserData.Status = solvedData.Problems[i].GetStatusWord()
					ojWebUserData.ProblemID = solvedData.Problems[i].GetProblemTitle()
					userDatas = append(userDatas, ojWebUserData)
				}
			}
		}

	}
	//操作db
	re = ojWebUserModel.AddOJWebUserDatas(userDatas)
	if re.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(re.Status, re.Msg, re.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "爬取成功", true))
	return
}
