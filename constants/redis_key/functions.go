package redis_key

import "strconv"

func ContestRankUser(contestID int, userID string) string {
	return "contest_rank"+strconv.Itoa(contestID)+
		"user_id"+userID
}

func ContestRank(contestID int) string {
	return "contest_rank"+strconv.Itoa(contestID)
}

func Balloon(contestID int) string {
	return "balloon"+strconv.Itoa(int(contestID))
}

func UserNotification(nick string, contestID int) string {
	return "User:" + nick + ":Notification" + strconv.Itoa(contestID)
}

func VerifyCode(email string) string {
	return "VerifyCode" + email
}

func UserLastSubmit(userID int) string {
	return "user_last_submit" + strconv.Itoa(userID)
}

func LastPrintRequest(userID int) string {
	return "user_last_print_request" + strconv.Itoa(userID)
}

func AuthInfo(userID int) string {
	return strconv.Itoa(userID) + "auth_info"
}