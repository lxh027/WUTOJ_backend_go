package main

import (
	"OnlineJudge/core/judger"
	"OnlineJudge/core/server"
	"github.com/gin-gonic/gin"
)
import "OnlineJudge/core/database"

var httpServer *gin.Engine

func main()  {
	defer func() {
		database.MySqlDb.Close()
		judger.CloseInstance()
	}()
	
	server.Run(httpServer)
}