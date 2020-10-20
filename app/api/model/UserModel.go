package model

import (
	"OnlineJudge/app/common"
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
	AllProblem	interface{}	`json:"all_problem" form:"all_problem"`
}

// 设定表名
func (User) TableName() string {
	return "users"
}

// 添加用户
func (model *User) AddUser(data User)common.ReturnType  {
	user := User{}
	// 判断昵称是否已存在
	if err := db.Where("nick = ?", data.Nick).First(user).Error; err == nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "昵称已存在",  Data: user}
	}
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

// 通过ID编辑用户
func (model *User) EditUserByID(userId uint, data User)common.ReturnType{
	err := db.Model(model).Where("user_id = ?", userId).Update(data).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "更新失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

// 通过nick编辑用户
func (model *User) EditUserByNick( nick string, data User)common.ReturnType{
	err := db.Model(model).Where("nick = ?", nick).Update(data).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "更新失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

// 通过ID查询用户
func (model *User) FindUserByID(userID uint)common.ReturnType  {
	user := User{}
	err := db.Where("user_id = ?", userID).First(user).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: user}
	}
}

// 通过nick查询用户
func (model *User) FindUserByNick(nick string)common.ReturnType  {
	user := User{}
	err := db.Where("nick = ?", nick).First(user).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: user}
	}
}

func (model *User) LoginCheck(data User)common.ReturnType  {
	user := User{}
	err := db.Where("nick = ? AND password = ?", data.Nick, data.Password).First(&user).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		submitModel := Submit{}
		res := submitModel.GetUserSubmits(data.UserID)
		if res.Status != common.CODE_SUCCESS {
			return res
		} else {
			user.AllProblem = res.Data
			return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "登录验证成功", Data: user}
		}
	}
}



