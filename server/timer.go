package server

import (
	"OnlineJudge/db_server"
	"fmt"
)

type userSubmitLog struct {
	UserID 	int 	`json:"user_id"`
	AC 		int 	`json:"ac"`
	WA 		int 	`json:"wa"`
	TLE 	int 	`json:"tle"`
	MLE 	int 	`json:"mle"`
	RE 		int 	`json:"re"`
	SE 		int 	`json:"se"`
	CE 		int 	`json:"ce"`
}

type problemSubmitLog struct {
	ProblemID 	int 	`json:"problem_id"`
	AC 		int 	`json:"ac"`
	WA 		int 	`json:"wa"`
	TLE 	int 	`json:"tle"`
	MLE 	int 	`json:"mle"`
	RE 		int 	`json:"re"`
	SE 		int 	`json:"se"`
	CE 		int 	`json:"ce"`
}

func addTimer() {
	db := db_server.MySqlDb
	sql := `select t1.user_id AS user_id, ifnull(t2.ac,0) AS ac, 
		t1.wa AS wa, t1.tle AS tle, t1.mle AS mle, t1.re AS re, t1.se AS se, t1.ce AS ce
		from (((select submit.user_id AS user_id, 
				count(distinct submit.problem_id) AS ac 
					from submit where (submit.status = 'AC') 
			group by submit.user_id)) t2 
			left join ((
				select submit.user_id AS user_id,
					count((case when (submit.status = 'WA') then submit.status end)) AS wa,
					count((case when (submit.status = 'TLE') then submit.status end)) AS tle,
					count((case when (submit.status = 'MLE') then submit.status end)) AS mle,
					count((case when (submit.status = 'RE') then submit.status end)) AS re,
					count((case when (submit.status = 'SE') then submit.status end)) AS se,
					count((case when (submit.status = 'CE') then submit.status end)) AS ce 
				from submit group by submit.user_id)) t1 
			on((t1.user_id = t2.user_id)))`
	rows, err := db.Raw(sql).Rows()
	if err == nil {
		defer func() { _ = rows.Close()}()
		var userSubmitLog userSubmitLog
		addSql := "INSERT INTO user_submit_log (user_id, ac, wa, tle, mle, re, se, ce) VALUES "
		index := 0
		for rows.Next() {
			_ = rows.Scan(&userSubmitLog)
			data := fmt.Sprintf("(%d, %d, %d, %d, %d, %d, %d, %d)",
				userSubmitLog.UserID, userSubmitLog.AC, userSubmitLog.WA, userSubmitLog.TLE,
				userSubmitLog.MLE, userSubmitLog.RE, userSubmitLog.SE, userSubmitLog.CE)
			if index != 0 {
				addSql = addSql+", "
			}
			addSql = addSql+data
			index++
		}
		if index != 0 {
			addSql = addSql+" ac = VALUES(ac), wa = VALUES(wa), tle = VALUES(tle), mle = VALUES(mle), re = VALUES(re), se = VALUES(se), ce = VALUES(ce)"
			db.Exec(addSql)
		}
	}
}


