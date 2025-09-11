package eventdriven

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func RunKafka() {
	// topic <- message <- producer
	// producer -> message ->topic
	// consumer <- message <- topic
	// data - message

	//producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:8080"})

	if err != nil {
		log.Fatal(err)
	}

	defer producer.Close()

	topic := "test_topic"

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte("Hello kafka, I am using it for Event driven API"),
	}, nil)

	//consumer

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:8080",
		"group.id":          "test-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	consumer.SubscribeTopics([]string{topic}, nil)

	msg, err := consumer.ReadMessage(-1)

	if err == nil {
		fmt.Println("Message: ", string(msg.Value))
	} else {
		log.Fatal(err)
	}
}
