package main

import (
	"tutorial-go-kafka/connection"

	"github.com/IBM/sarama"
)

func PullFromProducer(topic string) (sarama.Consumer, sarama.PartitionConsumer) {

	worker, err := connection.ConnectToConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	return worker, consumer
}
