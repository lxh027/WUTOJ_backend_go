package validate

import "OnlineJudge/app/helper"

var SubmitValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id":     "required",
		"problem_id":  "required",
		"contest_id":  "required",
		"source_code": "required",
		"language":    "required",
	}

	scenes := map[string][]string{
		"add": {"problem_id", "contest_id", "source_code", "language"},
	}

	SubmitValidate.Rules = rules
	SubmitValidate.Scenes = scenes
}
