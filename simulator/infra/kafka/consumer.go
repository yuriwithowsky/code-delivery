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

func(k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id": os.Getenv("KAFKA_CONSUMER_GROUP_ID"),
		"security.protocol": os.Getenv("KAFKA_BOOTSTRAP_SECURITY_PROTOCOL"),
		// "sasl.mechanisms": os.Getenv("KAFKA_BOOTSTRAP_SASL_MECHANISMS"),
		// "sasl.username": os.Getenv("KAFKA_BOOTSTRAP_SASL_USERNAME"),
		// "sasl.password": os.Getenv("KAFKA_BOOTSTRAP_SASL_PASSWORD"),
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	topics := []string{os.Getenv("KAFKA_READ_TOPIC")}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}