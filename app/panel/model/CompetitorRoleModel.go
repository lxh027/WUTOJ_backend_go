package model

// //自建
// import (
// 	"OnlineJudge/app/helper"
// 	"OnlineJudge/constants"
// )

// type CompetitorRole struct {
// 	Rid  int    `json:"rid" form:"rid"`
// 	Desc string `json:"desc" form:"desc"`
// 	Name string `json:"name" form:"name"`
// }

// func (model *CompetitorRole) GetAllCompetitorRole(offset int, limit int, name string, desc string) helper.ReturnType {
// 	var competitorRoles []CompetitorRole
// 	where := "name like ? AND `desc` like ?"
// 	var count int

// 	db.Model(&CompetitorRole{}).Where(where, "%"+name+"%", "%"+desc+"%").Count(&count)

// 	err := db.Offset(offset).
// 		Limit(limit).
// 		Where(where, "%"+name+"%", "%"+desc+"%").
// 		Find(&competitorRoles).
// 		Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
// 			Data: map[string]interface{}{
// 				"roles": competitorRoles,
// 				"count": count,
// 			},
// 		}
// 	}
// }

// func (model *CompetitorRole) GetContestCompetitor(contestID int) helper.ReturnType {
// 	var competitorRoles []CompetitorRole

// 	err := db.
// 		Joins("JOIN contest_competitor ON competitor_role.rid = contest_competitor.rid AND contest_competitor.contest_id = ? ", contestID).
// 		Find(&competitorRoles).
// 		Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: competitorRoles}
// 	}
// }

// func (model *CompetitorRole) GetCompetitorRoleNoRules() helper.ReturnType {
// 	/*var roles, rolesTotal []Role

// 	var countTotal, countRole int

// 	err1 := db.Order("rid").Find(&rolesTotal).
// 		Count(&countTotal).Error

// 	err2 := db.Joins("JOIN user_role ON role.rid = user_role.rid AND user_role.uid = ? ", uid).
// 		Order("rid").
// 		Find(&roles).
// 		Count(&countRole).Error

// 	countLeft := countTotal-countRole
// 	var rolesLeft []Role
// 	j := 0
// 	for i := 0; i < countRole; i++ {
// 		if roles[i].Rid == rolesTotal[j].Rid {
// 			j++
// 			continue
// 		}
// 		for roles[i].Rid != rolesTotal[j].Rid {
// 			rolesLeft = append(rolesLeft, rolesTotal[j])
// 			j++
// 		}
// 	}
// 	for j < countTotal {
// 		rolesLeft = append(rolesLeft, rolesTotal[j])
// 		j++
// 	}*/
// 	var competitorRolesTotal []CompetitorRole

// 	err := db.Find(&competitorRolesTotal).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: competitorRolesTotal}
// 	}
// }

// func (model *CompetitorRole) GetCompetitorRoleByID(rid int) helper.ReturnType { //jun
// 	var getCompetitorRole CompetitorRole

// 	err := db.Select([]string{"name", "`desc`"}).Where("rid = ?", rid).First(&getCompetitorRole).Error
// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: getCompetitorRole}
// 	}
// }

// func (model *CompetitorRole) AddCompetitorRole(newCompetitorRole CompetitorRole) helper.ReturnType { //jun
// 	competitorRole := CompetitorRole{}

// 	if err := db.Where("name = ? OR `desc` = ?", newCompetitorRole.Name, newCompetitorRole.Desc).First(&competitorRole).Error; err == nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "角色名或描述已存在", Data: false}
// 	}

// 	err := db.Create(&newCompetitorRole).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
// 	}
// }

// func (model *CompetitorRole) DeleteCompetitorRole(roleID int) helper.ReturnType {
// 	err := db.Where("rid = ?", roleID).Delete(CompetitorRole{}).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
// 	}
// }

// func (model *CompetitorRole) UpdateCompetitorRole(roleID int, updateCompetitorRole CompetitorRole) helper.ReturnType {
// 	err := db.Model(&CompetitorRole{}).Where("rid = ?", roleID).Update(updateCompetitorRole).Error

// 	if err != nil {
// 		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
// 	} else {
// 		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
// 	}
// }
