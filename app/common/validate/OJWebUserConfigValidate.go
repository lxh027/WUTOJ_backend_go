package validate

import "OnlineJudge/app/helper"

var OJWebUserConfigValidate helper.Validator

func init() {
	rules := map[string]string{
		"user_id": "required",
	}

	scenes := map[string][]string{
		"add":      {"user_id", "oj_name", "oj_user_name"},
		"findByID": {"id"},
		"delete":   {"id"},
		"update":   {"id"},
	}

	OJWebUserConfigValidate.Rules = rules
	OJWebUserConfigValidate.Scenes = scenes
}
