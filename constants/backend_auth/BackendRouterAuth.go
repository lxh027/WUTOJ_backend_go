package backend_auth

func GetBackendRouterAuth(url string) string {
	return auths[url]
}

const AuthPass = "PASS"

var auths = map[string]string{
	// user
	"/panel/user/getAllUser":  "getAllUser",
	"/panel/user/getUserByID": "getAllUser",
	"/panel/user/register":    AuthPass,
	"/panel/user/login":       AuthPass,
	"/panel/user/logout":      AuthPass,
	"/panel/user/getUserInfo": AuthPass,
	"/panel/user/updateUser":  "updateUser",
	"/panel/user/deleteUser":  "deleteUser",
	"/panel/user/setAdmin":    "roleAssign",
	// role
	"/panel/role/getAllRole":  "getAllRole",
	"/panel/role/getRoleByID": "getAllRole",
	"/panel/role/addRole":     "addRole",
	"/panel/role/deleteRole":  "deleteRole",
	"/panel/role/updateRole":  "updateRole",

	// userRole
	"/panel/userRole/getUserRolesList": "roleAssign",
	"/panel/userRole/addUserRoles":     "roleAssign",
	"/panel/userRole/deleteUserRoles":  "roleAssign",

	// auth
	"/panel/auth/getAllAuth":    "getAllAuth",
	"/panel/auth/getParentAuth": "getAllAuth",
	"/panel/auth/addAuth":       "addAuth",
	"/panel/auth/deleteAuth":    "deleteAuth",
	"/panel/auth/getAuthByID":   "getAllAuth",
	"/panel/auth/updateAuth":    "updateAuth",

	// roleAuth
	"/panel/roleAuth/getRoleAuthsList": "authAssign",
	"/panel/roleAuth/addRoleAuths":     "authAssign",
	"/panel/roleAuth/deleteRoleAuths":  "authAssign",

	// submitLog
	"/panel/submitLog/getUserSubmitStatus":       "getUserSubmit",
	"/panel/submitLog/getUserSubmitStatusByTime": "getUserSubmit",

	// tag
	"/panel/tag/getAllTag":       "getAllTag",
	"/panel/tag/addTag":          "addTag",
	"/panel/tag/deleteTag":       "deleteTag",
	"/panel/tag/updateTag":       "updateTag",
	"/panel/tag/findTagsByName":  "getAllTag",
	"/panel/tag/getTagByID":      "getAllTag",
	"/panel/tag/changeTagStatus": "updateTag",

	// notice
	"/panel/notice/getAllNotice":  "getAllNotice",
	"/panel/notice/addNotice":     "addNotice",
	"/panel/notice/deleteNotice":  "deleteNotice",
	"/panel/notice/updateNotice":  "updateNotice",
	"/panel/notice/getNoticeByID": "getAllNotice",

	// contest
	"/panel/contest/getAllContest":       "getAllContest",
	"/panel/contest/addContest":          "addContest",
	"/panel/contest/deleteContest":       "deleteContest",
	"/panel/contest/updateContest":       "updateContest",
	"/panel/contest/getContestByID":      "getAllContest",
	"/panel/contest/changeContestStatus": "updateContest",
	"/panel/contest/flushRank":           "updateContest",
	"/panel/contest/openOuterBoard":	  "updateContest",

	//contest/contestUser
	"/panel/contest/contestUser/getAllContestUsers": "getAllContest",
	"/panel/contest/contestUser/addContestUsers":    "getAllContest",

	// contest/notification
	"/panel/contest/notification/getAllNotification":       "getAllContest",
	"/panel/contest/notification/addNotification":          "updateContest",
	"/panel/contest/notification/deleteNotification":       "updateContest",
	"/panel/contest/notification/updateNotification":       "updateContest",
	"/panel/contest/notification/getNotificationByID":      "getAllContest",
	"/panel/contest/notification/changeNotificationStatus": "updateContest",

	// submit
	"/panel/submit/getAllSubmit":        "getAllSubmit",
	"/panel/submit/getSubmitByID":       "getAllSubmit",
	"/panel/submit/rejudgeGroupSubmits": "rejudge",
	"/panel/submit/rejudgeSubmitByID":   "rejudge",

	// balloon
	"/panel/balloon/getContestBalloon": "getBalloonStatus",
	"/panel/balloon/sendBalloon":       "setBalloonStatus",

	// print
	"/panel/print/getAllPrintRequest": "getPrintRequest",
	"/panel/print/handlePrintRequest": "getPrintRequest",

	// problem
	"/panel/problem/getAllProblem":          "getAllProblem",
	"/panel/problem/addProblem":             "addProblem",
	"/panel/problem/deleteProblem":          "deleteProblem",
	"/panel/problem/updateProblem":          "updateProblem",
	"/panel/problem/getProblemByID":         "getAllProblem",
	"/panel/problem/changeProblemStatus":    "updateProblem",
	"/panel/problem/changeProblemPublic":    "updateProblem",
	"/panel/problem/addSample":              "addProblem",
	"/panel/problem/deleteSample":           "deleteProblem",
	"/panel/problem/updateSample":           "updateProblem",
	"/panel/problem/findSamplesByProblemID": "getAllProblem",
	"/panel/problem/uploadData":             "uploadData",
	"/panel/problem/updateJudgeInfo":        "uploadData",
	"/panel/problem/uploadXML":              "uploadData",
	"/panel/problem/uploadImg":              "uploadData",
}
