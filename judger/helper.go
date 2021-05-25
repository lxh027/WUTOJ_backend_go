package judger

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io"
	"os"
)

func NewUndefinedError(msg string) JudgeResult {
	return JudgeResult{
		Status:     "UE",
		Time:       0,
		Memory:     0,
		Msg:        msg,
		IsFinished: true,
	}
}

func Max(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func CopyFile(srcPath string, dstPath string) error {
	srcFp, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFp.Close()

	dstFp, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFp.Close()

	_, err = io.Copy(dstFp, srcFp)
	return err
}

func EncodeTomlFile(filePath string, data interface{}) error {
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	err := e.Encode(data)
	if err != nil {
		return err
	}

	fp, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	return nil
}
