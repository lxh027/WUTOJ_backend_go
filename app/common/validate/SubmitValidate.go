package validate

import "OnlineJudge/app/helper"

var SubmitValidate helper.Validator

func init() {
	rules := map[string]string{
		"id": 			"required",
		"user_id":     "required",
		"problem_id":  "required",
		"contest_id":  "required",
		"source_code": "required",
		"language":    "required",
		"page_number": "required",
	}

	scenes := map[string][]string{
		"add":             {"problem_id", "contest_id", "source_code", "language"},
		"get_problem_log": {"problem_id", "user_id"},
		"get_contest_log": {"contest_id", "user_id", "page_number"},
		"get_all":         {"user_id", "page_number"},
		"find": 		{"id"},
	}

	SubmitValidate.Rules = rules
	SubmitValidate.Scenes = scenes
}
