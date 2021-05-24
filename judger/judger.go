package judger

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

var judgerInstance *judger

const (
	OPT_BASEDIRECTORY = iota
	OPT_SETENV
	OPT_SETADDR
	OPT_SETTEMPDIRECTORY
)

const (
	MAX_RETRY_TIMES     = 5
	MAX_RECONNECT_TIMES = 10
)

// initInstance - init singleton judger instance
func InitInstance() *judger {
	if judgerInstance == nil {
		judgerInstance = &judger{}
		judgerInstance.getConn()
	}

	return judgerInstance
}

// getInstance - get instance from global variable, should be called after `initInstance`
func GetInstance() *judger {
	return judgerInstance
}

// closeInstance - cleanup the resource which be used by judger
func CloseInstance() {
	judgerInstance.anaConn.Close()
}

func (j *judger) SetOpt(opt int, param interface{}) error {
	if param == nil {
		return errors.New("Param should not be nil")
	}

	switch {
	case opt == OPT_BASEDIRECTORY:
		j.baseDirectory = param.(string)
		return nil
	case opt == OPT_SETENV:
		j.env = param.(string)
		return nil
	case opt == OPT_SETADDR:
		j.anaAddress = param.(string)
		return nil
	case opt == OPT_SETTEMPDIRECTORY:
		j.tempDirectory = param.(string)
	}
	return errors.New("unsupported option")
}

// submit - create new submit task synchronized or asynchronized by `go j.submit()`
// param:
// 		submitData -
func (j *judger) Submit(submitData SubmitData, callback SubmitCallback) {
	// step1. generate workspace
	workspacePath, err := ioutil.TempDir(j.tempDirectory, "")
	if err != nil {
		callback(submitData.Id, NewUndefinedError("create tempdir failed"))
		return
	}
	defer os.RemoveAll(workspacePath)

	postfixIndex := strings.Index(submitData.Language, ".")
	if postfixIndex == -1 {
		callback(submitData.Id, NewUndefinedError("wrong language type"))
		return
	}
	sourceFile := "source." + submitData.Language[:postfixIndex]
	buildPath := path.Join(workspacePath, "build")

	err = os.Mkdir(buildPath, 0755)
	if err != nil {
		glog.Errorf("create build dir %s failed, err: %v", buildPath, err)
		callback(submitData.Id, NewUndefinedError("create build dir failed"))
		return
	}

	f, err := os.Create(path.Join(buildPath, sourceFile))
	if err != nil {
		glog.Errorf("create source file failed, err: %v", err)
		callback(submitData.Id, NewUndefinedError("create sourceFile failed"))
		return
	}

	_, err = io.WriteString(f, submitData.Code)
	if err != nil {
		f.Close()
		glog.Errorf("write source code failed, err:%v", err)
		callback(submitData.Id, NewUndefinedError("write sourceCode failed"))
		return
	}
	f.Close()

	var buildScript *string
	buildScript = nil

	if submitData.BuildScript != "" {
		script := "build.sh"
		buildScript = &script
		err = CopyFile(submitData.BuildScript, path.Join(buildPath, script))
		if err != nil {
			glog.Errorf("create build script failed, err: %v", err)
			callback(submitData.Id, NewUndefinedError("create build script failed"))
			return
		}
	}

	config := TomlConfig{
		Source:      sourceFile,
		Language:    submitData.Language,
		BuildScript: buildScript,
		Timeout: TimeConfig{
			Seconds: 5,
			Nanos:   0,
		},
	}

	err = EncodeTomlFile(path.Join(buildPath, "config.toml"), config)
	if err != nil {
		glog.Errorf("encode build config toml failed, err: %v", err)
		callback(submitData.Id, NewUndefinedError("encode build config toml failed"))
		return
	}

	err = os.Symlink(submitData.RunnerConfig, path.Join(workspacePath, "config.toml"))
	if err != nil {
		glog.Errorf("copy runner script failed, err: %v", err)
		callback(submitData.Id, NewUndefinedError("create workspace config toml failed"))
		return
	}

	err = os.Symlink(path.Join(j.baseDirectory, j.env, strconv.FormatUint(submitData.Pid, 10), "problem"), path.Join(workspacePath, "problem"))
	if err != nil {
		glog.Errorf("link problem path failed, err: %v", err)
		callback(submitData.Id, NewUndefinedError("link problem path failed"))
		return
	}

	// step2. submit task to Ana
	i := 0
	for {
		err = j.submitTask(workspacePath, submitData.Id, callback)
		if err != nil {
			glog.Warningf("submit task failed, err: %v", err)
			if errStatus, ok := status.FromError(err); ok {
				// if grpc connection break, try to reconnect
				if errStatus.Code() == codes.Unavailable {
					glog.Warning("reconnect ana")
					j.getConn()
				}
			}
			if i == MAX_RETRY_TIMES {
				glog.Errorf("max connect retry, submit task failed, err: %v", err)
				callback(submitData.Id, NewUndefinedError(fmt.Sprintf("submit task failed, err: %v", err)))
				break
			}
			i++
		} else {
			break
		}
	}
}

func (j *judger) getConn() {
	for i := 0; i < MAX_RECONNECT_TIMES; i++ {
		conn, err := grpc.Dial(j.anaAddress, grpc.WithInsecure())
		if err != nil {
			glog.Warningf("connect failed: %v", err)
			continue
		}
		j.anaConn = conn
		break
	}

}

func (j *judger) submitTask(workspacePath string, id uint64, callback SubmitCallback) error {
	source := Workspace{
		Id:   &wrappers.StringValue{Value: strconv.FormatUint(id, 10) + j.env},
		Path: &wrappers.StringValue{Value: workspacePath},
	}

	rldata := JudgeResult{
		Time:       0,
		Memory:     0,
		Status:     "UE",
		Msg:        "",
		Case:       0,
		IsFinished: false,
	}

	resultData := &Report{
		Usage: &Resource{
			RealTime: &duration.Duration{Seconds: 0, Nanos: 0},
			CpuTime:  &duration.Duration{Seconds: 0, Nanos: 0},
			Memory:   &wrappers.UInt64Value{Value: 0},
		},
	}

	anaClient := NewAnaClient(j.anaConn)

	_, err := anaClient.JudgeWorkspace(context.Background(), &source)
	if err != nil {
		return err
	}

	request := Request{
		Id: &wrappers.StringValue{Value: strconv.FormatUint(id, 10) + j.env},
	}

	var caseNum uint64
	caseNum = 0

	for {
		resultData, err = anaClient.GetReport(context.Background(), &request)
		i := 0
		if err != nil {
			if errStatus, ok := status.FromError(err); ok {
				if errStatus.Code() == codes.OutOfRange {
					rldata.IsFinished = true
					break
				}

				// if grpc connection break, try to reconnect
				if errStatus.Code() == codes.Unavailable {
					glog.Warning("reconnect ana")
					j.getConn()
					i++
					if i == MAX_RETRY_TIMES {
						glog.Errorf("max connect retry, get report failed, err: %v", err)
						callback(id, NewUndefinedError(fmt.Sprintf("get report failed, err: %v", err)))
					} else {
						break
					}
				}
			}

			glog.Errorf("Recv msg from Ana error: %v\n", err)
			callback(id, NewUndefinedError(fmt.Sprintf("get report failed, err: %v", err)))
			return err
		}

		glog.Warningf("%d: rldata is %v", id, rldata)

		if rldata.Status == "AC" || rldata.Status == "UE" {
			rldata.Status = Report_ResultType_name[int32(resultData.Result)]
			rldata.Msg = resultData.Message.Value
			// 只记录第一次非AC样例数
			caseNum += 1
			rldata.Case = caseNum
		}

		if resultData.Result != Report_Accepted {
			rldata.Memory = 0
			rldata.Time = 0
		} else {
			rldata.Memory = Max(rldata.Memory, resultData.Usage.Memory.Value)
			rldata.Time = Max(rldata.Time, uint64(resultData.Usage.CpuTime.Seconds*1_000_000_000+int64(resultData.Usage.CpuTime.Nanos)))
		}

		callback(id, rldata)
	}

	callback(id, rldata)
	return nil
}
