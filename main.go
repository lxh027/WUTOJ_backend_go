package main

import (
	"OnlineJudge/judger"
	"OnlineJudge/server"
	"github.com/gin-gonic/gin"
)
import "OnlineJudge/db_server"

var httpServer *gin.Engine

func main()  {
	defer func() {
		db_server.MySqlDb.Close()
		judger.CloseInstance()
	}()
	
	server.Run(httpServer)

}