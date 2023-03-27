package kafka

import (
	"encoding/json"
	"github.com/yuriwithowsky/full-cycle-code-delivery-simulator/infra/kafka"
	route2 "github.com/yuriwithowsky/full-cycle-code-delivery-simulator/application/route"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
	"os"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	if err := json.Unmarshal(msg.Value, &route); err != nil {
		log.Println(err.Error())
	}
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KAFKA_PRODUCE_TOPIC"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}