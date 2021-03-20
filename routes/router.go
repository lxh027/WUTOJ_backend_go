package routes

import (
	apiController "OnlineJudge/app/api/Controller"
	panelController "OnlineJudge/app/panel/Controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {

	// api
	api := router.Group("/api")
	{
		api.GET("/hello", apiController.Index)
		api.POST("/register", apiController.Register)

		api.POST("/do_login", apiController.DoLogin)
		api.POST("/do_logout", apiController.DoLogout)

		api.GET("/rotation")
		api.GET("/data")
		api.GET("/notice")

		api.GET("/rank/contest/:contest_id")

		contest := api.Group("/contests")
		{
			contest.GET("", apiController.GetAllContest)
			contest.GET("/contest/:param", apiController.SearchContest)
			contest.GET("/user/:contest_id", apiController.CheckContest)
			contest.POST("/user/:contest_id", apiController.JoinContest)
		}

		submit := api.Group("/submit")
		{
			submit.GET("", apiController.GetSubmitInfo)
			submit.POST("", apiController.Submit)
		}

		problem := api.Group("/problems")
		{
			problem.GET("", apiController.GetAllProblems)
			problem.GET("/contest/:contest_id", apiController.GetContestProblems)
			problem.GET("/{problem_id}", apiController.GetProblemByID)

		}

		discuss := api.Group("/discussions")
		{
			discuss.GET("", apiController.GetAllDiscuss)
			discuss.GET("/problem/:problem_id", apiController.GetProblemDiscussion)
			discuss.GET("/contest/:contest_id", apiController.GetContestDiscussion)
			discuss.GET("/discuss", apiController.GetDiscussionByID) // changed
			discuss.POST("", apiController.AddDiscussion)

		}

		replies := api.Group("/replies")
		{
			replies.POST("", apiController.AddReply)
		}

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
	}
	router.StaticFS("/admin/", http.Dir("./web"))
}
