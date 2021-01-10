package helper

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

type ReturnType struct {
	Status 	int
	Msg		string
	Data 	interface{}
}

// 结构体转换为map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		mapValue := v.Field(i).Interface()
		// 递归获取数据
		if reflect.TypeOf(mapValue).Kind() == reflect.Struct {
			innerMap := Struct2Map(mapValue)
			for key, value := range innerMap {
				data[key] = value
			}
			continue
		}
		// 转换驼峰为下划线
		upperField := t.Field(i).Name
		field := ""
		index := 0
		for j := 0; j < len(upperField)-1; j++ {
			if (upperField[j] >= 'a' && upperField[j] <= 'z') &&
				(upperField[j+1] >= 'A' && upperField[j+1] <= 'Z') {
				field += upperField[index:j+1]+"_"
				index = j+1
			}
		}
		field += upperField[index:]
		data[strings.ToLower(field)] = v.Field(i).Interface()
	}
	return data
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

