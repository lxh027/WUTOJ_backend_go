package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTag(c *gin.Context)  {
	if res := haveAuth(c, "getAllTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	tagModel := model.Tag{}

	tagJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Name 	string 	`json:"name" form:"name"`
		}
	}{}

	if c.ShouldBind(&tagJson) == nil {
		tagJson.Offset = (tagJson.Offset-1)*tagJson.Limit
		res := tagModel.GetAllTag(tagJson.Offset, tagJson.Limit, tagJson.Where.Name)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func FindTagsByName(c *gin.Context) {
	if res := haveAuth(c, "getAllTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	//tagValidate := validate.TagValidate
	tagModel := model.Tag{}

	var tagJson model.Tag

	if err := c.ShouldBind(&tagJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	/*tagMap := helper.Struct2Map(tagJson)
	if res, err:= tagValidate.ValidateMap(tagMap, "findByName"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}*/

	res := tagModel.GetAvailTag(tagJson.Name)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetTagByID(c *gin.Context) {
	if res := haveAuth(c, "getAllTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	tagValidate := validate.TagValidate
	tagModel := model.Tag{}

	var tagJson model.Tag

	if err := c.ShouldBind(&tagJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	tagMap := helper.Struct2Map(tagJson)
	if res, err:= tagValidate.ValidateMap(tagMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := tagModel.FindTagByID(tagJson.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddTag(c *gin.Context) {
	if res := haveAuth(c, "addTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	tagValidate := validate.TagValidate
	tagModel := model.Tag{}

	var tagJson model.Tag
	if err := c.ShouldBind(&tagJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	tagMap := helper.Struct2Map(tagJson)
	if res, err:= tagValidate.ValidateMap(tagMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := tagModel.AddTag(tagJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteTag(c *gin.Context) {
	if res := haveAuth(c, "deleteTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	tagValidate := validate.TagValidate
	tagModel := model.Tag{}

	var tagJson model.Tag
	if err := c.ShouldBind(&tagJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	tagMap := helper.Struct2Map(tagJson)
	if res, err:= tagValidate.ValidateMap(tagMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := tagModel.DeleteTag(tagJson.ID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateTag(c *gin.Context) {
	if res := haveAuth(c, "updateTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	tagValidate := validate.TagValidate
	tagModel := model.Tag{}

	var tagJson model.Tag
	if err := c.ShouldBind(&tagJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	tagMap := helper.Struct2Map(tagJson)
	if res, err:= tagValidate.ValidateMap(tagMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := tagModel.UpdateTag(tagJson.ID, tagJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeTagStatus(c *gin.Context) {
	if res := haveAuth(c, "updateTag"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	tagValidate := validate.TagValidate
	tagModel := model.Tag{}

	var tagJson model.Tag
	if err := c.ShouldBind(&tagJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	tagMap := helper.Struct2Map(tagJson)
	if res, err:= tagValidate.ValidateMap(tagMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := tagModel.ChangeTagStatus(tagJson.ID, tagJson.Status)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}