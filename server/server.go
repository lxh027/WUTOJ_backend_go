package server

import (
	"OnlineJudge/config"
	"OnlineJudge/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine)  {
	serverConfig  := config.GetServerConfig()
	sessionConfig := config.GetSessionConfig()

	// 运行模式
	gin.SetMode(serverConfig["mode"])
	httpServer = gin.Default()

	// 创建session存储引擎
	sessionStore := cookie.NewStore([]byte(sessionConfig["key"]))
	//设置session中间件
	httpServer.Use(sessions.Sessions(sessionConfig["name"], sessionStore))

	// 注册路由
	routes.Routes(httpServer)

	serverError := httpServer.Run(serverConfig["host"]+":"+serverConfig["port"])

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}
}