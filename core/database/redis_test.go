package database

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func Test(t *testing.T) {
	_ = PutToRedis("k", "132", 3600)
	_ = PutToRedis("q", 123, 3600)
	m := map[string]string {
		"asd": "asd",
		"as": "qq",
	}
	jsonMap, _ := json.Marshal(m)
	_ = PutToRedis("m", jsonMap, 3600)
	s := struct {
		AA int
		BB string
	}{1, "aa"}
	_ = PutToRedis("s", s, 3600)

	if k, err := GetFromRedis("m"); err == nil {
		mp := make(map[string]interface{})
		mpBytes, _ := redis.String(k, err)
		_ = json.Unmarshal([]byte(mpBytes), &mp)
		fmt.Println(mpBytes)
	} else {
		fmt.Println(err.Error())
	}

}

func TestSortedSet(t *testing.T) {
	m1 := map[string]interface{} {
		"asdaqsd": 123,
		"asd": map[string]interface{}{
			"ddd": "ddd",
		},
	}
	m2 := map[string]interface{} {
		"asdaqsd": 234,
		"asd": map[string]interface{}{
			"ddd": "555",
		},
	}
	jsonM1, _ := json.Marshal(m1)
	jsonM2, _ := json.Marshal(m2)
	_ = ZAddToRedis("rank", 123, jsonM1)
	_ = ZAddToRedis("rank", 123, jsonM2)

	if k, err := ZGetAllFromRedis("rank"); err == nil {
		//var mp []interface{}
		mpBytes, _ := redis.Strings(k, err)
		//_ = json.Unmarshal(mpBytes, &mp)
		fmt.Println(mpBytes)
	} else {
		fmt.Println(err.Error())
	}

}