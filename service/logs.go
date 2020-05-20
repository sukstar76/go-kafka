package service

import (
	"fmt"
	"github.com/sukstar76/go-kafka/kafka"
	"github.com/sukstar76/go-kafka/kafka/message"
	ckafka "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const LogTopic string = "db-log"

type LogServiceInterface interface {
	Send(string, message.LogMessage) error
}
type LogService struct {
	producer *kafka.Producertype
}

func NewLogService(producer *kafka.Producertype) *LogService {
	return &LogService{producer: producer}
}

func (ls *LogService) Send(topic string, msg message.LogMessage) error {
	var deliveryChan = make(chan ckafka.Event)
	defer close(deliveryChan)

	message.SetMessage(msg)

	err := ls.producer.Send(topic, message.Message, deliveryChan)

	if err != nil {
		return err
	}

	e := <-deliveryChan
	m := e.(*ckafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: &v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		return nil
	}
}
