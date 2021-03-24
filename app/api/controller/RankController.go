package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/db_server"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserRank(c *gin.Context)  {
	rankValidate := validate.RankValidate

	rankJson := struct {
		ContestID 	int 	`uri:"contest_id"`
	}{}

	if err := c.ShouldBindUri(&rankJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	rankMap := helper.Struct2Map(rankJson)
	if res, err := rankValidate.ValidateMap(rankMap, "getContestRank"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	contestModel := model.Contest{}
	res := contestModel.GetContestById(strconv.FormatInt(int64(rankJson.ContestID), 10))
	if res.Status != common.CodeSuccess {
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}


	// try get from redis
	if rank, err := db_server.ZGetAllFromRedis("contest_rank"+strconv.Itoa(rankJson.ContestID)); err == nil {
		userIDRank, _ := redis.Strings(rank, err)
		var rankBoard []interface{}
		for _, userID := range userIDRank {
			// get each user_id
			var item map[string]interface{}
			itemStr, err := redis.String(db_server.GetFromRedis("contest_rank"+strconv.Itoa(rankJson.ContestID)+"user_id"+userID))
			if err != nil {
				c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
				return
			}
			_ = json.Unmarshal([]byte(itemStr), &item)
			rankBoard = append(rankBoard, item)
		}
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "获取排行榜成功", rankBoard))
		return
	}
	// get from DB
	getRankFromDB(c, rankJson.ContestID)
}

// TODO get rank from DB
func getRankFromDB(c *gin.Context, contestID int) {



}

