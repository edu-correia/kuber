package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (kafkaConsumer *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id": os.Getenv("KafkaConsumerGroupId"),
	}

	consumer, err := ckafka.NewConsumer(configMap)
	if (err != nil) {
		log.Fatalf("Error consuming kafka message: " + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started")

	for {
		message, err := consumer.ReadMessage(-1)
		fmt.Println("message: " + string(message.Value))
		if (err == nil) {
			kafkaConsumer.MsgChan <- message
		}
	}

}

