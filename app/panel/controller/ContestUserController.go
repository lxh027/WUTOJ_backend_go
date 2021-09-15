package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"net/http"
	"encoding/csv"
	"fmt"
	"io"
	"github.com/gin-gonic/gin"
)

//GetAllContestUsers 获取所有参赛选手
func GetAllContestUsers(c *gin.Context) {
	contestUserModel := model.ContestUser{}

	contestuserJSON := struct {
		ContestID int `json:"contest_id" form:"contest_id"`
	}{}

	if c.ShouldBind(&contestuserJSON) == nil {
		res := contestUserModel.GetAllContestUsersByID(contestuserJSON.ContestID)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

//AddContestUsers 通过csv文件批量添加参赛选手
func AddContestUsers(c *gin.Context) {
	//TODO:完全没搞
	// session := sessions.Default(c)

	// contestUserModel := model.ContestUser{}

	// var contestUsers []model.ContestUser
	file, err := c.FormFile("file")

	if err == nil {
		csvFile, err := file.Open()
		if err != nil {
			//TODO:报错
			return
		}
		defer csvFile.Close()
		r := csv.NewReader(csvFile)
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				//TODO:报错
				return
			}
			//TODO:康康record
			fmt.Print(record)
			// fmt.Printf("Record has %d columns.\n", len(record))
			// city, _ := iconv.ConvertString(record[2], "gb2312", "utf-8")
			// fmt.Printf("%s %s %s \n", record[0], record[1], city)
		}
	} else {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	//TODO:delete this later
	// if err := c.ShouldBind(&noticeJson); err != nil {
	// 	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
	// 	return
	// }
	// noticeJson.SubmitTime = time.Now()
	// noticeJson.ModifyTime = time.Now()
	// noticeJson.UserID = session.Get("userId").(int)

	// res := noticeModel.AddNotification(noticeJson)

	//TODO:
	// c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}