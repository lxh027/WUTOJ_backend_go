package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type User struct {
	UserID   int    `json:"user_id" form:"user_id"`
	Nick     string `json:"nick" form:"nick"`
	Password string `json:"password" form:"password"`
	Realname string `json:"realname" form:"realname" gorm:"column:realname"`
	Avatar   string `json:"avatar" form:"avatar"`
	School   string `json:"school" form:"school"`
	Major    string `json:"major" form:"major"`
	Class    string `json:"class" form:"class"`
	Contact  string `json:"contact" form:"contact"`
	Identity uint   `json:"identity" form:"identity"`
	Desc     string `json:"desc" form:"desc"`
	Mail     string `json:"mail" form:"mail"`
	Status   int    `json:"status" form:"status"`
}

//TableName 设定表名
func (User) TableName() string {
	return "users"
}

func (model *User) SetAdmin(userID int, isAdmin int) helper.ReturnType {
	err := db.Model(&User{}).Where("user_id = ?", userID).Update("identity", isAdmin).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *User) UpdateUser(userID int, updateUser User) helper.ReturnType {
	err := db.Model(&User{}).Where("user_id = ?", userID).Update(updateUser).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *User) DeleteUser(userID int) helper.ReturnType {
	err := db.Where("user_id = ?", userID).Delete(User{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *User) AddUser(newUser User) helper.ReturnType {
	user := User{}

	if err := db.Where("nick = ? OR mail = ?", newUser.Nick, newUser.Mail).First(&user).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "昵称或邮箱已存在", Data: false}
	}

	err := db.Create(&newUser).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *User) CheckLogin(loginUser User) helper.ReturnType {
	user := User{}

	if err := db.Where("nick = ? AND password = ?", loginUser.Nick, loginUser.Password).First(&user).Error; err == nil {
		returnData := make(map[string]interface{})
		returnData["userInfo"] = user
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "验证成功", Data: returnData}
	} else {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户名或密码错误", Data: false}
	}
}

func (model *User) GetAllUser(offset int, limit int, nick string, email string) helper.ReturnType {
	var users []User
	where := "nick like ? AND mail like ?"
	var count int

	db.Model(&User{}).Where(where, "%"+nick+"%", "%"+email+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+nick+"%", "%"+email+"%").
		Find(&users).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"users": users,
				"count": count,
			},
		}
	}
}

func (model *User) GetUserByID(userID int) helper.ReturnType {
	var getUser User

	err := db.Select([]string{"nick", "mail"}).Where("user_id = ?", userID).First(&getUser).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: getUser}
	}
}
