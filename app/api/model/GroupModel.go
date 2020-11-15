package model

type Group struct {
	GroupID      int    `json:"group_id" form:"group_id"`
	GroupName    int    `json:"group_name" form:"group_name"`
	Avatar       string `json:"avatar" form:"avatar"`
	JoinCode     int    `json:"join_code" form:"join_code"`
	Desc         string `json:"desc" form:"desc"`
	GroupCreator int    `json:"group_creator" form:"group_creator"`
	Status       int    `json:"status" form:"status"`
}

func getAllGroup() {
	//config := config2.GetWutOjConfig()
	//PageLimit := config["page_limit"]

}

func getTheGroup(groupID int) {

}

func newGroup(data Group, userID uint) {

}
