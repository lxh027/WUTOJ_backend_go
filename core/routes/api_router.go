package routes

import (
	apiController "OnlineJudge/app/api/controller"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {

	// api
	api := router.Group("/api")
	{
		api.GET("/hello", apiController.Index)
		api.POST("/register", apiController.Register)

		api.POST("/login", apiController.DoLogin)
		api.POST("/logout", apiController.DoLogout)

		api.GET("/rotation")
		api.GET("/data")
		api.GET("/notice")

		api.GET("/notification", apiController.GetNotification)

		api.GET("/rank/contest/:contest_id", apiController.GetUserRank)

		api.POST("/avatar", apiController.UploadAvatar)

		api.POST("/password", apiController.ForgetPassword)
		api.PUT("/password", apiController.UpdatePassword)

		api.GET("/checklogin", apiController.Check)

		api.POST("/print", apiController.PrintRequest)

		users := api.Group("/users")
		{
			users.GET("/:user_id", apiController.GetUserByID)
			users.PUT("/:user_id", apiController.UpdateUserInfo)
		}

		contest := api.Group("/contests")
		{
			contest.GET("", apiController.GetAllContest)
			contest.GET("/contest/:contest_id", apiController.GetContestByID)
			contest.GET("/user", apiController.GetUserContest)
			contest.GET("/user/:contest_id", apiController.CheckContest)
			contest.POST("/user/:contest_id", apiController.JoinContest)
		}

		submit := api.Group("/submit")
		{
			submit.GET("", apiController.GetSubmitInfo)
			submit.GET("/id", apiController.GetSubmitByID)
			submit.GET("/problem", apiController.GetProblemSubmitInfo)
			submit.GET("/contest", apiController.GetUserContestSubmitInfo)
			submit.POST("", apiController.Submit)
		}

		problem := api.Group("/problems")
		{
			problem.GET("", apiController.GetAllProblems)
			problem.GET("/contest/:contest_id", apiController.GetContestProblems)
			problem.GET("/problem/:problem_id", apiController.GetProblemByID)

		}

		discuss := api.Group("/discussions")
		{
			discuss.GET("", apiController.GetAllDiscuss)
			discuss.GET("/problem/:problem_id", apiController.GetProblemDiscussion)
			discuss.GET("/contest/:contest_id", apiController.GetContestDiscussion)
			discuss.GET("/discuss", apiController.GetDiscussionByID) // changed
			discuss.GET("/user", apiController.GetUserDiscussion)
			discuss.POST("", apiController.AddDiscussion)

		}

		replies := api.Group("/replies")
		{
			replies.POST("", apiController.AddReply)
		}

	}
}
