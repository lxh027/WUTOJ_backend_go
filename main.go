package main

import (
	"OnlineJudge/core/judger"
	"OnlineJudge/core/server"
	"github.com/gin-gonic/gin"
)
import "OnlineJudge/core/db"

var httpServer *gin.Engine

func main()  {
	defer func() {
		db.MySqlDb.Close()
		judger.CloseInstance()
	}()
	
	server.Run(httpServer)
}