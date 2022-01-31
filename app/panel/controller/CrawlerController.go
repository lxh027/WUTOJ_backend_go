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

//GetUserSubmit 发送爬取请求
func GetUserSubmit(c *gin.Context) {

	ojWebDataValidate := validate.OJWebDataValidate
	ojWebUserConfigModel := model.OJWebUserConfig{}

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
	// 定义客户端发送的request
	for _, ojWebUserConfigRe := range ojWebUserConfigRes {

		spiderReq := &nsqueue.Request{
			Targets: []*nsqueue.TargetInfo{
				{
					UserInfo: fmt.Sprintf("%d", ojUserJSON.UserID),
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
