package main

import (
	"fmt"
	"tutorial-go-kafka/connection"

	"github.com/IBM/sarama"
)

func PushPizzaQueue(topic string, message []byte) error {
	urls := []string{"localhost:9092"}
	producer, err := connection.ConnectToProducer(urls)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
