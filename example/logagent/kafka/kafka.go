package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type logData struct {
	topic string
	data string
}

var (
	client sarama.SyncProducer
	logDataChan chan *logData
)

func Init(addrs []string, maxSize int)(err error){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	// 连接 kafka
	client, err = sarama.NewAsyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, error:", err)
		return
	}
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine 从通道中取数据发往kafka
	go sendToKafka()
	return
}

func SendToChan(topic, data string)  {
	msg := &logData{
		topic:topic,
		data:data,
	}
	logDataChan <- msg
}

func sendToKafka()  {
	for  {
		select {
		case ld := <- logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, error:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond*50)
		}
	}
}
