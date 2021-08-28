package server

import (
	"OnlineJudge/config"
	"OnlineJudge/core/judger"
	"OnlineJudge/core/routes"
	"io"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Run(httpServer *gin.Engine) {
	serverConfig := config.GetServerConfig()
	sessionConfig := config.GetSessionConfig()
	judgeConfig := config.GetJudgeConfig()
	// 运行模式
	gin.SetMode(serverConfig["mode"].(string))
	httpServer = gin.Default()

	// 创建session存储引擎
	sessionStore := cookie.NewStore([]byte(sessionConfig["key"].(string)))
	sessionStore.Options(sessions.Options{
		MaxAge: sessionConfig["age"].(int),
		Path:   sessionConfig["path"].(string),
	})
	//设置session中间件
	httpServer.Use(sessions.Sessions(sessionConfig["name"].(string), sessionStore))

	gin.DisableConsoleColor()
	// 生成日志
	logFile, _ := os.Create(config.GetLogPath())
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout, os.Stdin, os.Stderr)
	// 设置日志格式
	httpServer.Use(gin.LoggerWithFormatter(config.GetLogFormat))
	httpServer.Use(gin.Recovery())

	// 初始化judge
	instance := judger.InitInstance()

	instance.SetOpt(judger.OPT_SETENV, judgeConfig["env"].(string))
	instance.SetOpt(judger.OPT_SETADDR, judgeConfig["address"].(string))
	instance.SetOpt(judger.OPT_BASEDIRECTORY, judgeConfig["base_dir"].(string))
	instance.SetOpt(judger.OPT_SETTEMPDIRECTORY, judgeConfig["tmp_dir"].(string))

	go func() {
		for {
			addTimer()
			time.Sleep(10 * time.Minute)
		}
	}()

	// 注册路由
	routes.ApiRoutes(httpServer)
	routes.BackendRoutes(httpServer)

	serverError := httpServer.Run(serverConfig["host"].(string) + ":" + serverConfig["port"].(string))

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}
}
