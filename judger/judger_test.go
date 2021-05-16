package judger

import (
	"math/rand"
	"testing"
	"time"
)

func createJudgerInstance() {
	instance := InitInstance()

	instance.SetOpt(OPT_SETENV, "dev")
	instance.SetOpt(OPT_SETADDR, "127.0.0.1:8800")
	instance.SetOpt(OPT_BASEDIRECTORY, "/home/baka233/acmwhut/data")
	instance.SetOpt(OPT_SETTEMPDIRECTORY, "/home/baka233/acmwhut/ana_tmpdir")
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

	pyCommand := "/usr/local/bin/python"
	javaCommand := "java"

	langConfigs := []struct {
		lang    string
		rootfs  *RootfsConfig
		buildSh string
		command *string
		args    *[]string
		envs    *map[string]string
	}{
		{"c.gcc", nil, "", nil, nil, nil},
		{"py.cpython3.6",
			&RootfsConfig{
				BasePath: "/home/baka233/acmwhut/env/py.cpython3.6/lang_runtime",
				WithProc: false},
			"/home/baka233/acmwhut/env/py.cpython3.6/build.sh",
			&pyCommand,
			&[]string{"main.py"},
			&map[string]string{},
		},
		{"java.openjdk-10",
			&RootfsConfig{
				BasePath: "/home/baka233/acmwhut/env/java.openjdk-10/lang_runtime",
				WithProc: true,
			},
			"/home/baka233/acmwhut/env/java.openjdk-10/build.sh",
			&javaCommand,
			&[]string{
				"Main",
			},
			&map[string]string{
				"JAVA_HOME":           "/docker-java-home",
				"JAVA_VERSION":        "10.0.2",
				"JAVA_DEBIAN_VERSION": "10.0.2+13-2",
				"PATH":                "/opt/java/openjdk/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/opt/openjdk-10/bin",
			},
		},
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
		runnerConfig := RunnerConfig{
			Runner: Runner{
				Language: langConfig.lang,
				Command:  langConfig.command,
				Args:     langConfig.args,
				Envs:     langConfig.envs,
				Rootfs:   langConfig.rootfs,
			},
		}
		submitData := SubmitData{
			Id:          uint64(rand.Int()),
			Pid:         24,
			Language:    langConfig.lang,
			Code:        testCase.sourceCode,
			BuildScript: langConfig.buildSh,
			Runner:      runnerConfig,
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
