package server

import (
	"OnlineJudge/config"
	"OnlineJudge/routes"
	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine)  {
	serverConfig := config.GetServerConfig()

	// 运行模式
	gin.SetMode(serverConfig["mode"])
	httpServer = gin.Default()

	// 注册路由
	routes.Routes(httpServer)

	serverError := httpServer.Run(serverConfig["host"]+":"+serverConfig["port"])

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}
}