package service

import (
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/dto"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/repository"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework"
	"github.com/spf13/viper"
)

type ProducerService interface {
	Send(message dto.PushMessage) error
}

//===============================================================================
// Producer Service is a structure to define business methods for
// Producer API's.
// Business Logic MUST be placed in this structure.
//
type producerService struct {
	framework.BaseService
	repository repository.ProducerRepository
}

// Create and initialize the service
func NewProducerService(config *viper.Viper) ProducerService {
	return &producerService{
		BaseService: framework.NewBaseService("Producer Service", config),
		repository: repository.NewSecurityRepository(config),
	}
}

// Send a Push message data to the topic in Kafka server.
func (s *producerService) Send(message dto.PushMessage) error {
	s.ChecksInitialized()
	return s.repository.Send(message)
}
