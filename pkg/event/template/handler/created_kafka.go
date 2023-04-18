package handler

import (
	"sync"

	"github.com/marcelofelixsalgado/financial-commons/pkg/events"
	"github.com/marcelofelixsalgado/financial-commons/pkg/kafka"

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
	h.Kafka.Publish(message, nil, "Templates")
	logger.Debug("TemplateCreatedKafkaHandler: ", message.GetPayload())
}
