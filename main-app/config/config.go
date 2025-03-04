package config

import (
	"os"
)

type Config struct {
	AppEnvironment string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	SecretKey string

	GRPCServerPort string
	RESTServerPort string

	Minio_Endpoint  string
	Minio_AccessKey string
	Minio_SecretKey string
	Minio_UseSSL    string

	KafkaHost     string
	KafkaPort     string
	KafkaClientID string
	KafkaTopic    string
}

var ENV *Config

func LoadConfig() {
	ENV = &Config{
		AppEnvironment: os.Getenv("AppEnvironment"),

		DBHost:     os.Getenv("DBHost"),
		DBPort:     os.Getenv("DBPort"),
		DBUser:     os.Getenv("DBUser"),
		DBPassword: os.Getenv("DBPassword"),
		DBName:     os.Getenv("DBName"),
		DBSSLMode:  os.Getenv("DBSSLMode"),

		SecretKey: os.Getenv("SecretKey"),

		GRPCServerPort: os.Getenv("GRPCServerPort"),
		RESTServerPort: os.Getenv("RESTServerPort"),

		Minio_Endpoint:  os.Getenv("Minio_Endpoint"),
		Minio_AccessKey: os.Getenv("Minio_AccessKey"),
		Minio_SecretKey: os.Getenv("Minio_SecretKey"),
		Minio_UseSSL:    os.Getenv("Minio_UseSSL"),

		KafkaHost:     os.Getenv("KafkaHost"),
		KafkaPort:     os.Getenv("KafkaPort"),
		KafkaClientID: os.Getenv("KafkaClientID"),
		KafkaTopic:    os.Getenv("KafkaTopic"),
	}
}
