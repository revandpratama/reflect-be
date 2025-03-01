package helper

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Log struct {
	Level     string
	Msg       string
	Source    string
	Timestamp time.Time
}

func NewLog() *Log {
	return &Log{
		Timestamp: time.Now(),
		Source:    "auth-service",
	}
}

func (l *Log) Info(msg string) *Log {
	l.Level = "INFO"
	l.Msg = msg

	log.Info().Msg(msg)
	return l
}
func (l *Log) Fatal(msg string) *Log {
	l.Level = "FATAL"
	l.Msg = msg

	log.Fatal().Msg(msg)
	return l
}
func (l *Log) Error(msg string) *Log {
	l.Level = "ERROR"
	l.Msg = msg

	log.Error().Msg(msg)
	return l
}

var KafkaWriter *kafka.Writer

func (l *Log) ToKafka() {

	err := KafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("failed to write messages: %v ", err))
	}

	fmt.Println("wrote messages")
}
