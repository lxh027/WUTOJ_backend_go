package main

import (
	"OnlineJudge/core/database"
	"OnlineJudge/core/judger"
	"OnlineJudge/core/nsqueue"
	"OnlineJudge/core/server"

	"github.com/gin-gonic/gin"
)

var httpServer *gin.Engine

func main() {
	defer func() {
		database.MySqlDb.Close()
		judger.CloseInstance()
	}()

	nsqueue.InitNSQ("Responce", "lc")

	server.Run(httpServer)
}
