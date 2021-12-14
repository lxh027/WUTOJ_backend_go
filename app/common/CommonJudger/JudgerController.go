package CommonJudger

import (
	"OnlineJudge/core/database"
	"log"

	"github.com/garyburd/redigo/redis"
)

var rc redis.Conn
var sub_judger redis.PubSubConn

func init() {
	rc = database.RedisClient.Get()
	sub_judger = redis.PubSubConn{Conn: rc}
}

func RunCommonJudger() error {
	// rc.Send("SUBSCRIBE", "message")
	// rc.Flush()
	log.Println("=================")
	err := sub_judger.Subscribe("panda")
	if err != nil {
		return err
	}
	go func() {
		for {
			// reply, err := rc.Receive()
			switch msg := sub_judger.Receive().(type) {
			case error:
				log.Println(msg)
			case redis.Message:
				log.Printf("this is %d: %v", 1, msg)
			}
		}
	}()
	return nil
}

func CloseCommonJudger() error {
	err := rc.Close()
	if err != nil {
		return err
	}
	err = sub_judger.Close()
	if err != nil {
		return err
	}
	return nil
}
