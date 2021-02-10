package validate

import "OnlineJudge/app/helper"

var TagValidate helper.Validator

func init() {
	rules := map[string]string {
		"id"	: "required",
		"name"	: "required",
		"description"	: "required",
		"color"	: "required",
		"status": "required",
	}

	scenes := map[string] []string {
		"add" : {"name", "color"},
		"delete": {"id"},
		"findByName": {"name"},
		"findByID": {"id"},
		"update": {"id"},
		"changeStatus": {"id", "status"},
	}

	TagValidate.Rules = rules
	TagValidate.Scenes = scenes
}
