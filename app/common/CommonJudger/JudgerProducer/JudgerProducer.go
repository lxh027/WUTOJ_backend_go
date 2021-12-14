package JudgerProducer

import (
	"OnlineJudge/app/panel/model"
	"OnlineJudge/core/database"
	"log"

	"github.com/garyburd/redigo/redis"
)

var rc_producer redis.Conn

func init() {
	rc_producer = database.RedisClient.Get()
}

func AddSubmitToStream(submit model.Submit) {
	value := string(submit.Language) + submit.SourceCode
	streamID, err := redis.String(rc_producer.Do("xadd", "pandastream", "*", submit.UserID, value))
	if err != nil {
		log.Println(err)
	}
	log.Println(streamID)
}

func main() {

}
