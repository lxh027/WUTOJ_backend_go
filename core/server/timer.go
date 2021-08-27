package server

import (
	"OnlineJudge/core/database"
)

func addTimer() {
	//start := time.Now()

	db := database.MySqlDb
	update_user_sql := `INSERT IGNORE INTO user_submit_log(user_id,ac,wa,tle,mle,re,se,ce)
		select t1.user_id AS user_id, ifnull(t2.ac,0) AS ac, 
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
			on((t1.user_id = t2.user_id)))
            ON DUPLICATE KEY UPDATE user_submit_log.ac = values(ac), user_submit_log.wa = values(wa), user_submit_log.tle = values(tle), user_submit_log.mle = values(mle), user_submit_log.se = values(se), user_submit_log.ce = values(ce),user_submit_log.re = values(re)`
	db.Exec(update_user_sql)
	update_problem_sql := `INSERT IGNORE INTO problem_submit_log(problem_id,ac,wa,tle,mle,re,se,ce)
		select submit.problem_id AS problem_id,
			count((case when (submit.status = 'AC') then submit.status end)) AS ac,
			count((case when (submit.status = 'WA') then submit.status end)) AS wa,
			count((case when (submit.status = 'TLE') then submit.status end)) AS tle,
			count((case when (submit.status = 'MLE') then submit.status end)) AS mle,
			count((case when (submit.status = 'RE') then submit.status end)) AS re,
			count((case when (submit.status = 'SE') then submit.status end)) AS se,
			count((case when (submit.status = 'CE') then submit.status end)) AS ce 
		from submit group by submit.problem_id
            ON DUPLICATE KEY UPDATE problem_submit_log.ac = values(ac), problem_submit_log.wa = values(wa), problem_submit_log.tle = values(tle), problem_submit_log.mle = values(mle), problem_submit_log.se = values(se), problem_submit_log.ce = values(ce),problem_submit_log.re = values(re)`
	db.Exec(update_problem_sql)
	/*cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)*/
}


