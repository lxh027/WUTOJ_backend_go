package controller

import (
	"OnlineJudge/app/panel/model"
	"OnlineJudge/constants"
	"OnlineJudge/constants/redis_key"
	"OnlineJudge/core/database"
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type menuItem struct {
	Title  string     `json:"title"`
	Icon   string     `json:"icon"`
	Href   string     `json:"href"`
	Target string     `json:"target"`
	Child  []menuItem `json:"child"`
}

type redisData struct {
	Menu []menuItem `json:"menu"`
	Auth []string   `json:"auth"`
}

func ClearAuthRedis(userID int) {
	err := database.DeleteFromRedis(redis_key.AuthInfo(userID))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func GetUserAllAuth(userID int) ([]menuItem, []string, error) {
	if info, err := database.GetFromRedis(redis_key.AuthInfo(userID)); err == nil && info != nil {
		var authInfo redisData
		bytes, _ := redis.Bytes(info, err)
		_ = json.Unmarshal(bytes, &authInfo)
		return authInfo.Menu, authInfo.Auth, nil
	}

	authModel := model.Auth{}

	if res := authModel.GetUserAllAuth(userID); res.Status == constants.CodeSuccess {
		auths := res.Data.([]model.Auth)
		authsLeft := map[int]model.Auth{}
		var authName []string
		var menu []menuItem

		childMenu := map[int][]menuItem{}
		hasChild := map[int]int{}
		menuItemCount := 0

		for _, auth := range auths {
			if auth.Type == 2 {
				authName = append(authName, auth.Title)
			} else {
				hasChild[auth.Aid] = 0
				childMenu[auth.Aid] = make([]menuItem, 0)
				authsLeft[auth.Aid] = auth
				if auth.Type == 0 {
					menuItemCount++
				}
			}
		}

		for _, auth := range authsLeft {
			hasChild[auth.Parent]++
		}
		queue := list.New()
		for _, auth := range authsLeft {
			if hasChild[auth.Aid] == 0 {
				queue.PushBack(auth)
			}
		}

		for queue.Len() > 0 {
			auth := queue.Front().Value.(model.Auth)
			queue.Remove(queue.Front())
			item := menuItem{
				Title:  auth.Title,
				Icon:   auth.Icon,
				Href:   auth.Href,
				Target: auth.Target,
				Child:  childMenu[auth.Aid],
			}
			childMenu[auth.Parent] = append(childMenu[auth.Parent], item)
			hasChild[auth.Parent]--
			if hasChild[auth.Parent] == 0 {
				parentAuth := authsLeft[auth.Parent]
				if parentAuth.Type == 0 {
					menu = append(menu, menuItem{
						Title:  parentAuth.Title,
						Icon:   parentAuth.Icon,
						Href:   parentAuth.Href,
						Target: parentAuth.Target,
						Child:  childMenu[parentAuth.Aid],
					})
				} else {
					queue.PushBack(authsLeft[auth.Parent])
				}
			}
		}
		authInfo := redisData{menu, authName}
		dataJson, _ := json.Marshal(authInfo)
		_ = database.PutToRedis(redis_key.AuthInfo(userID), dataJson, 3600)
		return menu, authName, nil
	} else {
		return nil, nil, errors.New("获取权限列表错误")
	}

}

func getContestTime(contestID uint) (time.Time, time.Time, time.Time, error) {
	contestModel := model.Contest{}
	res := contestModel.GetContestById(contestID)
	now := time.Now()
	if res.Status != constants.CodeSuccess {
		return now, now, now, errors.New(res.Msg)
	}
	beginTime := res.Data.(model.Contest).BeginTime
	endTime := res.Data.(model.Contest).EndTime
	frozen := res.Data.(model.Contest).Frozen
	frozenTime := time.Unix(int64(float64(endTime.Unix())*frozen+float64(beginTime.Unix())), 0)

	return beginTime, frozenTime, endTime, nil
}
