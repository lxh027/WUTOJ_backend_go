package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)


func Test(c *gin.Context)  {

	var authModel = model.Authority{}
	var authValidate = validate.AuthorityValidate

	if res, err := authValidate.Validate(c, "find"); !res {
		log.Println(err.Error())
		return
	}
	id, err:= strconv.ParseUint(c.Query("id"), 10, 64)
	auths, err := authModel.GetAuthorityByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ERROR",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
		"data": auths,
	})
}