package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/dto"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/service"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//===============================================================================
// ProducerController is a structure to define methods for Producer API
//
type producerController struct {
	service service.ProducerService
}

func (c *producerController) GetControllerName() string {
	return "Producer Controller"
}

func (c *producerController) Setup(app *iris.Application, config *viper.Viper) {
	c.service = service.NewProducerService(config)

	app.Post("/v1/push/send", c.send)
}

// Send a Push message data to the topic in Kafka server.
func (c *producerController) send(ctx iris.Context) {
	var msg dto.PushMessage
	err := ctx.ReadJSON(msg)

	if err != nil {
		logrus.Error(err)
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.StopExecution()
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
}

func init() {
	framework.RegisterController(&producerController{})
}
