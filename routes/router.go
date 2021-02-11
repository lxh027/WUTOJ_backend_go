package routes

import (
	 apiController "OnlineJudge/app/api/Controller"
	 panelController "OnlineJudge/app/panel/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine)  {

	// api
	api := router.Group("/api")
	{
		api.GET("/hello", apiController.Index)
		api.POST("/register", apiController.Register)

		api.POST("/do_login", apiController.DoLogin)
		api.POST("/do_logout", apiController.DoLogout)
	}

	panel := router.Group("/panel")
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

		roleAuth := panel.Group("roleAuth")
		{
			roleAuth.POST("/getRoleAuthsList", panelController.GetRoleAuthsList)
			roleAuth.POST("/addRoleAuths", panelController.AddRoleAuths)
			roleAuth.POST("/deleteRoleAuths", panelController.DeleteRoleAuths)
		}

		submitLog := panel.Group("submitLog")
		{
			submitLog.POST("/getUserSubmitStatus", panelController.GetAllUserSubmitStatus)
		}

		tag := panel.Group("tag")
		{
			tag.POST("/getAllTag", panelController.GetAllTag)
			tag.POST("/addTag", panelController.AddTag)
			tag.POST("/deleteTag", panelController.DeleteTag)
			tag.POST("/updateTag", panelController.UpdateTag)
			tag.POST("/findTagsByName", panelController.FindTagsByName)
			tag.POST("/getTagByID", panelController.GetTagByID)
			tag.POST("/changeTagStatus", panelController.ChangeTagStatus)
		}

		notice := panel.Group("notice")
		{
			notice.POST("/getAllNotice", panelController.GetAllNotice)
			notice.POST("/addNotice", panelController.AddNotice)
			notice.POST("/deleteNotice", panelController.DeleteNotice)
			notice.POST("/updateNotice", panelController.UpdateNotice)
			notice.POST("/getNoticeByID", panelController.GetNoticeByID)
		}

		contest := panel.Group("contest")
		{
			contest.POST("/getAllContest", panelController.GetAllContest)
			contest.POST("/addContest", panelController.AddContest)
			contest.POST("/deleteContest", panelController.DeleteContest)
			contest.POST("/updateContest", panelController.UpdateContest)
			contest.POST("/getContestByID", panelController.GetContestByID)
			contest.POST("/changeContestStatus", panelController.ChangeContestStatus)
		}
	}
	router.StaticFS("/admin/", http.Dir("./web"))
}