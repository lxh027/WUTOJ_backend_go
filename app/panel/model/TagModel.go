package model

import (
	"OnlineJudge/app/helper"
	"OnlineJudge/constants"
)

type Tag struct {
	ID 		int 	`json:"id" form:"id"`
	Name 	string 	`json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Status 	int 	`json:"status" form:"status"`
	Color 	string 	`json:"color" form:"color"`
}

func (model *Tag) GetAllTag(offset int, limit int, name string) helper.ReturnType {
	var tags []Tag
	where := "name like ?"
	var count int

	db.Model(&Tag{}).Where(where, "%"+name+"%").Count(&count)


	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%").
		Find(&tags).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"tags": tags,
				"count": count,
			},
		}
	}
}

func (model *Tag) GetAvailTag(name string) helper.ReturnType {
	var tags []Tag
	where := "status = 1 AND name like ?"


	err := db.
		Where(where, "%"+name+"%").
		Find(&tags).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: tags}
	}
}

func (model *Tag) FindTagByID(id int) helper.ReturnType {
	var tag Tag

	err := db.Where("id = ?", id).First(&tag).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: tag}
	}
}

func (model *Tag) AddTag(newTag Tag) helper.ReturnType {//jun
	tag :=Tag{}

	if err := db.Where("name = ?", newTag.Name).First(&tag).Error; err == nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "标签已存在",  Data: false}
	}

	err := db.Create(&newTag).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: true}
	}
}

func (model *Tag) DeleteTag(tagID int) helper.ReturnType  {
	err := db.Where("id = ?", tagID).Delete(Tag{}).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "删除失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "删除成功", Data: true}
	}
}

func (model *Tag) UpdateTag(tagID int, updateTag Tag) helper.ReturnType  {
	err := db.Model(&Tag{}).Where("id = ?", tagID).Update(updateTag).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}

func (model *Tag) ChangeTagStatus(tagID int, status int) helper.ReturnType  {
	err := db.Model(&Tag{}).Where("id = ?", tagID).Update("status", status).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "更新失败", Data: false}
	} else {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "更新成功", Data: true}
	}
}