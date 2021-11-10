package routes

import (
	"OnlineJudge/app/middleware"
	panelController "OnlineJudge/app/panel/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BackendRoutes(router *gin.Engine) {
	panel := router.Group("/panel")
	// Use MiddleWare to auth backend
	panel.Use(middleware.BackendAuth())
	{
		user := panel.Group("/user")
		{
			user.POST("/getAllUser", panelController.GetAllUser)
			user.POST("/getUserByID", panelController.GetUserByID)
			user.POST("/register", panelController.Register)
			user.POST("/login", panelController.Login)
			user.POST("/logout", panelController.Logout)
			user.POST("/getUserInfo", panelController.GetUserInfo)
			user.POST("/updateUser", panelController.UpdateUser)
			user.POST("/deleteUser", panelController.DeleteUser)
			user.POST("/setAdmin", panelController.SetUserAdmin)
		}

		role := panel.Group("/role")
		{
			role.POST("/getAllRole", panelController.GetAllRole)
			role.POST("/getRoleByID", panelController.GetRoleByID)
			role.POST("/addRole", panelController.AddRole)
			role.POST("/deleteRole", panelController.DeleteRole)
			role.POST("/updateRole", panelController.UpdateRole)
		}

		userRole := panel.Group("/userRole")
		{
			userRole.POST("/getUserRolesList", panelController.GetUserRolesList)
			userRole.POST("/addUserRoles", panelController.AddUserRoles)
			userRole.POST("/deleteUserRoles", panelController.DeleteUserRoles)
		}

		auth := panel.Group("/auth")
		{
			auth.POST("/getAllAuth", panelController.GetAllAuth)
			auth.POST("/getParentAuth", panelController.GetParentAuth)
			auth.POST("/addAuth", panelController.AddAuth)
			auth.POST("/deleteAuth", panelController.DeleteAuth)
			auth.POST("/getAuthByID", panelController.GetAuthByID)
			auth.POST("/updateAuth", panelController.UpdateAuth)
		}

		roleAuth := panel.Group("/roleAuth")
		{
			roleAuth.POST("/getRoleAuthsList", panelController.GetRoleAuthsList)
			roleAuth.POST("/addRoleAuths", panelController.AddRoleAuths)
			roleAuth.POST("/deleteRoleAuths", panelController.DeleteRoleAuths)
		}

		submitLog := panel.Group("/submitLog")
		{
			submitLog.POST("/getUserSubmitStatus", panelController.GetAllUserSubmitStatus)
			submitLog.POST("/getUserSubmitStatusByTime", panelController.GetUserSubmitStatusByTime)
		}

		tag := panel.Group("/tag")
		{
			tag.POST("/getAllTag", panelController.GetAllTag)
			tag.POST("/addTag", panelController.AddTag)
			tag.POST("/deleteTag", panelController.DeleteTag)
			tag.POST("/updateTag", panelController.UpdateTag)
			tag.POST("/findTagsByName", panelController.FindTagsByName)
			tag.POST("/getTagByID", panelController.GetTagByID)
			tag.POST("/changeTagStatus", panelController.ChangeTagStatus)
		}

		notice := panel.Group("/notice")
		{
			notice.POST("/getAllNotice", panelController.GetAllNotice)
			notice.POST("/addNotice", panelController.AddNotice)
			notice.POST("/deleteNotice", panelController.DeleteNotice)
			notice.POST("/updateNotice", panelController.UpdateNotice)
			notice.POST("/getNoticeByID", panelController.GetNoticeByID)
		}

		contest := panel.Group("/contest")
		{
			contest.POST("/getAllContest", panelController.GetAllContest)
			contest.POST("/addContest", panelController.AddContest)
			contest.POST("/deleteContest", panelController.DeleteContest)
			contest.POST("/updateContest", panelController.UpdateContest)
			contest.POST("/getContestByID", panelController.GetContestByID)
			contest.POST("/changeContestStatus", panelController.ChangeContestStatus)
			contest.POST("/flushRank", panelController.ClearContestRedis)
			contest.POST("/openOuterBoard", panelController.SetOuterBoard)
			contestUser := contest.Group("/contestUser")
			{
				contestUser.POST("/getAllContestUsers", panelController.GetAllContestUsers)
				contestUser.POST("/addContestUsers", panelController.AddContestUsers)
			}
			notification := contest.Group("/notification")
			{
				notification.POST("/getAllNotification", panelController.GetAllNotification)
				notification.POST("/addNotification", panelController.AddNotification)
				notification.POST("/deleteNotification", panelController.DeleteNotification)
				notification.POST("/updateNotification", panelController.UpdateNotification)
				notification.POST("/getNotificationByID", panelController.GetNotificationByID)
				notification.POST("/changeNotificationStatus", panelController.ChangeNotificationStatus)
			}
		}

		submit := panel.Group("/submit")
		{
			submit.POST("/getAllSubmit", panelController.GetAllSubmit)
			submit.POST("/getSubmitByID", panelController.GetSubmitByID)
			submit.POST("/rejudgeGroupSubmits", panelController.RejudgeGroupSubmits)
			submit.POST("/rejudgeSubmitByID", panelController.RejudgeSubmitByID)
		}

		balloon := panel.Group("/balloon")
		{
			balloon.POST("/getContestBalloon", panelController.GetContestBalloon)
			balloon.POST("/sendBalloon", panelController.SentBalloon)
		}

		printRequest := panel.Group("/print")
		{
			printRequest.POST("/getAllPrintRequest", panelController.GetAllPrintRequest)
			printRequest.POST("/handlePrintRequest", panelController.PrintRequest)
		}

		problem := panel.Group("/problem")
		{
			problem.POST("/getAllProblem", panelController.GetAllProblem)
			problem.POST("/addProblem", panelController.AddProblem)
			problem.POST("/deleteProblem", panelController.DeleteProblem)
			problem.POST("/updateProblem", panelController.UpdateProblem)
			problem.POST("/getProblemByID", panelController.GetProblemByID)
			problem.POST("/changeProblemStatus", panelController.ChangeProblemStatus)
			problem.POST("/changeProblemPublic", panelController.ChangeProblemPublic)
			problem.POST("/addSample", panelController.AddSample)
			problem.POST("/deleteSample", panelController.DeleteSample)
			problem.POST("/updateSample", panelController.UpdateSample)
			problem.POST("/findSamplesByProblemID", panelController.GetSamplesByProblemID)
			problem.POST("/uploadData", panelController.UploadData)
			problem.POST("/updateJudgeInfo", panelController.SetProblemTimeAndSpace)
			problem.POST("/uploadXML", panelController.UploadXML)
			problem.POST("/uploadImg", panelController.UploadImg)
		}
	}
	router.StaticFS("/admin/", http.Dir("./web/admin"))
	router.StaticFS("/board/", http.Dir("./web/board"))
	router.StaticFS("/uploads/", http.Dir("./web/uploads"))
}
