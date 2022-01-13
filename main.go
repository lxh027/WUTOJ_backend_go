package main

import (
	"OnlineJudge/core/database"
	"OnlineJudge/core/grpc/rpcconn"
	"OnlineJudge/core/judger"
	"OnlineJudge/core/server"

	"github.com/gin-gonic/gin"
)

var httpServer *gin.Engine

func main() {
	defer func() {
		database.MySqlDb.Close()
		judger.CloseInstance()
		rpcconn.RPCConn.Close()
	}()

	server.Run(httpServer)
}
