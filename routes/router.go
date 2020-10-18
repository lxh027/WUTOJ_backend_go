package routes

import (
	"OnlineJudge/app/api/controller"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine)  {
	router.GET("/", controller.Index)
	router.GET("/test", controller.Test)
}