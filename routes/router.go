package routes

import (
	"OnlineJudge/app/api/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine)  {

	// api
	api := router.Group("/api")
	{
		api.GET("/", controller.Index)
		api.POST("/register", controller.Register)

		api.POST("/do_login", controller.DoLogin)
		api.POST("/do_logout", controller.DoLogout)
	}

}