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
	userModel := model.User{}
	var users []model.User
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

			if len(record) != len(title) {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "文件读取出错", false))
				return
			}

			var user model.User

			//csv格式:0:teamID,1:realname,2:school,3:major,4:class,5:contact,6:password

			user.Nick = fmt.Sprintf("c_%d_%s_%s", contestUserJSON.ContestID, record[0], record[1])
			user.Realname = record[1]
			user.School = record[2]
			user.Major = record[3]
			user.Class = record[4]
			user.Contact = record[5]
			user.Password = helper.GetMd5(record[6])
			users = append(users, user)
		}
		res := userModel.AddUsersAndContestUsers(users, contestUserJSON.ContestID)
		if res.Status != constants.CodeSuccess {
			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
			return
		}

	} else {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
		return
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "批量添加成功", true))
	return
}
