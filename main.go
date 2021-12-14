package main

import (
	JudgerConsumer "OnlineJudge/app/common/CommonJudger/JudgerConsumer"
	"OnlineJudge/core/database"
	"OnlineJudge/core/judger"
	"OnlineJudge/core/server"

	"github.com/gin-gonic/gin"
)

var httpServer *gin.Engine

func main() {
	defer func() {
		database.MySqlDb.Close()
		judger.CloseInstance()
		JudgerConsumer.CloseJudgerConsumer()
	}()

	JudgerConsumer.RunJudgerConsumer()
	server.Run(httpServer)
}
