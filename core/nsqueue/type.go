package nsqueue

type Response struct {
	ScrapyTime string         `json:"scrapy_time"`
	StartTime  string         `json:"start_time"`
	EndTime    string         `json:"end_time"`
	Data       []*CrawlerData `json:"data"`
}

type CrawlerData struct {
	UserInfo *UserInfo              `json:"user_info"`
	Data     map[string]*SolvedData `json:"data"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type SolvedData struct {
	Data  *SolvedData_Data `json:"data"`
	Error *Error           `json:"error"`
}

type SolvedData_Data struct {
	Statistics *Statistics   `json:"statistics"`
	Problems   []*SolvedInfo `json:"problems"`
}

type Statistics struct {
	SolvedNum int32 `json:"solved_num"`
	SubmitNum int32 `json:"submit_num"`
}

type SolvedInfo struct {
	ProblemTitle string `json:"problem_title"`
	StatusWord   string `json:"status_word"`
	Status       bool   `json:"status"`
	SubmitTime   string `json:"submit_time"`
	RelatedUrl   string `json:"related_url"`
}

type Error struct {
	Msg string `json:"msg"`
}

type TargetInfo struct {
	UserInfo *UserInfo    `json:"user_info"`
	Oj       []*OJAccount `json:"oj"`
}

type OJAccount struct {
	OjName string   `json:"oj_name"`
	Id     []string `json:"id"`
}

type Request struct {
	Targets []*TargetInfo `json:"targets"`
	Status  int           `json:"status"`
}

type GetBetweenRequest struct {
	Targets []*TargetInfo `json:"targets"`
	Start   string        `json:"start"`
	End     string        `json:"end"`
}
