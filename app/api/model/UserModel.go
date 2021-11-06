package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
	"os"
)

type User struct {
	UserID   uint   `json:"user_id" form:"user_id"`
	Nick     string `json:"nick" form:"nick"`
	Password string `json:"password,omitempty" form:"password"`
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

// 设定表名
func (User) TableName() string {
	return "users"
}

// 添加用户
func (model *User) AddUser(data User) helper.ReturnType {
	user := User{}
	// 判断昵称是否已存在
	err := db.
		Where("nick = ?", data.Nick).
		First(&user).
		Error
	if err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "昵称已存在", Data: user}
	}
	// 创建记录
	err = db.Create(&data).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: 1}
	}
}

// 通过ID编辑用户
func (model *User) EditUserByID(userId uint, data User) helper.ReturnType {
	err := db.
		Select([]string{"user_id", "nick", "realname", "school", "major", "class", "contact"}).
		Model(&data).
		Where("user_id = ?", userId).
		Update(data).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: 1}
	}
}

// 通过nick编辑用户
// MARK: 此接口没有被使用过
func (model *User) EditUserByNick(nick string, data User) helper.ReturnType {
	err := db.Model(model).Where("nick = ?", nick).Update(data).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: 1}
	}
}

// 通过ID查询用户
// MARK:此处使用该方法查询用户身份
func (model *User) FindUserByID(userID uint) helper.ReturnType {
	user := User{}
	err := db.Model(&User{}).
		Select([]string{"user_id", "nick", "realname", "avatar", "school", "major", "class", "contact", "identity"}).
		Where("user_id = ?", userID).
		First(&user).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: user}
	}
}

// 通过nick查询用户
// MARK: 此接口没有被使用过
func (model *User) FindUserByNick(nick string) helper.ReturnType {
	user := User{}
	err := db.Where("nick = ?", nick).First(&user).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: user}
	}
}

func (model *User) LoginCheck(data User) helper.ReturnType {
	user := User{}
	err := db.
		Select([]string{"user_id", "nick", "password"}).
		Where("nick = ? AND password = ?", data.Nick, data.Password).
		First(&user).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "用户名或密码错误", Data: err.Error()}
	} else {
		userSubmitLog := UserSubmitLog{}
		res := userSubmitLog.GetUserSubmitLog(user.UserID)

		resp := make(map[string]interface{})
		resp["userInfo"] = user
		resp["submitLog"] = res.Data
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "登录验证成功", Data: resp}
	}
}

func (model *User) UpdatePassword(user User) helper.ReturnType {

	err := db.
		Model(&user).
		Select([]string{"mail", "password"}).
		Where("mail = ?", user.Mail).
		Update("password", user.Password).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "修改密码失败", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "修改密码成功", Data: ""}
}

func (model *User) SearchUser(param string) helper.ReturnType {
	user := User{}
	err := db.
		Model(&User{}).
		Select([]string{"user_id", "nick", "realname", "avatar", "school", "major", "class", "contact", "identity"}).
		Where("user_id = ?", param).
		Find(&user).Error

	if err == nil {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: user}
	}
	return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}

}

func (model *User) AddUserAvatar(UserID int, avatar string) helper.ReturnType {
	user := User{}
	err := db.
		Select([]string{"user_id", "avatar"}).
		Where("user_id = ?", UserID).
		First(&user).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取用户信息失败", Data: err.Error()}
	}

	if user.Avatar != "null" {
		err := os.Remove(user.Avatar)
		if err != nil {
			return helper.ReturnType{Status: constants.CodeError, Msg: "移除原头像失败", Data: err.Error()}
		}
	}
	user.Avatar = avatar

	err = db.
		Select([]string{"user_id", "avatar"}).
		Model(&user).
		Where("user_id = ?", UserID).
		Update(&user).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新用户信息失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "添加头像成功", Data: ""}

}


func (model *User) GetContestUser(contestID int, isStar int) helper.ReturnType {
	var users []User

	fields := []string{"users.user_id", "nick", "realname", "school"}

	err := db.
		Joins("JOIN contest_users ON users.user_id = contest_users.user_id").
		Where("contest_users.contest_id = ? AND contest_users.status = ?", contestID, isStar).
		Select(fields).
		Find(&users).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: users}
}