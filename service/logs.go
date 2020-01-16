package service

import(
	"fmt"
	"github.com/sukstar76/kafka"
	"github.com/sukstar76/kafka/message"
	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)


const LogTopic string ="db-log"

var LogService *kafka.Producertype

func Send(topic string, msg message.LogMessage) error{
	var deliveryChan = make(chan ckafka.Event)
	defer close(deliveryChan)

	message.SetMessage(msg)


	err := LogService.Send(topic,message.Message,deliveryChan)

	if err!= nil{
		return err
	}

	e :=<-deliveryChan
	m := e.(*ckafka.Message)

	if m.TopicPartition.Error != nil{
		fmt.Printf("Delivery failed: &v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	} else{
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic,  m.TopicPartition.Partition, m.TopicPartition.Offset)
		return nil
	}
}

