package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Role struct {
	Rid 	int 	`json:"rid" form:"rid"`
	Desc 	string 	`json:"desc" form:"desc"`
	Name 	string 	`json:"name" form:"name"`
}





func (model *Role) GetAllRole(offset int, limit int, name string, desc string) helper.ReturnType {
	var roles []Role
	where := "name like ? AND `desc` like ?"
	var count int

	db.Model(&Role{}).Where(where, "%"+name+"%", "%"+desc+"%").Count(&count)


	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%", "%"+desc+"%").
		Find(&roles).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"roles": roles,
				"count": count,
			},
		}
	}
}

func (model *Role) GetUserRole(userID int) helper.ReturnType {
	var roles []Role

	err := db.
		Joins("JOIN user_role ON role.rid = user_role.rid AND user_role.user_id = ? ", userID).
		Find(&roles).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: roles,
		}
	}
}

func (model *Role) GetRoleNoRules() helper.ReturnType {
	/*var roles, rolesTotal []Role

	var countTotal, countRole int

	err1 := database.Order("rid").Find(&rolesTotal).
		Count(&countTotal).Error

	err2 := database.Joins("JOIN user_role ON role.rid = user_role.rid AND user_role.uid = ? ", uid).
		Order("rid").
		Find(&roles).
		Count(&countRole).Error

	countLeft := countTotal-countRole
	var rolesLeft []Role
	j := 0
	for i := 0; i < countRole; i++ {
		if roles[i].Rid == rolesTotal[j].Rid {
			j++
			continue
		}
		for roles[i].Rid != rolesTotal[j].Rid {
			rolesLeft = append(rolesLeft, rolesTotal[j])
			j++
		}
	}
	for j < countTotal {
		rolesLeft = append(rolesLeft, rolesTotal[j])
		j++
	}*/
	var rolesTotal []Role

	err := db.Find(&rolesTotal).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: rolesTotal,
		}
	}
}

func (model *Role) GetRoleByID(rid int) helper.ReturnType {//jun
	var getRole Role

	err := db.Select([]string{"name", "`desc`"}).Where("rid = ?", rid).First(&getRole).Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: getRole}
	}
}

func (model *Role) AddRole(newRole Role) helper.ReturnType {//jun
	role :=Role{}

	if err := db.Where("name = ? OR `desc` = ?", newRole.Name,newRole.Desc).First(&role).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "角色名或描述已存在",  Data: false}
	}

	err := db.Create(&newRole).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *Role) DeleteRole(roleID int) helper.ReturnType  {
	err := db.Where("rid = ?", roleID).Delete(Role{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Role) UpdateRole(roleID int, updateRole Role) helper.ReturnType  {
	err := db.Model(&Role{}).Where("rid = ?", roleID).Update(updateRole).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}