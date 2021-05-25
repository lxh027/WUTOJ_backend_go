package validate

import "OnlineJudge/app/helper"

var ContestValidate helper.Validator

func init() {
	rules := map[string]string{
		"contest_id":   "required",
		"contest_name": "required",
		"problems":     "required",
		"prize":        "required",
		"begin_time":   "required",
		"end_time":     "required",
		"colors":       "required",
		"rule":         "required",
		"status":       "required",
	}

	scenes := map[string][]string{
		"add":         {"contest_name", "problems", "prize", "colors", "begin_time", "end_time"},
		"delete":      {"contest_id"},
		"findByID":    {"contest_id"},
		"update":      {"contest_id"},
		"getProblems": {"contest_id"},
		"join":        {"contest_id"},
	}

	ContestValidate.Rules = rules
	ContestValidate.Scenes = scenes
}
