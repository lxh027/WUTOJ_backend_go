package validate

import "OnlineJudge/app/helper"

var NotificationValidate helper.Validator

func init() {
	rules := map[string]string{
		"id":          "required",
		"title":       "required",
		"content":     "required",
		"submit_time": "required",
		"modify_time": "required",
		"end_time":    "required",
		"contest_id":  "required",
		"user_id":     "required",
		"status":      "required",
	}

	scenes := map[string][]string{
		"get": {"contest_id"},
	}

	NotificationValidate.Rules = rules
	NotificationValidate.Scenes = scenes
}
