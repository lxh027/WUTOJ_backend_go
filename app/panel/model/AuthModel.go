package model

import "OnlineJudge/app/common"
import "OnlineJudge/app/helper"

type Auth struct {
	Aid 	int 	`json:"aid" form:"aid"`
	Icon 	string 	`json:"icon" form:"icon"`
	Title	string 	`json:"title" form:"title"`
	Href	string 	`json:"href" form:"href"`
	Target 	string 	`json:"target" form:"target"`
	Type 	int 	`json:"type" form:"type"`
	Parent	int 	`json:"parent" form:"parent"`
}

func (model *Auth) GetUserAllAuth(userID int) helper.ReturnType  {
	var auths []Auth

	db.Joins("JOIN role_auth ON auth.aid = role_auth.aid").
		Joins("JOIN user_role ON role_auth.rid = user_role.rid AND user_role.user_id = ?", userID).
		Find(&auths)

	var returnAuths []Auth
	isExist := make(map[int]bool)
	for _, auth := range auths {
		if _, ok := isExist[auth.Aid]; !ok {
			returnAuths = append(returnAuths, auth)
			isExist[auth.Aid] = true
		}
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "OK", Data: returnAuths}
}

func (model *Auth) GetAllAuth(offset int, limit int, title string) helper.ReturnType {
	var auths []Auth
	where := "title like ?"
	var count int

	db.Model(&Auth{}).Where(where, "%"+title+"%").Count(&count)

	err := db.
		Where(where, "%"+title+"%").
		Limit(limit).Offset(offset).
		Find(&auths).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"auths": auths,
				"count": count,
			},
		}
	}
}

func (model *Auth) GetParentAuth(tp int) helper.ReturnType {
	var auths []Auth

	var err error
	if tp == 2 || tp == 0 {
		err = db.Where("type = ?", tp-1).Find(&auths).Error
	} else {
		err = db.Where("type = ? || type = ?", 0, 1).Find(&auths).Error
	}


	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: auths}
	}
}

func (model *Auth) AddAuth(newAuth Auth) helper.ReturnType {
	auth :=Auth{}

	if err := db.Where("type = ? AND title = ?", newAuth.Type,newAuth.Title).First(&auth).Error; err == nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "此类型权限名已存在",  Data: false}
	}

	var err error
	if newAuth.Type == 0 {
		sqlLine := "INSERT INTO auth (icon, title, href, target, type) VALUES (?, ?, ?, ?, ?)"
		err = db.Exec(sqlLine, newAuth.Icon, newAuth.Title, newAuth.Href, newAuth.Target, newAuth.Type).Error
	} else {
		err = db.Create(&newAuth).Error
	}


	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "创建成功", Data: true}
	}

}

func (model *Auth) DeleteAuth(authID int) helper.ReturnType  {
	err := db.Where("aid = ?", authID).Delete(Auth{}).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Auth) UpdateAuth(authID int, updateAuth Auth) helper.ReturnType  {
	fields := [][]string {
		{"title", "icon"},
		{"title", "icon", "href", "parent"},
		{"title", "icon", "parent"},
	}
	err := db.Model(&Auth{}).Select(fields[updateAuth.Type]).Where("aid = ?", authID).Update(updateAuth).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Auth) GetRoleAuth(rid int) helper.ReturnType {
	var auths []Auth

	err := db.
		Joins("JOIN role_auth ON auth.aid = role_auth.aid AND role_auth.rid = ? ", rid).
		Find(&auths).
		Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: auths,
		}
	}
}

func (model *Auth) GetAuthNoRules() helper.ReturnType {
	var authsTotal []Auth

	err := db.Find(&authsTotal).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: authsTotal,
		}
	}
}

func (model *Auth) GetAuthByID(aid int) helper.ReturnType {//jun
	var getAuth Auth

	var err error

	err = db.Select([]string {"icon", "title", "type", "href", "parent"}).
		Where("aid = ?", aid).First(&getAuth).Error

	if err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功", Data: getAuth}
	}
}