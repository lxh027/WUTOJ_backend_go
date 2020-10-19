package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetLogFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func GetLogPath() string  {
	timeObj := time.Now()
	datetime := timeObj.Format("2006-01-02-15-04-05")
	return "log/OnlineJudge"+datetime+".log"
}
