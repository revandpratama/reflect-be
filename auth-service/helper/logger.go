package helper

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
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
		Source: "auth-service",
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

func (l *Log) ToKafka() {
	fmt.Println("produced to kafka")
}