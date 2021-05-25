package validate

import "OnlineJudge/app/helper"

var ReplyValidate helper.Validator

func init() {
	rules := map[string]string{
		"id":         "required",
		"discuss_id": "required",
		"user_id":    "required",
		"content":    "required",
		"time":       "required",
	}

	scenes := map[string][]string{
		"add": {"discuss_id", "user_id", "content"},
	}

	ReplyValidate.Rules = rules
	ReplyValidate.Scenes = scenes
}
