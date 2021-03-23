package validate

import "OnlineJudge/app/helper"

var RankValidate helper.Validator

func init() {
	rules := map[string]string{
		"contest_id": "required",
		"user_id":     "required",
	}

	scenes := map[string][]string{
		"getContestRank":      {"contest_id"},
	}

	RankValidate.Rules = rules
	RankValidate.Scenes = scenes
}
