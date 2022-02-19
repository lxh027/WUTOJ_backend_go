package nsqueue

import (
	"OnlineJudge/app/panel/model"
	cfg "OnlineJudge/config"
	"OnlineJudge/constants"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

type ResponceHandler struct {
	q *nsq.Consumer
}

type Producer struct {
	producer *nsq.Producer
}

var RequestProducer *Producer

func (h *ResponceHandler) HandleMessage(message *nsq.Message) error {
	type Data struct {
		Msg Response
	}

	var (
		data *Data
		err  error
	)
	data = &Data{}
	buf := bytes.NewBuffer(message.Body)
	dec := gob.NewDecoder(buf)

	err = dec.Decode(&data.Msg)
	if err != nil {
		log.Println("Error decoding GOB data:", err)
		return err
	}
	// log.Println(data.Msg)
	SaveOJUserData(data.Msg)

	message.Finish()
	return nil
}

func InitConsumer(topic string, channel string) {
	var (
		config *nsq.Config
		h      *ResponceHandler
		err    error
	)
	h = &ResponceHandler{}

	config = nsq.NewConfig()
	nsqConfig := cfg.GetNSQConfig()
	if h.q, err = nsq.NewConsumer(topic, channel, config); err != nil {
		log.Println(err)
		return
	}

	h.q.AddHandler(h)
	if err = h.q.ConnectToNSQD(fmt.Sprintf("%s:%s", nsqConfig["host"], nsqConfig["port"])); err != nil {
		log.Println(err)
	}

}

//InitNSQ 初始化nsq
func InitNSQ(channel string) {
	nsqConfig := cfg.GetNSQConfig()
	InitConsumer(fmt.Sprintf("%s", nsqConfig["consumer_topic"]), channel)
	RequestProducer = &Producer{}
	// RequestProducer.producer, _ = InitProducer(cfg.nsqConfig["host"] + ":" + cfg.nsqConfig["port"])
	RequestProducer.producer, _ = InitProducer(fmt.Sprintf("%s:%s", nsqConfig["host"], nsqConfig["port"]))
}

func (p *Producer) Publish(data Request) (err error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err = enc.Encode(data)
	if err != nil {
		log.Println(err)
		return err
	}

	if err = p.producer.Publish(fmt.Sprintf("%s", cfg.GetNSQConfig()["producer_topic"]), buf.Bytes()); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func InitProducer(addr string) (p *nsq.Producer, err error) {
	// var config *nsq.Config
	config := nsq.NewConfig()
	if p, err = nsq.NewProducer(addr, config); err != nil {
		return nil, err
	}
	return p, nil
}

//SaveOJUserData 储存爬虫爬取的提交信息
func SaveOJUserData(res Response) {
	var userDatas []model.OJWebUserData
	var ojWebUserData model.OJWebUserData
	ojWebUserModel := model.OJWebUserData{}
	for _, data := range res.Data {
		var err error
		ojWebUserData.UserID, err = strconv.Atoi(data.UserInfo)
		if err != nil {
			log.Println("crawler data error:", err)
			return
		}
		for key, value := range data.Data {
			ojWebUserData.OJName = key
			solvedData := value.Data
			for i := range solvedData.Problems {
				submitTime, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", solvedData.Problems[i].SubmitTime)
				if err != nil {
					log.Println("crawler data error:", err)
					continue
				}
				ojWebUserData.SubmitTime = submitTime
				ojWebUserData.Status = solvedData.Problems[i].StatusWord
				ojWebUserData.ProblemID = solvedData.Problems[i].ProblemTitle
				userDatas = append(userDatas, ojWebUserData)
			}
		}
	}
	re := ojWebUserModel.AddOJWebUserDatas(userDatas)
	if re.Status != constants.CodeSuccess {
		log.Println("crawler data error:", re.Data)
		return
	}
	log.Println("receive the crawler data successful")
}
