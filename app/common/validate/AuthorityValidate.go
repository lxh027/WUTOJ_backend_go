package validate

import (
	"OnlineJudge/app/helper"
)

var AuthorityValidate helper.Validator

func init()  {
	rules := map[string]string{
		"id": "required",
	}

	scenes := map[string][]string {
		"find": {"id"},
	}
	AuthorityValidate.Rules = rules
	AuthorityValidate.Scenes = scenes
}


