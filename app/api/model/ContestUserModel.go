package model

import "OnlineJudge/app/common"

type ContestUser struct {
	ContestID int `json:"contest_id" form:"contest_id"`
	UserID    int `json:"user_id" form:"user_id"`
}

// 设定表名
func (ContestUser) TableName() string {
	return "contest_users"
}

func searchUser(ContestID int, UserID int) common.ReturnType {
	contestUser := ContestUser{}
	err := db.Where("contest_id = ?", ContestID).Where("user_id = ?", UserID).First(contestUser).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "数据库错误"}
	} else {
		if contestUser != (ContestUser{}) {
			return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "已参加比赛", Data: ""}
		} else {
			return common.ReturnType{Status: common.CODE_ERROE, Msg: "未参加比赛", Data: ""}
		}
	}
}

func searchUserContest(UserID int) common.ReturnType {
	contestUser := ContestUser{}
	err := db.Where("user_id = ?", UserID).First(&contestUser).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "数据库错误"}
	} else {
		if contestUser != (ContestUser{}) {
			return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: contestUser}
		} else {
			return common.ReturnType{Status: common.CODE_ERROE, Msg: "无比赛信息", Data: ""}
		}
	}
}

func addInfo(ContestID int, UserID int) common.ReturnType {
	contestUser := ContestUser{ContestID: ContestID, UserID: UserID}
	err := db.Create(&contestUser)
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "参加失败，数据库错误"}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "参加成功", Data: ""}
	}
}
