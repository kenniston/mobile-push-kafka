package framework

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

//===============================================================================
// BaseController is a interface to define a Controller basic methods
//
type BaseController interface {
	Setup(app *iris.Application, config *viper.Viper)
	GetControllerName() string
}
