package validate

import "OnlineJudge/app/helper"

var DiscussValidate helper.Validator

func init() {
	rules := map[string]string{
		"discuss_id":  "required",
		"contest_id":  "required",
		"problem_id":  "required",
		"user_id":     "required",
		"title":       "required",
		"content":     "required",
		"time":        "required",
		"status":      "required",
		"page_number": "required",
	}

	scenes := map[string][]string{
		"add":             {"contest_id", "problem_id", "content", "title"},
		"findByContestID": {"contest_id", "page_number"},
		"findByProblemID": {"problem_id", "page_number"},
		"findByID":        {"discuss_id", "page_number"},
	}

	DiscussValidate.Rules = rules
	DiscussValidate.Scenes = scenes
}
