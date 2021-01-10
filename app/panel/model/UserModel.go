package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
)

type User struct {
	Uid 	int 	`json:"uid" form:"uid"`
	Nick	string 	`json:"nick" form:"nick"`
	Password string	`json:"password" form:"password"`
	Mail 	string 	`json:"mail" form:"mail"`
	IsAdmin	int 	`json:"is_admin" form:"is_admin"`
}

func (model *User) SetAdmin(uid int, isAdmin int) helper.ReturnType {
	err := db.Model(&User{}).Where("uid = ?", uid).Update("is_admin", isAdmin).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *User) UpdateUser(userID int, updateUser User) helper.ReturnType  {
	err := db.Model(&User{}).Where("uid = ?", userID).Update(updateUser).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *User) DeleteUser(userID int) helper.ReturnType  {
	err := db.Where("uid = ?", userID).Delete(User{}).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *User) AddUser(newUser User) helper.ReturnType {
	user :=User{}

	if err := db.Where("nick = ? OR mail = ?", newUser.Nick, newUser.Mail).First(&user).Error; err == nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "昵称或邮箱已存在",  Data: false}
	}

	err := db.Create(&newUser).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *User) CheckLogin(loginUser User) helper.ReturnType {
	user := User{}

	if err := db.Where("nick = ? AND password = ?", loginUser.Nick, loginUser.Password).First(&user).Error; err == nil {
		returnData := make(map[string]interface{})
		returnData["userInfo"] = user
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "验证成功", Data: returnData}
	} else {
		return helper.ReturnType{Status: common.CodeError, Msg: "用户名或密码错误", Data: false}
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
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"users": users,
				"count": count,
			},
		}
	}
}

func (model *User) GetUserByID(uid int) helper.ReturnType {
	var getUser User

	err := db.Select([]string{"nick", "mail"}).Where("uid = ?", uid).First(&getUser).Error
	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: getUser}
	}
}