package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUserSubmitStatus(c *gin.Context) {
	if res := haveAuth(c, "getUserSubmit"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userLogModel := model.UserSubmitLog{}

	userLogJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Nick string `json:"nick" form:"nick"`
		}
	}{}

	if c.ShouldBind(&userLogJson) == nil {
		userLogJson.Offset = (userLogJson.Offset - 1) * userLogJson.Limit
		res := userLogModel.GetAllUserSubmitStatus(userLogJson.Offset, userLogJson.Limit, userLogJson.Where.Nick)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}
