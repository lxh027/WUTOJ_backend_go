package controller

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/config"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetAllProblem(c *gin.Context)  {
	if res := haveAuth(c, "getAllProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemModel := model.Problem{}

	problemJson := struct {
		Offset 	int 	`json:"offset" form:"offset"`
		Limit 	int 	`json:"limit" form:"limit"`
		Where 	struct{
			Title 	string 	`json:"title" form:"title"`
		}
	}{}

	if c.ShouldBind(&problemJson) == nil {
		problemJson.Offset = (problemJson.Offset-1)*problemJson.Limit
		res := problemModel.GetAllProblem(problemJson.Offset, problemJson.Limit, problemJson.Where.Title)
		c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", false))
	return
}

func GetProblemByID(c *gin.Context) {
	if res := haveAuth(c, "getAllProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem

	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.FindProblemByID(problemJson.ProblemID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddProblem(c *gin.Context) {
	if res := haveAuth(c, "addProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	problemJson.Path = config.GetJudgeConfig()["base_dir"].(string)+"/tmp/0"

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "add"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.AddProblem(problemJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteProblem(c *gin.Context) {
	if res := haveAuth(c, "deleteProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.DeleteProblem(problemJson.ProblemID)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateProblem(c *gin.Context) {
	if res := haveAuth(c, "updateProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.UpdateProblem(problemJson.ProblemID, problemJson)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeProblemStatus(c *gin.Context) {
	if res := haveAuth(c, "updateProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.ChangeProblemStatus(problemJson.ProblemID, problemJson.Status)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeProblemPublic(c *gin.Context) {
	if res := haveAuth(c, "updateProblem"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err:= problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.ChangeProblemPublicStatus(problemJson.ProblemID, problemJson.Public)
	c.JSON(http.StatusOK, helper.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UploadData(c *gin.Context) {
	if res := haveAuth(c, "uploadData"); res != common.Authed {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "权限不足", res))
		return
	}

	//var problemModel model.Problem

	form, _ := c.MultipartForm()
	files := form.File["file[]"]
	problemDataJson := struct {
		ProblemID 	int 	`json:"problem_id" form:"problem_id"`
		Time 		float64 `json:"time" form:"time"`
		Memory 		float64 `json:"memory" form:"memory"`
		Spj 		bool 	`json:"spj" form:"spj"`
		Language 	string 	`json:"language" form:"language"`
		Code 		string 	`json:"code" form:"code"`
	}{}
	// 绑定数据
	if err := c.ShouldBind(&problemDataJson); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	//c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "绑定数据模型成功", 1))
	judgeConfig := config.GetJudgeConfig()
	var problemType string
	if problemDataJson.Spj {
		problemType = "Special Judge"
	} else {
		problemType = "Normal"
	}
	// 解析数据
	secsFloat, nanosFloat := math.Modf(problemDataJson.Time)
	// 秒 纳秒 内粗你
	secs, nanos, memory := int(secsFloat), int(nanosFloat*100000000), int(problemDataJson.Memory * 1024 * 1024)
	// 数据路径 path = {base_dir}/{id}{env}/problem
	dataPath := judgeConfig["base_dir"].(string)+"/"+judgeConfig["env"].(string)+"/"+strconv.Itoa(problemDataJson.ProblemID)+"/problem"
	//fmt.Println(dataPath)
	// 删除原目录
	_ = os.RemoveAll(dataPath)
	if err := os.MkdirAll(dataPath, 755); err != nil {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建路径失败", err.Error()))
		return
	}
	timeType := map[string]int { "secs": secs, "nanos": nanos}

	tomlPath := dataPath+"/config.toml"
	// 创建config.toml
	if tomlFile, err := os.Create(tomlPath); err == nil {
		defer tomlFile.Close()
		tomlEncode := toml.NewEncoder(tomlFile)
		tomlMap := map[string]interface{} {
			"problem_type": problemType,
			"limit": map[string]interface{} {
				"memory": memory, "real_time": timeType, "cpu_time": timeType,
			},
		}
		if problemDataJson.Spj {
			// SPJ
			// spj文件名
			spjFileName := "spj."+problemDataJson.Language[:strings.Index(problemDataJson.Language, ".")]
			tomlMap["spj"] = map[string]interface{} {
				"source": spjFileName, "language": problemDataJson.Language,
			}
		}
		err = tomlEncode.Encode(tomlMap)
		if err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "保存配置文件失败", err.Error()))
			return
		}
	} else {
		c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "打开配置文件失败", err.Error()))
		return
	}

	// 创建spj文件
	if problemDataJson.Spj {
		// spj文件名
		spjFileName := "spj."+problemDataJson.Language[:strings.Index(problemDataJson.Language, ".")]
		spjPath := dataPath+"/extern_program"
		if err := os.MkdirAll(spjPath, 755); err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建路径失败", err.Error()))
			return
		}
		if tomlFile, err := os.Create(spjPath+"/config.toml"); err == nil {
			defer tomlFile.Close()
			tomlEncode := toml.NewEncoder(tomlFile)
			spjToml := map[string]interface{} {
				"source": spjFileName,
				"language": problemDataJson.Language,
				"timeout": timeType,
			}
			if err = tomlEncode.Encode(spjToml); err != nil {
				c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "保存SPJ配置文件失败", err.Error()))
				return
			}
		} else {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "打开SPJ配置文件失败", err.Error()))
			return
		}
		if spjFile, err := os.Create(spjPath+"/"+spjFileName); err == nil {
			defer spjFile.Close()
			if _, err = spjFile.WriteString(problemDataJson.Code); err != nil {
				c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "写入spj失败", err.Error()))
				return
			}
		} else {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "打开SPJ配置文件失败", err.Error()))
			return
		}
	}
	// 保存数据文件
	filePairs := map[string]map[string]*multipart.FileHeader{}
	for _, file := range files {
		filename := file.Filename
		nameSlice := strings.Split(filename, ".")
		if len(filePairs[nameSlice[0]]) == 0{
			filePairs[nameSlice[0]] = make(map[string]*multipart.FileHeader)
		}
		filePairs[nameSlice[0]][nameSlice[1]] = file
	}

	index := 0
	for _, filePair := range filePairs {
		dataPairPath := dataPath+"/"+strconv.Itoa(index)
		index++
		if err := os.MkdirAll(dataPairPath, 755); err != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "创建路径失败", err.Error()))
			return
		}
		/*inputFile, err1 := os.Create(dataPath+"/input")
		outputFile, err2 := os.Create(dataPath+"/answer")
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "打开输入/输出文件失败", 1))
		}*/
		fmt.Println(filePair["in"].Filename, filePair["out"].Filename)
		err1 := c.SaveUploadedFile(filePair["in"], dataPairPath+"/input")
		err2 := c.SaveUploadedFile(filePair["out"], dataPairPath+"/answer")
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "保存输入/输出文件失败", 1))
			return
		}
	}
	//problemModel.SaveProblemPath(problemDataJson.ProblemID, dataPath)
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeSuccess, "上传成功", "OK"))
}