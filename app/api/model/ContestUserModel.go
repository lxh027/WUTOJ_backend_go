package model

type ContestUser struct {
	ContestID int `json:"contest_id" form:"contest_id"`
	UserID    int `json:"user_id" form:"user_id"`
	ID        int `json:"id" form:"id"`
	Status    int `json:"status" form:"status"`
}
