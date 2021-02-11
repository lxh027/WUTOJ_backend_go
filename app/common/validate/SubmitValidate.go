package validate

import "OnlineJudge/app/helper"

var SubmitValidate helper.Validator

func init() {
	rules := map[string]string {
		"id"	: "required",
	}

	scenes := map[string] []string {
		"rejudge" : {"id"},
		"find"	: {"id"},
	}

	SubmitValidate.Rules = rules
	SubmitValidate.Scenes = scenes
}
