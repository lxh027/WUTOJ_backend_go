package helper

import (
	"crypto/md5"
	"fmt"
)

func GetMd5(str string) string {
	hash := md5.Sum([]byte(str))
	hashBase64 := fmt.Sprintf("%x", hash)
	return hashBase64
}
