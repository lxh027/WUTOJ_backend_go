package validate

import "OnlineJudge/app/helper"

var ContestCompetitorValidate helper.Validator

func init() {
	rules := map[string]string{
		"contest_id": "required",
		"rid":        "required",
		"rids":       "required",
	}

	scenes := map[string][]string{
		"add":         {"contest_id", "rid"},
		"addGroup":    {"contest_id", "rids"},
		"deleteGroup": {"contest_id", "rids"},
		"delete":      {"contest_id", "rid"},
		"getUserRole": {"contest_id"},
	}

	ContestCompetitorValidate.Rules = rules
	ContestCompetitorValidate.Scenes = scenes
}
