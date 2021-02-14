package validate

import "OnlineJudge/app/helper"

var ProblemValidate helper.Validator

func init() {
	rules := map[string]string {
		"problem_id"	: "required",
		"title"	: "required",
		"content"	: "required",
		"link"	: "required",
		"begintime": "required",
		"endtime": "required",
	}

	scenes := map[string] []string {
		"add" : {"title"},
		"delete": {"problem_id"},
		"findByID": {"problem_id"},
		"update": {"problem_id"},
	}

	ProblemValidate.Rules = rules
	ProblemValidate.Scenes = scenes
}
