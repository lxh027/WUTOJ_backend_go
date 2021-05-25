package validate

import "OnlineJudge/app/helper"

var NoticeValidate helper.Validator

func init() {
	rules := map[string]string {
		"id"	: "required",
		"title"	: "required",
		"content"	: "required",
		"link"	: "required",
		"begintime": "required",
		"endtime": "required",
	}

	scenes := map[string] []string {
		"add" : {"title", "content", "begintime", "endtime"},
		"delete": {"id"},
		"findByID": {"id"},
		"update": {"id"},
	}

	NoticeValidate.Rules = rules
	NoticeValidate.Scenes = scenes
}
