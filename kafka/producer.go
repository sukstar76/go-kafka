package kafka

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type Producertype struct {
	Producer *kafka.Producer
}

func ConnectProducer() *Producertype {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "127.18.0.8:9092"})
	if err != nil {
		fmt.Printf("Failed to create producer: %s \n", err)
	}

	return &Producertype{
		Producer: p,
	}
}

func (p *Producertype) Send(topic string, msg []byte, deliveryChan chan kafka.Event) error {
	return p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, deliveryChan)
}
