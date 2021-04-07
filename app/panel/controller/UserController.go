package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateUser(c *gin.Context) {
	if res := haveAuth(c, "updateUser"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	var userJson model.User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "updateUser"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := userModel.UpdateUser(userJson.UserID, userJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteUsers(c *gin.Context) {
	if res := haveAuth(c, "deleteUser"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	userArrayJson := struct {
		Users []int `json:"users" form:"users"`
	}{}

	if err := c.ShouldBind(&userArrayJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", false))
		return
	}

	userArrayMap := helper.Struct2Map(userArrayJson)
	if res, err := userValidate.ValidateMap(userArrayMap, "groupDelete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	for _, userID := range userArrayJson.Users {
		res := userModel.DeleteUser(userID)
		if res.Status != common.CodeSuccess {
			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, "user_id为"+string(strconv.Itoa(userID))+"的用户删除失败", res.Data))
			return
		}
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeSuccess, "删除成功", true))
	return
}

func SetUserAdmin(c *gin.Context) {
	if res := haveAuth(c, "roleAssign"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	userIDJson := struct {
		UserID   int `json:"user_id" form:"user_id"`
		Identity int `json:"identity" form:"identity"`
	}{}

	if err := c.ShouldBind(&userIDJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userIDMap := helper.Struct2Map(userIDJson)
	if res, err := userValidate.ValidateMap(userIDMap, "updateUser"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := userModel.SetAdmin(userIDJson.UserID, userIDJson.Identity)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteUser(c *gin.Context) {
	if res := haveAuth(c, "deleteUser"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	userIDJson := struct {
		UserID int `json:"user_id" form:"user_id"`
	}{}

	if err := c.ShouldBind(&userIDJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userIDMap := helper.Struct2Map(userIDJson)
	if res, err := userValidate.ValidateMap(userIDMap, "deleteUser"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := userModel.DeleteUser(userIDJson.UserID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetAllUser(c *gin.Context) {
	if res := haveAuth(c, "getAllUser"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userModel := model.User{}

	userJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Nick string `json:"nick" form:"nick"`
			Mail string `json:"mail" form:"mail"`
		}
	}{}

	if c.ShouldBind(&userJson) == nil {
		userJson.Offset = (userJson.Offset - 1) * userJson.Limit
		res := userModel.GetAllUser(userJson.Offset, userJson.Limit, userJson.Where.Nick, userJson.Where.Mail)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetUserByID(c *gin.Context) {
	if res := haveAuth(c, "getAllUser"); res != common.Authed {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "权限不足", res))
		return
	}
	userValidate := validate.UserValidate
	userModel := model.User{}

	var userJson model.User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "searchUser_id"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := userModel.GetUserByID(userJson.UserID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func Register(c *gin.Context) {
	userValidate := validate.UserValidate
	userModel := model.User{}

	var userJson struct {
		model.User
		PasswordCheck string `json:"password_check" form:"password_check"`
	}
	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userMap := helper.Struct2Map(userJson)
	if res, err := userValidate.ValidateMap(userMap, "register"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	if userJson.Password != userJson.PasswordCheck {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "两次密码输入不一致", false))
		return
	}

	userJson.Password = helper.GetMd5(userJson.Password)
	res := userModel.AddUser(userJson.User)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	if id := session.Get("userId"); id != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeSuccess, "已登陆", data))
		return
	}

	userValidate := validate.UserValidate
	userModel := model.User{}

	var loginUser model.User

	if err := c.ShouldBind(&loginUser); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	userMap := helper.Struct2Map(loginUser)
	if res, err := userValidate.ValidateMap(userMap, "login"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	loginUser.Password = helper.GetMd5(loginUser.Password)
	res := userModel.CheckLogin(loginUser)
	if res.Status == common.CodeSuccess {
		userInfo := res.Data.(map[string]interface{})["userInfo"].(model.User)
		returnData := map[string]interface{}{
			"user_id":  userInfo.UserID,
			"nick":     userInfo.Nick,
			"identity": userInfo.Identity,
		}
		/*if menu, auths, err := getUserAllAuth(userInfo.UserID); err == nil {
			returnData["auths"] = auths
			returnData["menu"] = menu
		} else {
			c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "获取权限失败", err.Error()))
			return
		}*/
		jsonData, _ := json.Marshal(returnData)
		session.Set("userId", returnData["user_id"])
		session.Set("identity", returnData["identity"])
		session.Set("data", string(jsonData))
		if err := session.Save(); err == nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, "登录成功", returnData))
		} else {
			c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "登录失败", err.Error()))
		}
	} else {
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, res.Msg, false))
	}
	return
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	if id := session.Get("userId"); id != nil {
		session.Clear()
		session.Save()
		ClearAuthRedis(id.(int))
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeSuccess, "注销成功", session.Get("userId")))
}

func GetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	if id := session.Get("userId"); id != nil {
		data := make(map[string]interface{}, 0)
		_ = json.Unmarshal([]byte(session.Get("data").(string)), &data)
		if menu, auths, err := getUserAllAuth(id.(int)); err == nil {
			data["auths"] = auths
			data["menu"] = menu
		} else {
			session.Clear()
			session.Save()
			c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "获取权限失败，请重新登陆", err.Error()))
			return
		}
		c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeSuccess, "已登陆", data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "未登陆", false))
}
