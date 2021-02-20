package model

type Notice struct {
	ID        int
	Title     string
	Content   string
	Link      string
	BeginTime int64
	EndTime   int64
}
