package model

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"fmt"
	"time"
)

type UserSubmitLog struct {
	UserID 	int 	`json:"user_id" form:"user_id"`
	AC 		int 	`json:"ac" form:"ac"`
	WA 		int 	`json:"wa" form:"wa"`
	TLE 	int 	`json:"tle" form:"tle"`
	MLE 	int 	`json:"mle" form:"mle"`
	RE 		int 	`json:"re" form:"re"`
	CE 		int 	`json:"ce" form:"ce"`
	SE 		int 	`json:"se" form:"se"`
}

func (model *UserSubmitLog) GetAllUserSubmitStatus(offset int, limit int, nick string) helper.ReturnType {
	type userNickSubmitLog struct {
		UserSubmitLog
		RealName string	`json:"realname"`
		Nick 	string	`json:"nick"`
	}
	nick = "%"+nick+"%"
	sql := fmt.Sprintf("select count(*) from user_submit_log left join users on users.user_id = user_submit_log.user_id where users.nick like '%s'", nick)
	var count int
	db.Raw(sql).Row().Scan(&count)
	sql = fmt.Sprintf("select users.user_id, users.realname, ac, wa, tle, mle, re, se, ce, nick from user_submit_log left join users on users.user_id = user_submit_log.user_id where users.nick like '%s' limit %d offset %d", nick, limit, offset)
	var logs []userNickSubmitLog
	rows, err := db.Raw(sql).Rows()
	if rows != nil {
		defer rows.Close()
	}
	if err == nil {
		for rows.Next() {
			var userLog userNickSubmitLog
			err := rows.Scan(&userLog.UserID, &userLog.RealName, &userLog.AC, &userLog.WA, &userLog.TLE, &userLog.MLE, &userLog.CE, &userLog.RE, &userLog.SE, &userLog.Nick)
			if err != nil {
				fmt.Println(err.Error())
			}
			logs = append(logs, userLog)
		}
		return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
			Data: map[string]interface{}{
				"statuses": logs,
				"count": count,
			},
		}
	}
	return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
}

func (model *UserSubmitLog) GetUserSubmitStatusByTime(userId []int, startTime string, endTime string) helper.ReturnType {
	type UserSubmitLogByTime struct {
		AC int `json:"ac" form:"ac"`
		DateTime 	time.Time	`json:"date_time" form:"date_time"`
	}
	var logs [][]UserSubmitLogByTime
	for i := 0; i < len(userId); i++ {
		var result []UserSubmitLogByTime
		sql := fmt.Sprintf("SELECT count((case when (submit.status = 'AC') then submit.status end)) as ac, submit_time as date_time FROM `submit` WHERE user_id = '%d' and submit_time BETWEEN '%s' and '%s' group BY day(submit_time)", userId[i], startTime, endTime)
		rows, err := db.Raw(sql).Rows()
		if err == nil {
			for rows.Next() {
				var userLog UserSubmitLogByTime
				err := rows.Scan(&userLog.AC, &userLog.DateTime)
				if err != nil {
					fmt.Println(err.Error())
				}
				result = append(result, userLog)
			}
		} else {
			return helper.ReturnType{Status: common.CodeError, Msg: "查询失败", Data: err.Error()}
		}
		rows.Close()
		logs = append(logs, result)
	}

	return helper.ReturnType{Status: common.CodeSuccess, Msg: "查询成功",
		Data: logs,
	}
}