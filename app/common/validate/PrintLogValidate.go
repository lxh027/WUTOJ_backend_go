package validate

import "OnlineJudge/app/helper"

var PrintLogValidate helper.Validator

func init() {
	rules := map[string]string{
		"id":        "required",
		"submit_id": "required",
		"code":      "required",
	}

	scenes := map[string][]string{
		"add": {"code"},
	}

	PrintLogValidate.Rules = rules
	PrintLogValidate.Scenes = scenes
}
