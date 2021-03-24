package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromSession(c *gin.Context) uint {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		return id.(uint)
	}
	return 0
}

func CheckLogin(c *gin.Context) helper.ReturnType {
	//return helper.ReturnType{Status: common.CodeSuccess, Msg: "已登陆", Data: 0}
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "未登录，请先登录", Data: 1}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "已登陆", Data: 0}
}

//func Upload(FileDst string, file *multipart.FileHeader) helper.ReturnType {
//
//	FileNameMd5 := helper.GetMd5(file.Filename)
//
//	dst := "../uploads/image/" + FileNameMd5 + path.Ext(file.Filename)
//
//	UserID := middleware.GetUserIdFromSession(c)
//	userModel := model.User{}
//
//	if err := c.SaveUploadedFile(file, dst); err != nil {
//		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "上传文件失败", err.Error()))
//		return
//	}
//
//	res = userModel.AddUserAvatar(int(UserID), dst)
//	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
//	return
//}
