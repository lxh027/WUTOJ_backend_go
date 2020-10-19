package common

import "github.com/gin-gonic/gin"

type ReturnType struct {
	Status 	int
	Msg		string
	Data 	interface{}
}

// 模块内统一返回格式
func ReturnRes(status int, msg string, data interface{}) ReturnType {
	returnType := ReturnType{status, msg, data}
	return returnType
}

func ApiReturn(status int, msg string, data interface{}) gin.H  {
	return gin.H{
		"status"	: status,
		"msg"		: msg,
		"data"		: data,
	}
}

