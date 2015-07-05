package main

import (
	"log"

	"github.com/Shopify/sarama"
)

const (
	topic     = "my-messages"
	partition = 0
)

var kafkaAddrs = []string{"localhost:9092"}

func main() {

	log.Println("Starting ... ")

	asyncProducer, err := sarama.NewAsyncProducer(kafkaAddrs, sarama.NewConfig())

	if err != nil {
		log.Fatalf("cannot create the AsyncProducer")
	}

	asyncProducer.Input() <- &sarama.ProducerMessage{Topic: topic}

	//asyncProducer.
}
