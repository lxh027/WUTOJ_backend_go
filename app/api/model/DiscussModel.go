package model

import "time"

type Discuss struct {
	ID        int
	ContestID int
	ProblemID int
	UserID    int
	Title     string
	Content   string
	Time      time.Time
	Status    int
}
