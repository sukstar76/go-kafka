package kafka

import (
	"time"
	"context"
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const(
	topic string  = "db-log"
	numParts int = 2
	replicationFactor int =  2
)


func SetTopic(){
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers" : "127.18.0.5:9092"})

	if err != nil{
		fmt.Printf("Failed to create Admin client : %s \n",err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timeout,err:= time.ParseDuration("60s")
	if err !=nil{
		panic("PaseDuration(60s)")
	}

	results, err := admin.CreateTopics(
		ctx,
		[]kafka.TopicSpecification{{
			Topic:  topic,
			NumPartitions: numParts,
			ReplicationFactor : replicationFactor}},
		kafka.SetAdminOperationTimeout(timeout))
	if err != nil{
		fmt.Printf("Failed to create topic")
	}
	for _, result := range results {
		fmt.Printf("%s\n", result)
	}
	admin.Close()

}
