package handler

import (
	"sync"

	"github.com/marcelofelixsalgado/financial-commons/pkg/events"
	"github.com/marcelofelixsalgado/financial-commons/pkg/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/marcelofelixsalgado/financial-commons/pkg/commons/logger"
)

type TemplateCreatedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewTemplateCreatedKafkaHandler(kafka *kafka.Producer) *TemplateCreatedKafkaHandler {
	return &TemplateCreatedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *TemplateCreatedKafkaHandler) Handle(message events.IEvent, wg *sync.WaitGroup) {
	logger := logger.GetLogger()

	defer wg.Done()

	deliveryChan := make(chan ckafka.Event)

	go h.Kafka.Publish(message, nil, "templates", deliveryChan)

	event := <-deliveryChan
	channelMessage := event.(*ckafka.Message)

	if channelMessage.TopicPartition.Error != nil {
		logger.Errorf("Erro sending the message to Kafka - Topic: %s / Partition: %s", channelMessage.TopicPartition.Topic, channelMessage.TopicPartition)
	}

	logger.Debug("TemplateCreatedKafkaHandler: ", channelMessage.TopicPartition)
	logger.Debug("TemplateCreatedKafkaHandler: ", message.GetPayload())
}
