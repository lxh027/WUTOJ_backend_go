package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func Test(c *gin.Context)  {

	var authModel = model.Authority{}
	var authValidate = validate.AuthorityValidate

	if res, err := authValidate.Validate(c, "find"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, err.Error(), ""))
		return
	}
	id, _:= strconv.ParseUint(c.Query("id"), 10, 64)
	res := authModel.GetAuthorityByID(id)
	if res.Status != common.CODE_SUCCESS{
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
}