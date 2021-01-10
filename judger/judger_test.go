package judger

import (
	"math/rand"
	"testing"
	"time"
)

func createJudgerInstance() {
	instance := InitInstance()

	instance.SetOpt(OPT_SETENV, "master")
	instance.SetOpt(OPT_SETADDR, "127.0.0.1:8800")
	instance.SetOpt(OPT_BASEDIRECTORY, "/home/acmwhut/data")
	instance.SetOpt(OPT_SETTEMPDIRECTORY, "/tmpdir")
}

const ac_code = `
#include <stdio.h>

int main() {
	int a,b;
	scanf("%d%d",&a,&b);
	printf("%d\n", a + b);
	return 0;
}
`

const wa_code = `
#include <stdio.h>

int main() {
	int a,b;
	scanf("%d%d",&a,&b);
	printf("%d\n", a - b);
	return 0;
}
`

const tle_code = `
#include <stdio.h>

int main() {
	while(1) {}
	return 0;
}
`

const ce_code = `
#include <stdio.h>

int main() {
	int a, b;
`

const mle_code = `
#include <stdio.h>
#include <stdlib.h>

#define MAXN 40000000

int data[MAXN] = {0};

int main() {
	memset(data, 0, sizeof(data));
	return 0;
}
`

func TestSubmit(t *testing.T) {
	createJudgerInstance()
	defer CloseInstance()

	testCases := []struct {
		status 	   string
		sourceCode string
	} {
		{"AC", ac_code},
		{"WA", wa_code},
		{"CE", ce_code},
		{"TLE", tle_code},
		{"MLE", mle_code},
	}

	rand.Seed(time.Now().Unix())

	for index, testCase := range testCases {
		submitData := SubmitData{
			Id:           uint64(rand.Int()),
			Pid:          1000,
			Language:     "c.gcc",
			Code:         testCase.sourceCode,
			BuildScript:  "",
			RootfsConfig: nil,
		}

		ch := make(chan JudgeResult)

		// 这里使用chan传递结果是为了方便进行测试assert,
		// 实际使用中可以完全异步的将结果存放到数据库中
		callback := func(id uint64, result JudgeResult) {
			ch <- result
		}

		j := GetInstance()

		go j.Submit(submitData, callback)

		for {
			result := <-ch
			if result.IsFinished {
				if result.Status != testCase.status {
					t.Errorf("case %d,  expected %s(get %s),code `%s`, msg: %s", index, testCase.status, result.Status, testCase.sourceCode, result.Msg)
				}
				break
			}
		}
	}
}
