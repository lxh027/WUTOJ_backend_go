package model

type Sample struct {
	SampleId  int    `json:"sample_id" form:"sample_id"`
	ProblemId int    `json:"problem_id" form:"problem_id"`
	Input     string `json:"input" form:"input"`
	Output    string `json:"output" form:"output"`
}
