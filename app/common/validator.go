package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"log"
)

type Validator struct {
	Rules map[string]string
	Scenes map[string] []string
}

func (validator *Validator)Validate(c *gin.Context, scene string) (bool, error) {
	// 判断scene是否存在
	if _, ok := validator.Scenes[scene]; !ok {
		msg := errors.New("scene is not exists")
		log.Println(msg.Error()) // all error messages
		return false, msg
	}
	httpData, err := validate.FromRequest(c.Request)
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	// 创建验证器
	v := httpData.Create()
	//  添加规则
	for _, field := range validator.Scenes[scene] {
		v.StringRule(field, validator.Rules[field])
	}

	if v.Validate() {
		return true, errors.New("")
	} else {
		log.Println(v.Errors) // all error messages
		return false, errors.New(v.Errors.One())
	}
}
