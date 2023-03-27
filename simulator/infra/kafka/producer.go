package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"fmt"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"security.protocol": os.Getenv("KAFKA_BOOTSTRAP_SECURITY_PROTOCOL"),
		"sasl.mechanisms": os.Getenv("KAFKA_BOOTSTRAP_SASL_MECHANISMS"),
		"sasl.username": os.Getenv("KAFKA_BOOTSTRAP_SASL_USERNAME"),
		"sasl.password": os.Getenv("KAFKA_BOOTSTRAP_SASL_PASSWORD"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		fmt.Printf("Failed to create producer. message: %v\n", err)
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value: []byte(msg),
	}
  fmt.Printf("%+v\n", message)
	err :=  producer.Produce(message, nil)
	if err != nil {
		if err.(ckafka.Error).Code() == ckafka.ErrQueueFull {
			// Producer queue is full, wait 1s for messages
			// to be delivered then try again.
		}
		fmt.Printf("Failed to produce message: %v\n", err)
		return err
	}
	return nil
}