package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

//
//func UploadImage(c *gin.Context) {
//
//	res := CheckLogin(c)
//	if res.Status == common.CodeError {
//		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
//		return
//	}
//
//	file,err := c.FormFile("image")
//	if err != nil {
//		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "获取文件失败", err.Error()))
//		return
//	}
//
//	FileNameMd5 := helper.GetMd5(file.Filename)
//
//	dst := "../uploads/image/" + FileNameMd5 + path.Ext(file.Filename)
//
//	if err := c.SaveUploadedFile(file, dst); err != nil {
//		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "上传文件失败", err.Error()))
//		return
//	}
//
//	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "上传图片成功",""))
//	return
//}

func UploadAvatar(c *gin.Context) {

	res := checkLogin(c)
	if res.Status == constants.CodeError {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "获取文件失败", err.Error()))
		return
	}

	FileNameMd5 := helper.GetMd5(file.Filename)

	dst := "../uploads/image/" + FileNameMd5 + path.Ext(file.Filename)

	UserID := GetUserIdFromSession(c)
	userModel := model.User{}

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(constants.CodeError, "上传文件失败", err.Error()))
		return
	}

	res = userModel.AddUserAvatar(int(UserID), dst)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}
