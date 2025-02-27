package helper

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Log struct {
	Level     string
	Msg       string
	Source    string
	Timestamp time.Time
}

func NewLog(level, msg string) {
	switch level {
	case "INFO":
		log.Info().Msg(msg)
	case "WARN":
		log.Warn().Msg(msg)
	case "ERROR":
		log.Error().Msg(msg)
	case "FATAL":
		log.Fatal().Msg(msg)
	default:
		log.Info().Msg(msg)
	}

	fullLog := Log{
		Level: level,
		Msg: msg,
		Source: "auth-service",
		Timestamp: time.Now(),
	}

	produceToKafka(fullLog)
}

func produceToKafka(fullLog Log) {

}