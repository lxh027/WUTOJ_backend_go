package main

import (
	"OnlineJudge/server"
	"github.com/gin-gonic/gin"
)
import "OnlineJudge/db_server"

var httpServer *gin.Engine

func main()  {
	defer db_server.MySqlDb.Close()

	server.Run(httpServer)

}