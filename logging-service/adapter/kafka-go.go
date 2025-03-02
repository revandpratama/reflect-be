package adapter

import "github.com/segmentio/kafka-go"

type KafkaGo struct {
	// reader *kafka.Reader
}

var KafkaReader *kafka.Reader

func NewKafka() *KafkaGo {
	return &KafkaGo{}
}

func (k *KafkaGo) InitKafka() {
	KafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "logging-service",
		GroupTopics: []string{"auth-service", "main"},
	})
}
