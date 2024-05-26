package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	topic := "Pizza"

	worker, consumer := PullFromProducer(topic)

	log.Println("Consumer started")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	messageCount := 0

	doneChan := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				messageCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", messageCount, string(msg.Topic), string(msg.Value))
			case <-sigChan:
				fmt.Println("Interrupt is detected")
				doneChan <- struct{}{}
			}
		}
	}()

	<-doneChan
	fmt.Println("Processed", messageCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}
}
