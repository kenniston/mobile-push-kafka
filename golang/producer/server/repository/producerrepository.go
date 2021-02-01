package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/dto"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type ProducerRepository interface {
	Send(message dto.PushMessage) error
}

//===============================================================================
// ProducerRepository is a structure to integration with others
// systems (internal and external)
// Methods for this structure MUST ALWAYS return DTO Entities and/or error
//
type producerRepository struct {
	framework.BaseRepository
	conn *kafka.Conn
}

// Create and initialize the repository
func NewSecurityRepository(config *viper.Viper) ProducerRepository {
	return &producerRepository{
		BaseRepository: framework.NewBaseRepository("", "", "Producer Repository", config),
	}
}

// Return a Security list object
func (r *producerRepository) Send(message dto.PushMessage) error {
	r.ChecksInitialized()

	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	logrus.Debug("Message: %s", string(msg))

	address := r.GetConfig().GetString("kafka-address")
	topic := r.GetConfig().GetString("kafka-topic")

	kafkaWriter :=  &kafka.Writer{
		Addr:     kafka.TCP(address),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer kafkaWriter.Close()

	kakfaMsg := kafka.Message{
		Key:   []byte(fmt.Sprintf("push-%d", time.Now().Unix())),
		Value: msg,
	}

	err = kafkaWriter.WriteMessages(context.Background(), kakfaMsg)

	return err
}
