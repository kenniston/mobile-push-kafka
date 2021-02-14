package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kataras/i18n"
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
	kafkaAddress string
	kafkaTopic string
}

// Create and initialize the repository
func NewSecurityRepository(config *viper.Viper) ProducerRepository {
	address := config.GetString("kafka-address")
	topic := config.GetString("kafka-topic")

	return &producerRepository{
		BaseRepository: framework.NewBaseRepository("", "", "Producer Repository", config),
		kafkaAddress: address,
		kafkaTopic: topic,
	}
}

// Return a Security list object
func (r *producerRepository) Send(message dto.PushMessage) error {
	r.ChecksInitialized()

	kafkaWriter := &kafka.Writer{
		Addr:     kafka.TCP(r.kafkaAddress),
		Topic:    r.kafkaTopic,
		Balancer: &kafka.LeastBytes{},
		RequiredAcks: 1,
		Async: true,
		Completion: func(messages []kafka.Message, err error) {
			if err != nil {
				logrus.Error(i18n.Tr("en", "message-send-error"), err)
				return
			}
		},
	}
	defer func() {
		err := kafkaWriter.Close()
		if err != nil {
			logrus.Error(i18n.Tr("en", "message-kafka-close-error"), err)
		}
	}()

	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}
	logrus.Infof("Message: %s", string(msg))

	kafkaMsg := kafka.Message{
		Key:   []byte(fmt.Sprintf("push-%d", time.Now().Unix())),
		Value: msg,
	}
	err = kafkaWriter.WriteMessages(context.Background(), kafkaMsg)

	return err
}

