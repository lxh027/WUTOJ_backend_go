package model

import (
	"OnlineJudge/app/helper"
	_ "OnlineJudge/config"
	"OnlineJudge/constants"
	_ "debug/elf"
	"time"

	_ "github.com/go-playground/locales/mgh"
)

type Contest struct {
	ContestID   int       `json:"contest_id,omitempty" form:"contest_id" uri:"contest_id"`
	ContestName string    `json:"contest_name,omitempty" form:"contest_name"`
	BeginTime   time.Time `json:"begin_time,omitempty" form:"begin_time"`
	EndTime     time.Time `json:"end_time,omitempty" form:"end_time"`
	Frozen      float64   `json:"frozen,omitempty" form:"frozen"`
	Problems    string    `json:"problems,omitempty" form:"problems"`
	Prize       string    `json:"prize,omitempty" form:"prize"`
	Colors      string    `json:"colors,omitempty" form:"colors"`
	Rule        int       `json:"rule,omitempty" form:"rule"`
	Status      int       `json:"status,omitempty" form:"status"`
}

func (Contest) TableName() string {
	return "contest"
}

func (model *Contest) GetContest(param string) helper.ReturnType {
	contest := Contest{}
	err := db.
		Select([]string{"contest_id", "contest_name", "begin_time", "end_time", "problems", "status"}).
		Where("contest_id = ?", param).
		Find(&contest).
		Error

	if err == nil {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: contest}
	} else {
		if err = db.Where("contest_name = ?", param).Find(&contest).Error; err == nil {
			return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查询成功", Data: contest}
		}
	}

	return helper.ReturnType{Status: constants.CodeError, Msg: "查询失败", Data: ""}
}

// MARK: 此接口没用
func (model *Contest) GetContestByName(contestName string) helper.ReturnType {
	contest := Contest{}
	if contestName != "" {
		err := db.Where("contest_name = ?", contestName).First(&contest).Error
		if err != nil {
			return helper.ReturnType{Status: constants.CodeError, Msg: "未找到数据", Data: err.Error()}
		} else {
			return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查找成功", Data: contest}
		}
	}
	return helper.ReturnType{Status: constants.CodeError, Msg: "未找到数据", Data: ""}
}

func (model *Contest) GetContestById(contestID int) helper.ReturnType {
	contest := Contest{}
	err := db.
		Select([]string{"contest_id", "contest_name", "begin_time", "end_time", "problems", "status", "frozen"}).
		Where("contest_id = ?", contestID).
		First(&contest).
		Error
	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "未找到数据", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查找成功", Data: contest}

}

func (model *Contest) GetAllContest(Offset int, Limit int) helper.ReturnType {
	var contests []Contest
	//
	//database.Model(&Contest{}).Find(&contests).Error
	//
	err := db.
		Select([]string{"contest_id", "contest_name", "begin_time", "end_time", "problems", "status"}).
		Offset(Offset).
		Limit(Limit).
		Find(&contests).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查找失败，数据库错误", Data: ""}
	}
	if contests != nil {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查找成功", Data: contests}
	}
	return helper.ReturnType{Status: constants.CodeError, Msg: "当前无比赛", Data: ""}

}

// MARK: 此接口没用
func (model *Contest) GetAllContestWithoutLimit() helper.ReturnType {
	var contests []Contest
	err := db.Model(&Contest{}).Find(&contests).Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "查找失败，数据库错误", Data: ""}
	}
	if contests != nil {
		return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查找成功", Data: contests}
	}
	return helper.ReturnType{Status: constants.CodeError, Msg: "当前无比赛", Data: ""}

}

func (model *Contest) GetContestStatus(ContestID int) helper.ReturnType {
	var contest = Contest{}
	err := db.
		Select([]string{"contest_id", "status"}).
		Where("contest_id = ?", ContestID).
		First(&contest).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取比赛状态失败", Data: ""}
	}

	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取比赛状态成功", Data: contest.Status}
}

func (model *Contest) GetContestProblems(ContestID int) helper.ReturnType {
	var contest = Contest{}
	err := db.
		Select([]string{"contest_id", "problems"}).
		Where("contest_id = ?", ContestID).
		First(&contest).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "获取题目失败，数据库错误", Data: ""}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "获取题目成功", Data: contest.Problems}
}

func (model *Contest) GetContestsByProblemID(problemID int, fields []string) helper.ReturnType {
	var contests []Contest

	err := db.
		Joins("JOIN contest_problem ON contest.contest_id = contest_problem.contest_id").
		Where("contest_problem.problem_id = ?", problemID).
		Select(fields).
		Find(&contests).
		Error

	if err != nil {
		return helper.ReturnType{Status: constants.CodeError, Msg: "sql query error", Data: err.Error()}
	}
	return helper.ReturnType{Status: constants.CodeSuccess, Msg: "success", Data: contests}
}

/*func (model *Contest) GetContestByProblemId(problemId int)  helper.ReturnType {
	contestsJson := model.GetAllContestWithoutLimit()
	if contestsJson.Status == constants.CodeSuccess {
		contests := contestsJson.Data.([]Contest)
		for _, contest := range contests {
			var problems []int
			err := json.Unmarshal([]byte(contest.Problems), &problems);
			if err != nil {
				// unmarshal failed
				log.Print("ProblemController: unmarshal contest problems failed, err: ", err);
				continue
			}
			for _, problem := range problems {
				if problem == problemId {
					log.Print("problem %d in contest %d", problemId, contest.ContestID)
					return helper.ReturnType{Status: constants.CodeSuccess, Msg: "查找Contest成功", Data: contest}
				}
			}
		}
	}

	return helper.ReturnType{Status: constants.CodeError, Msg: "问题不属于任何contest", Data: nil}
}*/
