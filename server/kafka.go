package server

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func log(message string) error {
	producer, err := sarama.NewSyncProducer([]string{"localhost:19092", "localhost:29092", "localhost:39092"}, nil)

	if err != nil {
		return err
	}

	defer func() {
		if errP := producer.Close(); errP != nil {
			fmt.Println(errP)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "test_melkor", Value: sarama.StringEncoder(message)}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("> message sent to partition %d at offset %d\n", partition, offset)

	return nil
}
