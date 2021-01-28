package framework

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BaseService interface {
	ChecksInitialized()
	GetConfig() *viper.Viper
}

//===============================================================================
// BaseRepository is a structure with common methods for use with
// Web Framework
//
type baseService struct {
	serverConfig *viper.Viper
	name         string
	built        bool
}

func (r *baseService) GetConfig() *viper.Viper {
	return r.serverConfig
}

func (r *baseService) ChecksInitialized() {
	if !r.built || r.name == "" {
		logrus.Fatal("(Service) %s has not been initialized", r.name)
	}
}

func NewBaseService(name string, config *viper.Viper) BaseService {
	return &baseService{name: name, serverConfig: config, built: true}
}
