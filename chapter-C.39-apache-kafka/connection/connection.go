package connection

import "github.com/IBM/sarama"

func ConnectToProducer(urls []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 6

	conn, err := sarama.NewSyncProducer(urls, config)
	if err != nil {
		return nil, err
	}

	return conn, err
}

//
func ConnectToConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
