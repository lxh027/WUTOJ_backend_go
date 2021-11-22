package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"errors"
	"fmt"
	"net/http"
	"time"

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

func GetUserNickFromSession(c *gin.Context) string {
	session := sessions.Default(c)
	if nick := session.Get("nick"); nick != "" {
		return nick.(string)
	}
	return ""
}

func Check(c *gin.Context) {
	res := checkLogin(c)
	if res.Status == constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, gin.H{
			"user_id": res.Data,
		}))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
}

func checkLogin(c *gin.Context) helper.ReturnType {
	//return helper.ReturnType{Status: common.CodeSuccess, Msg: "已登陆", Data: 0}
	session := sessions.Default(c)
	var id interface{}
	if id = session.Get("user_id"); id == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "未登录，请先登录", Data: -1}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "已登陆", Data: id}
}

// return begin, frozen, end
func getContestTime(contestID uint) (time.Time, time.Time, time.Time, error) {
	contestModel := model.Contest{}
	res := contestModel.GetContestById(int(contestID))
	now := time.Now()
	if res.Status != constants.CodeSuccess {
		return now, now, now, errors.New(res.Msg)
	}
	beginTime := res.Data.(model.Contest).BeginTime
	endTime := res.Data.(model.Contest).EndTime
	frozen := res.Data.(model.Contest).Frozen
	frozenTime := time.Unix(int64((float64(endTime.Unix())-float64(beginTime.Unix()))*(1-frozen)+float64(beginTime.Unix())), 0)
	frozenTime = frozenTime.UTC()
	//format := "2006-01-02 15:04:05"
	//frozenF, _ := time.Parse(format, frozenTime.Format(format))
	fmt.Printf("beginTime: %v, endTime: %v, frozenTime: %v, frozenF: %v",
		beginTime, endTime, frozenTime)
	return beginTime, endTime, frozenTime, nil
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
