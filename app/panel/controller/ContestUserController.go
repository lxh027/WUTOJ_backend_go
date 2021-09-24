package controller

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetAllContestUsers 获取所有参赛选手
func GetAllContestUsers(c *gin.Context) {
	contestUserModel := model.ContestUser{}

	contestUserJSON := struct {
		ContestID int `json:"contest_id" form:"contest_id"`
	}{}

	if c.ShouldBind(&contestUserJSON) != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}
	res := contestUserModel.GetAllContestUsersByID(contestUserJSON.ContestID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

//AddContestUsers 通过csv文件批量添加参赛选手
func AddContestUsers(c *gin.Context) {
	contestUserModel := model.ContestUser{}
	userModel := model.User{}

	contestUserJSON := struct {
		ContestID int `json:"contest_id" form:"contest_id"`
	}{}

	err := c.ShouldBind(&contestUserJSON)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}
	//这里动文件
	file, err := c.FormFile("file")

	if err == nil {
		csvFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "文件打开失败", false))
			return
		}
		defer csvFile.Close()
		r := csv.NewReader(csvFile)
		title, err := r.Read()
		if err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "文件读取失败", false))
			return
		}
		for {
			// Read each record from csv
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "文件读取失败", false))
				return
			}

			//用于康康record
			// fmt.Print(record)
			// fmt.Printf("Record has %d columns.\n", len(record))
			// city, _ := iconv.ConvertString(record[2], "gb2312", "utf-8")
			// fmt.Printf("%s %s %s \n", record[0], record[1], city)

			if len(record) != len(title) {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "文件读取出错", false))
				return
			}
			var user model.User
			//csv格式:0:teamID,1:realname,2:school,3:major,4:class,5:contact
			user.Nick = "c_" + fmt.Sprintf("%d", contestUserJSON.ContestID) + "_" + record[0]
			user.Realname = record[1]
			user.School = record[2]
			user.Major = record[3]
			user.Class = record[4]
			user.Contact = record[5]
			res := userModel.AddUserWithoutCheckEMail(user)
			if res.Msg != "创建成功" {
				c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
				return
			}
			var contestUser model.ContestUser
			// var findUser model.User
			contestUser.ContestID = contestUserJSON.ContestID
			findUser := userModel.GetUserByNick(user.Nick)
			user_to_id, ok := (findUser.Data).(model.User)
			if !ok {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "数据库错误", false))
			}
			contestUser.UserID = user_to_id.UserID
			res = contestUserModel.AddContestUser(contestUser)
			if res.Msg != "参加比赛成功" {
				c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
				return
			}
		}
	} else {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "批量添加成功", true))
	return
}
