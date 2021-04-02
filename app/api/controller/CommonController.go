package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetUserIdFromSession(c *gin.Context) uint {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		return id.(uint)
	}
	return 0
}

func GetUserNickFromSession(c *gin.Context) string {
	session := sessions.Default(c)
	if nick := session.Get("nick"); nick != "" {
		return nick.(string)
	}
	return ""
}

func Check(c *gin.Context) {
	res := checkLogin(c)
	if res.Data == 0 {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func checkLogin(c *gin.Context) helper.ReturnType {
	//return helper.ReturnType{Status: common.CodeSuccess, Msg: "已登陆", Data: 0}
	session := sessions.Default(c)
	if id := session.Get("user_id"); id == nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "未登录，请先登录", Data: 1}
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "已登陆", Data: 0}
}

// return begin, frozen, end
func getContestTime(contestID uint) (time.Time, time.Time, time.Time, error) {
	contestModel := model.Contest{}
	res := contestModel.GetContestById(strconv.FormatInt(int64(contestID), 10))
	now := time.Now()
	if res.Status != common.CodeSuccess {
		return now, now, now, errors.New(res.Msg)
	}
	beginTime := res.Data.(model.Contest).BeginTime
	endTime := res.Data.(model.Contest).EndTime
	frozen := res.Data.(model.Contest).Frozen
	frozenTime := time.Unix(int64(float64(endTime.Unix())*frozen+float64(beginTime.Unix())), 0)

	return beginTime, frozenTime, endTime, nil
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
