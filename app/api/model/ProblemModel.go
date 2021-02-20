package model

type Problem struct {
	ProblemID    uint    `json:"problem_id" form:"problem_id"`
	Title        string  `json:"title" form:"title" `
	Background   string  `json:"background" form:"background"`
	Describe     string  `json:"describe" form:"describe"`
	InputFormat  string  `json:"input_format" form:"input_format"`
	OutputFormat string  `json:"output_format" form:"output_format"`
	Hint         string  `json:"hint" form:"hint"`
	Public       uint    `json:"public" form:"public"`
	Source       string  `json:"source" form:"source`
	Time         float64 `json:"time" form:"time"`
	Memory       int     `json:"memory" form:"memory"`
	Type         string  `json:"type" form:"type"`
	Tag          string  `json:"tag" form:"tag"`
	Path         string  `json:"path" form:"path"`
	Status       int     `json:"status" form:"status"`
}
