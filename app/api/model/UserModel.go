package model

import (
	"errors"
)

type User struct {
	UserID 		uint 	`json:"user_id" form:"user_id"`
	Nick   		string 	`json:"nick" form:"nick"`
	Password	string 	`json:"password" form:"password"`
	RealName	string	`json:"realname" form:"realname" gorm:"column:realname"`
	Avatar		string 	`json:"avatar" form:"avatar"`
	School		string	`json:"school" form:"school"`
	Major		string	`json:"major" form:"major"`
	Class       string  `json:"class" form:"class"`
	Contact		string	`json:"contact" form:"contact"`
	Identity 	uint	`json:"identity" form:"identity"`
	Desc    	string 	`json:"desc" form:"desc"`
	Mail 		string 	`json:"mail" form:"mail"`
	Status 		int 	`json:"status" form:"status"`
	RoleGroup	string 	`json:"role_group" form:"role_group"`
}

// 设定表名
func (User) TableName() string {
	return "users"
}

// 添加用户
func (model *User) AddUser(data User)(bool, error)  {
	err := db.Create(&data).Error
	if err != nil {
		return false, err
	} else {
		return true, errors.New("创建成功")
	}

}

