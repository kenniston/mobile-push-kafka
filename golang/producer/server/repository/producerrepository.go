package repository

import (
	"encoding/json"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/dto"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	serverConfig *viper.Viper
}

// Create and initialize the repository
func NewSecurityRepository(config *viper.Viper) ProducerRepository {
	return &producerRepository{serverConfig: config}
}

// Return a Security list object
func (r *producerRepository) Send(message dto.PushMessage) error {
	r.ChecksInitialized()

	str, err := json.Marshal(message)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Debug("Message: %s", string(str))

	//TODO: Send message to the Kafka's topic

	return nil
}
