package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/panel/model"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type menuItem struct {
	Title 	string `json:"title"`
	Icon	string `json:"icon"`
	Href	string `json:"href"`
	Target 	string `json:"target"`
	Child 	[]menuItem `json:"child"`
}


func haveAuth(c *gin.Context, authQuery string) int {
	session := sessions.Default(c)
	id := session.Get("user_id")
	if  id == nil {
		return common.UnLoggedIn
	} else if session.Get("is_admin").(int) == 0 {
		return common.UnAuthed
	}
	_, auths, err := getUserAllAuth(id.(int))
	if err != nil {
		return common.AuthError
	} else {
		for _, auth := range auths {
			if auth == authQuery {
				return common.Authed
			}
		}
		return common.UnAuthed
	}
}

func getUserAllAuth(userID int) ([]menuItem, []string, error) {
	authModel := model.Auth{}

	if res := authModel.GetUserAllAuth(userID); res.Status == common.CodeSuccess {
		auths := res.Data.([]model.Auth)
		var authsLeft []model.Auth
		var authName []string
		var menu []menuItem

		menuItemCount := 0
		type2Pos := map[int]int{}

		for _, auth := range auths {
			if auth.Type == 2 {
				authName = append(authName, auth.Title)
			} else if auth.Type == 0 {
				item := menuItem{
					Title: auth.Title,
					Target: auth.Target,
					Icon: auth.Icon,
					Href: auth.Href,
				}
				menu = append(menu, item)
				type2Pos[auth.Aid] = menuItemCount
				menuItemCount++
			} else if auth.Type == 1 {
				authsLeft = append(authsLeft, auth)
			}
		}
		if menu != nil {
			for _, auth := range authsLeft {
				item := menuItem{
					Title: auth.Title,
					Target: auth.Target,
					Icon: auth.Icon,
					Href: auth.Href,
				}
				pos := type2Pos[auth.Parent]
				menu[pos].Child = append(menu[pos].Child, item)
			}
		} else {
			menu = make([]menuItem, 0)
		}

		return menu, authName, nil
	} else {
		return nil, nil, errors.New("获取权限列表错误")
	}

}
