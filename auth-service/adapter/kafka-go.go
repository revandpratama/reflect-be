package adapter

import (
	"time"

	"github.com/revandpratama/reflect/auth-service/config"
	"github.com/revandpratama/reflect/auth-service/helper"
	"github.com/segmentio/kafka-go"
)

type KafkaGoOption struct {
	writer *kafka.Writer
}

func (k *KafkaGoOption) Start(a *Adapter) error {
	w := &kafka.Writer{
		Addr:         kafka.TCP(config.ENV.KafkaPort),
		Topic:        config.ENV.KafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 1 * time.Second,
		// AllowAutoTopicCreation: true,
	}

	k.writer = w
	a.KafkaGo = w
	helper.KafkaWriter = w

	return nil
}

func (k *KafkaGoOption) Stop() error {
	err := k.writer.Close()

	return err
}
