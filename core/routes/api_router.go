package routes

import (
	apiController "OnlineJudge/app/api/controller"
	"OnlineJudge/app/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {

	// api
	api := router.Group("/api")
	{
		// 登陆接口
		api.POST("/login", apiController.DoLogin)
		api.POST("/register", apiController.Register)
		api.GET("/hello", apiController.Index)

		api.POST("/logout", apiController.DoLogout)

		//api.GET("/rotation")
		//api.GET("/data")
		//api.GET("/notice")
		api.GET("/outer/:contest_id", apiController.GetContestInfo)

		api.Use(middleware.ApiAuth())
		{
			api.GET("/notification", apiController.GetNotification)
			api.POST("/password", apiController.ForgetPassword)
			api.PUT("/password", apiController.UpdatePassword)
			problem := api.Group("/problems")
			// 题目获取鉴权

			// TODO: 参赛，比赛是否开始，题目是否为比赛题目且是否参赛
			api.POST("/print", middleware.ContestAuth(), apiController.PrintRequest)
			problem.Use(middleware.ProblemAuth())
			{
				problem.GET("", apiController.GetAllProblems)
				problem.GET("/contest/:contest_id", apiController.GetContestProblems)
				problem.GET("/problem/:problem_id", apiController.GetProblemByID)
			}

			api.GET("/rank/contest/:contest_id", apiController.GetUserRank)

			api.POST("/avatar", apiController.UploadAvatar)

			api.GET("/checklogin", apiController.Check)

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
}
