package judger

import (
	"flag"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func createJudgerInstance() {
	instance := InitInstance()

	instance.SetOpt(OPT_SETENV, "dev")
	instance.SetOpt(OPT_SETADDR, "127.0.0.1:8800")
	instance.SetOpt(OPT_BASEDIRECTORY, "/home/acmwhut/data")
	instance.SetOpt(OPT_SETTEMPDIRECTORY, "/home/ana_tmpdir")
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

const py_ac_code = `
(a, b) = map(int, input().split())
print(a + b)
`

const java_ac_code = `
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int a = scan.nextInt();
        int b = scan.nextInt();
        System.out.println(a + b);
    }
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

const py_wa_code = `
(a, b) = map(int, input().split())
print(a-b)
`

const java_wa_code = `
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int a = scan.nextInt();
        int b = scan.nextInt();
        System.out.println(a - b);
    }
}
`

const tle_code = `
#include <stdio.h>

int main() {
	while(1) {}
	return 0;
}
`

const py_tle_code = `
while True:
    pass
`

const java_tle_code = `
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        while (true) {
        }
    }
}
`

const ce_code = `
#include <stdio.h>

int main() {
	int a, b;
`

const java_ce_code = `
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        }
    }
}
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

const py_mle_code = `
list(range(int(1e7)))
`

const java_mle_code = `
import java.util.*;

public class Main {
    public static void main(String[] args) {
        Vector v= new Vector(100000000, 3);
    }
}
`

const py_re_code = `
print(a+b)
`

const java_re_code = `
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);
        int a = scan.nextInt();
        int b = scan.nextInt();
        System.out.println(a + b / 0);
    }
}
`

func TestSubmit(t *testing.T) {
	createJudgerInstance()
	defer CloseInstance()

	langBasePath := "/home/env"

	langConfigs := []struct {
		lang         string
		buildSh      string
		runnerConfig string
	}{
		{"c.gcc", "", langBasePath + "/c.gcc/runner.toml"},
		{"python.cpython3.6", langBasePath + "/python.cpython3.6/build.sh", langBasePath + "/python.cpython3.6/runner.toml"},
		{"java.openjdk8", langBasePath + "/java.openjdk8/build.sh", langBasePath + "/java.openjdk8/runner.toml"},
	}

	testCases := []struct {
		langIndex  int
		status     string
		sourceCode string
	}{
		{0, "AC", ac_code},
		{0, "WA", wa_code},
		{0, "CE", ce_code},
		{0, "TLE", tle_code},
		{0, "MLE", mle_code},
		{1, "AC", py_ac_code},
		{1, "WA", py_wa_code},
		{1, "RE", py_re_code},
		{1, "TLE", py_tle_code},
		{1, "MLE", py_mle_code},
		{2, "AC", java_ac_code},
		{2, "WA", java_wa_code},
		{2, "CE", java_ce_code},
		{2, "RE", java_re_code},
		{2, "TLE", java_tle_code},
		{2, "MLE", java_mle_code},
	}

	rand.Seed(time.Now().Unix())

	for index, testCase := range testCases {
		langConfig := langConfigs[testCase.langIndex]
		submitData := SubmitData{
			Id:           uint64(rand.Int()),
			Pid:          24,
			Language:     langConfig.lang,
			Code:         testCase.sourceCode,
			BuildScript:  langConfig.buildSh,
			RunnerConfig: langConfig.runnerConfig,
		}

		ch := make(chan JudgeResult, 100)

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

func TestPressure(t *testing.T) {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "/tmp")
	flag.Set("v", "3")
	// flag.Parse()
	createJudgerInstance()
	defer CloseInstance()

	langBasePath := "/home/env"

	langConfigs := []struct {
		lang         string
		buildSh      string
		runnerConfig string
	}{
		{"c.gcc", "", langBasePath + "/c.gcc/runner.toml"},
		{"python.cpython3.6", langBasePath + "/python.cpython3.6/build.sh", langBasePath + "/python.cpython3.6/runner.toml"},
		{"java.openjdk8", langBasePath + "/java.openjdk8/build.sh", langBasePath + "/java.openjdk8/runner.toml"},
	}

	testCases := []struct {
		langIndex  int
		status     string
		sourceCode string
	}{}

	for i := 0; i < 4; i++ {
		testCases = append(testCases, struct {
			langIndex  int
			status     string
			sourceCode string
		}{
			0,
			"AC",
			ac_code,
		})
	}

	wg := sync.WaitGroup{}

	rand.Seed(time.Now().Unix())

	times := []uint64{}

	for index, testCase := range testCases {
		langConfig := langConfigs[testCase.langIndex]
		submitData := SubmitData{
			Id:           uint64(rand.Int()),
			Pid:          24,
			Language:     langConfig.lang,
			Code:         testCase.sourceCode,
			BuildScript:  langConfig.buildSh,
			RunnerConfig: langConfig.runnerConfig,
		}

		// 这里使用chan传递结果是为了方便进行测试assert,
		// 实际使用中可以完全异步的将结果存放到数据库中
		callback := func(id uint64, result JudgeResult) {
			if result.IsFinished {
				if result.Status != testCase.status {
					t.Errorf("case %d,  expected %s(get %s),code `%s`, msg: %s", index, testCase.status, result.Status, testCase.sourceCode, result.Msg)
				}
				times = append(times, result.Time)
				wg.Done()
			}
		}

		j := GetInstance()

		wg.Add(1)
		go j.Submit(submitData, callback)
	}

	wg.Wait()
	t.Logf("avg time is %d, p99 time is %d", avg(times), p99(times, t))
}

func avg(vec []uint64) uint64 {
	var sum uint64 = 0
	for _, data := range vec {
		sum += data
	}
	ans := sum / uint64(len(vec))
	return ans
}

func p99(vec []uint64, t *testing.T) uint64 {
	tmp := uint64(float64(len(vec)) * 0.99)
	p99 := uint64(len(vec) - 1)
	if tmp < p99 {
		p99 = tmp
	}
	sort.Slice(vec, func(i, j int) bool {
		return vec[i] < vec[j]
	})
	t.Logf("vec is %v", vec)
	return vec[p99]
}
