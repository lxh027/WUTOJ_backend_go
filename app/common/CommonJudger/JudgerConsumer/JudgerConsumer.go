package JudgerConsumer

import (
	"OnlineJudge/core/database"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var rc_consumer redis.Conn
var Running chan (bool)

func init() {
	Running = make(chan bool)
	rc_consumer = database.RedisClient.Get()
}

// SubmitInfo
// Key : User ID + Problem ID + Submit ID
// Value : language Type + Code
type SubmitInfo struct {
	Key   string
	Value string
}

type Msg struct {
	ID          string
	SubmitInfos []SubmitInfo
}

type Info struct {
	StreamName string
	Msgs       []Msg
}

// 解析redis查询结果为所需结构体
func ParseRedisValue(values []interface{}) (Info, error) {
	var ResponseInfo Info
	StreamNameBytes := values[0].([]interface{})[0]
	// var cache interface{}
	StreamNameStr, err := redis.String(StreamNameBytes, nil)
	if err != nil {
		log.Println(err)
	}
	ResponseInfo.StreamName = StreamNameStr
	MsgBytes := values[0].([]interface{})[1].([]interface{})

	var Msgs []Msg
	Msgs = make([]Msg, 0)
	for idx, val := range MsgBytes {
		MsgIDInterface := val.([]interface{})[0]
		MsgValue := val.([]interface{})[1].([]interface{})
		var SubmitInfos []SubmitInfo
		SubmitInfos = make([]SubmitInfo, 0)
		log.Println(idx)
		MsgID, err := redis.String(MsgIDInterface, nil)
		if err != nil {
			log.Println(err)
		}
		KeyAndValue := MsgValue

		submitKey, err := redis.String(KeyAndValue[0], nil)
		if err != nil {
			log.Println(err)
		}
		submitValue, err := redis.String(KeyAndValue[1], err)
		if err != nil {
			log.Println(err)
		}
		SubmitInfos = append(SubmitInfos, SubmitInfo{
			Key:   submitKey,
			Value: submitValue,
		})

		Msgs = append(Msgs, Msg{
			ID:          MsgID,
			SubmitInfos: SubmitInfos,
		})
	}

	ResponseInfo.Msgs = Msgs

	if err != nil {
		log.Println(err)
	}
	return ResponseInfo, nil
}

func RunJudgerConsumer() error {
	var StopFlag bool
	go func() {
		for {
			// check the symbol of channel
			select {
			case StopFlag = <-Running:
				if !StopFlag {
					log.Println("Stop Consumer Success")
					return
				}
			default:
				break
			}

			time.Sleep(time.Second * 2)
			reply, err := redis.Values(rc_consumer.Do("xread", "count", 1, "STREAMS", "pandastream", "0-0"))
			if err != nil {
				log.Println("nothing to do ", err)
				continue
			}
			var info Info
			info, _ = ParseRedisValue(reply)
			for _, val := range info.Msgs {
				rc_consumer.Do("xdel", info.StreamName, val.ID)
				for _, submit := range val.SubmitInfos {
					ok, err := ConsumeMsg(submit)
					if err != nil {

					}
					if ok {

					}
				}
			}
			log.Println(info)
		}
	}()
	return nil
}

func ConsumeMsg(submit SubmitInfo) (bool, error) {
	return true, nil
	// splitValue := strings.Split(submit.Value, redis_key.SplitString)
	// LanguageType, err := strconv.Atoi(splitValue[0])
	// if err != nil {
	// 	log.Println(err)
	// }
	// SubmitCode := splitValue[1]
	// instance := judger.GetInstance()

	// callback := func() {

	// }

	// instance.Submit()
}

func CloseJudgerConsumer() error {
	err := rc_consumer.Close()
	Running <- false
	// close(Running)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	RunJudgerConsumer()
	defer CloseJudgerConsumer()
}
