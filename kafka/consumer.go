package kafka

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"os"
	"os/signal"
	"syscall"
)

type ConsumerConnector struct{
	consumer *kafka.Consumer
}



func ConsumerConnect() (*ConsumerConnector,error){
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":"127.18.0.5:9092",
		"broker.address.family":"v4",
		"group.id": "0",
		"session.timeout.ms": 6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"enable.partition.eof": true,
		"auto.offset.reset": "earliest"})

	if err !=nil{
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n",c)

	return &ConsumerConnector{
		consumer : c,
	},nil
}

func(c *ConsumerConnector) Consume() error {
	sigchan := make(chan os.Signal,1)
	signal.Notify(sigchan,syscall.SIGINT,syscall.SIGTERM)
	err := c.consumer.SubscribeTopics([]string{"db-log"},nil)
	if err!=nil{
		return err
	}
	defer c.consumer.Close()
	run := true

	for run == true{
		select{
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			case ev := <-c.consumer.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.consumer.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.consumer.Unassign()
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	return nil
}
