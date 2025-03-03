package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	SecretKey string

	GRPCServerPort string
	RESTServerPort string

	KafkaHost     string
	KafkaPort     string
	KafkaClientID string
	KafkaTopic    string
}

var ENV *Config

func LoadConfig() {
	ENV = &Config{
		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBName:     os.Getenv("DBName"),
		DBSSLMode:  os.Getenv("DBSSLMode"),

		SecretKey: os.Getenv("SecretKey"),

		GRPCServerPort: os.Getenv("GRPCServerPort"),
		RESTServerPort: os.Getenv("RESTServerPort"),

		KafkaHost:     os.Getenv("KafkaHost"),
		KafkaPort:     os.Getenv("KafkaPort"),
		KafkaClientID: os.Getenv("KafkaClientID"),
		KafkaTopic:    os.Getenv("KafkaTopic"),
	}
}
