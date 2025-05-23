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
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := KafkaWriter.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte(l.Source),
				Value: []byte(l.Msg),
				Time:  l.Timestamp,
			},
		)
		if err != nil {
			log.Error().Msg(fmt.Sprintf("failed to write messages: %v ", err))
			return
		}
		log.Info().Msgf("write message to kafka: %v", l.Msg)
	}()
}
