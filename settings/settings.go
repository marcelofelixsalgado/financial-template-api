package settings

import (
	"log"
	"os"

	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/marcelofelixsalgado/financial-commons/settings"
)

// Common ConfigType struct to resolve env vars
type ConfigType struct {
	AppName     string `env:"APP_NAME" default:"financial-movement-api"`
	Environment string `env:"ENVIRONMENT" default:"development"`

	// Database Connection String (MySQL)
	DatabaseConnectionUser          string `env:"DATABASE_USER"`
	DatabaseConnectionPassword      string `env:"DATABASE_PASSWORD"`
	DatabaseConnectionServerAddress string `env:"DATABASE_SERVER_ADDRESS"`
	DatabaseConnectionServerPort    int    `env:"DATABASE_SERVER_PORT" default:"3306"`
	DatabaseName                    string `env:"DATABASE_NAME"`

	// HTTP Port to expose the API
	ApiHttpPort int `env:"API_PORT"`

	// Key used to sign the token
	SecretKey []byte `env:"SECRET_KEY"`

	ServerCloseWait int `env:"SERVER_CLOSEWAIT" default:"10"`

	// Log files
	LogAccessFile string `env:"LOG_ACCESS_FILE" default:"./access.log"`
	LogAppFile    string `env:"LOG_APP_FILE" default:"./app.log"`
	LogLevel      string `env:"LOG_LEVEL" default:"INFO"`

	// Kafka settings
	KafkaServerHost        string `env:"KAFKA_SERVER_HOST"`
	KafkaServerPort        int    `env:"KAFKA_SERVER_PORT"`
	KafkaDeliveryTimeout   int    `env:"KAFKA_DELIVERY_TIMEOUT" default:"0"`
	KafkaAcks              int    `env:"KAFKA_ACKS" default:"0"`                   // number of acknowledgements the leader broker must receive from ISR brokers before responding to the request: 0=Broker does not send any response/ack to client, -1 or all=Broker will block until message is committed by all in sync replicas
	KafkaEnableIdempotence string `env:"KAFKA_ENABLE_IDEMPOTENCE" default:"false"` // When set to true, the producer will ensure that messages are successfully produced exactly once and in the original produce order.
}

var Config ConfigType

// InitConfigs initializes the environment settings
func Load() {
	settings.Load()

	loadDotEnv()

	// bind env vars
	if err := env.Set(&Config); err != nil {
		log.Fatal(err)
	}

	if _, err := logrus.ParseLevel(Config.LogLevel); err != nil {
		log.Fatal(err)
	}
}

func loadDotEnv() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			logrus.Fatal("Error loading .env file")
		}
		logrus.Println("Using .env file")
	}
}
