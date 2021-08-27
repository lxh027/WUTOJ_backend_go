package database

import (
	"OnlineJudge/config"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)


var RedisClient *redis.Pool

func init()  {
	redisConf := config.GetRedisConfig()
	RedisClient = &redis.Pool{
		MaxIdle: redisConf["maxIdle"].(int),
		MaxActive:   redisConf["maxActive"].(int),
		IdleTimeout: redisConf["timeout"].(time.Duration),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(redisConf["type"].(string), redisConf["host"].(string))
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}
			/*if _, err := c.Do("AUTH", redisConf["auth"].(string)); err != nil {
				_ = c.Close()
				fmt.Println(err.Error())
				return nil, err
			}*/
			return c, nil
		},
	}
}

func ZAddToRedis(key string, score int64, member interface{}) error  {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("ZADD", key, score, member)
	return err
}

func ZGetAllFromRedis(key string) (interface{}, error)  {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	return rc.Do("ZRANGE", key, 0, -1)
}

func SAddToRedisSet(key string, member interface{}) error  {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("SADD", key, member)
	return err
}

func SIsNumberOfRedisSet(key string, member interface{}) (bool, error)  {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	value, err := redis.Bool(rc.Do("SISMEMBER", key, member))
	return value, err
}

func GetFromRedis(key string) (interface{}, error)  {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	value, err := rc.Do("GET", key)
	return value, err
}

func PutToRedis(key string, value interface{}, timeout int)  error {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("SET", key, value, "EX", timeout)
	return err
}

func DeleteFromRedis(key string) error {
	key = appendPrefix(key)
	rc := RedisClient.Get()
	defer rc.Close()
	_, err := rc.Do("DEL", key)
	return err
}

func appendPrefix(key string) string {
	prefix := config.GetRedisConfig()["env"].(string)
	return prefix+"."+key
}


