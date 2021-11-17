package controller

import (
	"OnlineJudge/app/common/validate"
	"OnlineJudge/app/helper"
	"OnlineJudge/app/panel/model"
	"OnlineJudge/config"
	"OnlineJudge/constants"
	"encoding/xml"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func GetAllProblem(c *gin.Context) {
	problemModel := model.Problem{}

	problemJson := struct {
		Offset int `json:"offset" form:"offset"`
		Limit  int `json:"limit" form:"limit"`
		Where  struct {
			Title string `json:"title" form:"title"`
		}
	}{}

	if c.ShouldBind(&problemJson) == nil {
		problemJson.Offset = (problemJson.Offset - 1) * problemJson.Limit
		res := problemModel.GetAllProblem(problemJson.Offset, problemJson.Limit, problemJson.Where.Title)
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", false))
	return
}

func GetProblemByID(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem

	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "findByID"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.FindProblemByID(problemJson.ProblemID)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func AddProblem(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}
	problemJson.Path = config.GetJudgeConfig()["base_dir"].(string) + "/tmp/0"

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "add"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.AddProblem(problemJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func DeleteProblem(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "delete"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	res := problemModel.DeleteProblem(problemJson.ProblemID)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
		return
	}
	if err := deleteProblemData(problemJson.ProblemID) ; err != nil{
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UpdateProblem(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.UpdateProblem(problemJson.ProblemID, problemJson)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeProblemStatus(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.ChangeProblemStatus(problemJson.ProblemID, problemJson.Status)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func ChangeProblemPublic(c *gin.Context) {
	problemValidate := validate.ProblemValidate
	problemModel := model.Problem{}

	var problemJson model.Problem
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	problemMap := helper.Struct2Map(problemJson)
	if res, err := problemValidate.ValidateMap(problemMap, "update"); !res {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
		return
	}

	res := problemModel.ChangeProblemPublicStatus(problemJson.ProblemID, problemJson.Public)
	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

type tomlData struct {
	ProblemID int     `json:"problem_id" form:"problem_id"`
	Time      float64 `json:"time" form:"time"`
	Memory    float64 `json:"memory" form:"memory"`
	Spj       bool    `json:"spj" form:"spj"`
	Language  string  `json:"language" form:"language"`
	Code      string  `json:"code" form:"code"`
}

func SetProblemTimeAndSpace(c *gin.Context) {
	problemModel := model.Problem{}

	problemJson := struct {
		ProblemID int     `json:"problem_id" form:"problem_id"`
		Time      float64 `json:"time" form:"time"`
		Memory    float64 `json:"memory" form:"memory"`
		Spj       bool    `json:"spj" form:"spj"`
		Language  string  `json:"language" form:"language"`
		Code      string  `json:"code" form:"code"`
	}{}
	if err := c.ShouldBind(&problemJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	updateToml(c, problemJson, false)
	var problemType string
	if problemJson.Spj {
		problemType = "Special Judge"
	} else {
		problemType = "Normal"
	}
	problemData := model.Problem{
		ProblemID: problemJson.ProblemID,
		Type: problemType,
		Time: float32(problemJson.Time),
		Memory: int(problemJson.Memory),
	}

	res := problemModel.UpdateProblem(problemJson.ProblemID, problemData)

	c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
	return
}

func UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": constants.CodeError,
			"msg": "upload img error",
			"url":"",
		})
		return
	}

	FileNameMd5 := helper.GetMd5(file.Filename)

	dst := "/uploads/image/" + FileNameMd5 + helper.RandString(11) + path.Ext(file.Filename)

	if err := c.SaveUploadedFile(file, "web"+dst); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": constants.CodeError,
			"msg": "upload img error "+err.Error(),
			"url":"",
		})
		return
	}
	serverConfig := config.GetServerConfig()
	finalUrl := fmt.Sprintf("http://%s:%s%s", serverConfig["domain"], serverConfig["port"], dst)
	c.JSON(http.StatusOK, gin.H{
		"code": constants.CodeSuccess,
		"msg": "upload img success",
		"url": finalUrl,
	})
	return
}


func UploadData(c *gin.Context) {
	//var problemModel model.Problem

	form, _ := c.MultipartForm()
	files := form.File["file[]"]
	problemDataJson := tomlData{}
	// 绑定数据
	if err := c.ShouldBind(&problemDataJson); err != nil {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "绑定数据模型失败", err.Error()))
		return
	}

	updateToml(c, problemDataJson, true)

	judgeConfig := config.GetJudgeConfig()

	dataPath := judgeConfig["base_dir"].(string) + "/" + judgeConfig["env"].(string) + "/" + strconv.Itoa(problemDataJson.ProblemID) + "/problem"

	var problemType string
	if problemDataJson.Spj {
		problemType = "Special Judge"
	} else {
		problemType = "Normal"
	}

	problemData := model.Problem{
		ProblemID: problemDataJson.ProblemID,
		Type: problemType,
		Time: float32(problemDataJson.Time),
		Memory: int(problemDataJson.Memory),
	}

	problemModel := model.Problem{}

	res := problemModel.UpdateProblem(problemDataJson.ProblemID, problemData)
	if res.Status != constants.CodeSuccess {
		c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data) )
	}
	// 保存数据文件
	filePairs := map[string]map[string]*multipart.FileHeader{}
	for _, file := range files {
		filename := file.Filename
		nameSlice := strings.Split(filename, ".")
		if len(filePairs[nameSlice[0]]) == 0 {
			filePairs[nameSlice[0]] = make(map[string]*multipart.FileHeader)
		}
		filePairs[nameSlice[0]][nameSlice[1]] = file
	}

	index := 0
	for _, filePair := range filePairs {
		dataPairPath := dataPath + "/" + strconv.Itoa(index)
		index++
		if err := os.MkdirAll(dataPairPath, 755); err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "创建路径失败", err.Error()))
			return
		}
		/*inputFile, err1 := os.Create(dataPath+"/input")
		outputFile, err2 := os.Create(dataPath+"/answer")
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(common.CodeError, "打开输入/输出文件失败", 1))
		}*/
		fmt.Println(filePair["in"].Filename, filePair["out"].Filename)
		err1 := c.SaveUploadedFile(filePair["in"], dataPairPath+"/input")
		err2 := c.SaveUploadedFile(filePair["out"], dataPairPath+"/answer")
		if err1 != nil || err2 != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "保存输入/输出文件失败", 1))
			return
		}
	}
	//problemModel.SaveProblemPath(problemDataJson.ProblemID, dataPath)
	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "上传成功", "OK"))
}


func updateToml(c *gin.Context, problemDataJson tomlData, isRebuild bool) {
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
	secs, nanos, memory := int(secsFloat), int(nanosFloat*100000000), int(problemDataJson.Memory*1024*1024)
	// 数据路径 path = {base_dir}/{id}{env}/problem
	dataPath := judgeConfig["base_dir"].(string) + "/" + judgeConfig["env"].(string) + "/" + strconv.Itoa(problemDataJson.ProblemID) + "/problem"
	//fmt.Println(dataPath)
	// 删除原目录
	if isRebuild {
		_ = os.RemoveAll(dataPath)
		if err := os.MkdirAll(dataPath, 755); err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "创建路径失败", err.Error()))
			return
		}
	}
	timeType := map[string]int{"secs": secs, "nanos": nanos}

	tomlPath := dataPath + "/config.toml"
	// 创建config.toml
	if tomlFile, err := os.Create(tomlPath); err == nil {
		defer tomlFile.Close()
		tomlEncode := toml.NewEncoder(tomlFile)
		tomlMap := map[string]interface{}{
			"problem_type": problemType,
			"limit": map[string]interface{}{
				"memory": memory, "real_time": timeType, "cpu_time": timeType,
			},
		}
		if problemDataJson.Spj {
			// SPJ
			// spj文件名
			spjFileName := "spj." + problemDataJson.Language[:strings.Index(problemDataJson.Language, ".")]
			tomlMap["spj"] = map[string]interface{}{
				"source": spjFileName, "language": problemDataJson.Language,
			}
		}
		err = tomlEncode.Encode(tomlMap)
		if err != nil {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "保存配置文件失败", err.Error()))
			return
		}
	} else {
		c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "打开配置文件失败", err.Error()))
		return
	}

	// 创建spj文件
	if problemDataJson.Spj {
		// spj文件名
		spjFileName := "spj." + problemDataJson.Language[:strings.Index(problemDataJson.Language, ".")]
		spjPath := dataPath + "/extern_program"
		if isRebuild {
			if err := os.MkdirAll(spjPath, 755); err != nil {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "创建路径失败", err.Error()))
				return
			}
		}
		if tomlFile, err := os.Create(spjPath + "/config.toml"); err == nil {
			defer tomlFile.Close()
			tomlEncode := toml.NewEncoder(tomlFile)
			spjToml := map[string]interface{}{
				"source":   spjFileName,
				"language": problemDataJson.Language,
				"timeout":  timeType,
			}
			if err = tomlEncode.Encode(spjToml); err != nil {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "保存SPJ配置文件失败", err.Error()))
				return
			}
		} else {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "打开SPJ配置文件失败", err.Error()))
			return
		}
		if spjFile, err := os.Create(spjPath + "/" + spjFileName); err == nil {
			defer spjFile.Close()
			if _, err = spjFile.WriteString(problemDataJson.Code); err != nil {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "写入spj失败", err.Error()))
				return
			}
		} else {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "打开SPJ配置文件失败", err.Error()))
			return
		}
	}
}

func UploadXML(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file[]"]
	problemValidate := validate.ProblemValidate
	var problemJson model.Problem

	for _,fileheader := range files{
		f,err := fileheader.Open()
		if err != nil{
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "打开XML文件失败", err.Error()))
			return
		}
		problemItems, err := parseProblemXml(&f)
		if err != nil{
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "解析XML文件失败", err.Error()))
			return
		}
		for _, problemItem := range problemItems{
			problemJson.Title = problemItem.Title
		problemJson.Describe = problemItem.Description
		problemJson.InputFormat = problemItem.InputFormat
		problemJson.OutputFormat = problemItem.OutputFormat
		problemJson.Hint = problemItem.Hint
		problemJson.Public = 1
		problemJson.Source = problemItem.Source
		problemJson.Time = problemItem.TimeLimit
		problemJson.Memory = problemItem.MemoryLimit
		problemJson.Type = "Normal"
		problemJson.Status = 1
		problemJson.Path = config.GetJudgeConfig()["base_dir"].(string) + "/tmp/0"

		problemMap := helper.Struct2Map(problemJson)
		if res, err := problemValidate.ValidateMap(problemMap, "add"); !res {
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, err.Error(), 0))
			return
		}

		problemModel := model.Problem{}
		res := problemModel.AddProblem(problemJson)
		if res.Status == constants.CodeError{
			c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
			return
		}
		problemID := res.Data.(int)

		problemDataJson := tomlData{}
		problemDataJson.ProblemID = problemID
		problemDataJson.Time = float64(problemItem.TimeLimit)
		problemDataJson.Memory = float64(problemItem.MemoryLimit)
		problemDataJson.Spj = false

		updateToml(c, problemDataJson, true)

		judgeConfig := config.GetJudgeConfig()

		dataPath := judgeConfig["base_dir"].(string) + "/" + judgeConfig["env"].(string) + "/" + strconv.Itoa(problemDataJson.ProblemID) + "/problem"

		if len(problemItem.TestInput) != len(problemItem.TestOutput){
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "输入输出数量不一致", err.Error()))
			return
		}
		if len(problemItem.SampleInput) != len(problemItem.SampleOutput){
			c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "样例输入输出数量不一致", err.Error()))
			return
		}
		for index,_ := range problemItem.SampleInput{
			sampleModel := model.Sample{}
			var sampleJson model.Sample
			sampleJson.ProblemID = problemID
			sampleJson.Input = problemItem.SampleInput[index]
			sampleJson.Output = problemItem.SampleOutput[index]
			res := sampleModel.AddSample(sampleJson)
			if res.Status == constants.CodeError{
				c.JSON(http.StatusOK, helper.BackendApiReturn(res.Status, res.Msg, res.Data))
				return
			}
		}
		for index,_ := range problemItem.TestInput{
			dataPairPath := dataPath + "/" + strconv.Itoa(index)
			if err := os.MkdirAll(dataPairPath, 755); err != nil {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "创建路径失败", err.Error()))
				return
			}
			err1 := ioutil.WriteFile(dataPairPath+"/input",[]byte(problemItem.TestInput[index]),0666)
			err2 := ioutil.WriteFile(dataPairPath+"/answer",[]byte(problemItem.TestOutput[index]),0666)
			if err1 != nil || err2 != nil {
				c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeError, "保存输入/输出文件失败", 1))
				return
			}
		}
		}
		
	}

	c.JSON(http.StatusOK, helper.BackendApiReturn(constants.CodeSuccess, "解析XML成功", "OK"))
}

type ProblemItem struct {
	Title        string   `xml:"title"`
	TimeLimit    float32  `xml:"time_limit"`
	MemoryLimit  int      `xml:"memory_limit"`
	Description  string   `xml:"description"`
	InputFormat  string   `xml:"input"`
	OutputFormat string   `xml:"output"`
	SampleInput  []string `xml:"sample_input"`
	SampleOutput []string `xml:"sample_output"`
	TestInput    []string `xml:"test_input"`
	TestOutput   []string `xml:"test_output"`
	Hint         string   `xml:"hint"`
	Source       string   `xml:"source"`
}

type ProblemXml struct {
	XMLName xml.Name    `xml:"fps"`
	Item    []ProblemItem `xml:"item"`
}

func parseProblemXml(file *multipart.File) ([]ProblemItem, error) {
	v := ProblemXml{}
	data, err := ioutil.ReadAll(*file)
	if err != nil {
		return v.Item, fmt.Errorf("read file error:%v", err)
	}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		return v.Item, fmt.Errorf("unmarshal error:%v", err)
	}
	reg1 := regexp.MustCompile(`class=".*"`)
	reg2 := regexp.MustCompile(`id=".*"`)

	for index,_ := range v.Item{
	v.Item[index].Description = reg1.ReplaceAllString(v.Item[index].Description, "")
	v.Item[index].Hint = reg1.ReplaceAllString(v.Item[index].Hint, "")
	v.Item[index].InputFormat = reg1.ReplaceAllString(v.Item[index].InputFormat, "")
	v.Item[index].OutputFormat = reg1.ReplaceAllString(v.Item[index].OutputFormat, "")
	v.Item[index].Source = reg1.ReplaceAllString(v.Item[index].Source, "")

	v.Item[index].Description = reg2.ReplaceAllString(v.Item[index].Description, "")
	v.Item[index].Hint = reg2.ReplaceAllString(v.Item[index].Hint, "")
	v.Item[index].InputFormat = reg2.ReplaceAllString(v.Item[index].InputFormat, "")
	v.Item[index].OutputFormat = reg2.ReplaceAllString(v.Item[index].OutputFormat, "")
	v.Item[index].Source = reg2.ReplaceAllString(v.Item[index].Source, "")
	}


	return v.Item, nil
}

func deleteProblemData(ProblemID int) error {
	judgeConfig := config.GetJudgeConfig()
	dataPath := judgeConfig["base_dir"].(string) + "/" + judgeConfig["env"].(string) + "/" + strconv.Itoa(ProblemID) 
	if isExist(dataPath) == false {
		return fmt.Errorf("DeleteProblemData error: No data exists")
	}
	if err := os.RemoveAll(dataPath); err != nil{
		return fmt.Errorf("DeleteProblemData error:"+err.Error())
	}
	return nil
}


func isExist(path string)(bool){
    _, err := os.Stat(path)
    if err != nil{
        if os.IsExist(err){
            return true
        }
        if os.IsNotExist(err){
            return false
        }
        return false
    }
    return true
}