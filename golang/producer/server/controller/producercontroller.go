package controller

import (
	"github.com/kataras/i18n"
	"github.com/kataras/iris/v12"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/dto"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/service"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework/cmd"
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
	err := ctx.ReadBody(&msg)

	if err != nil {
		title := i18n.Tr("en", "message-parse-error")
		logrus.Error(title)
		ctx.StopWithProblem(iris.StatusBadRequest,
			iris.NewProblem().Title(title).Key("erros", framework.WrapValidationErrors(err)))
		return
	}

	if err := c.service.Send(msg); err != nil {
		title := i18n.Tr("en", "message-send-error")
		logrus.Error(title)
		ctx.StopWithProblem(iris.StatusServiceUnavailable,
			iris.NewProblem().Title(title).DetailErr(err))
		return
	}

	ctx.StatusCode(iris.StatusAccepted)
}

func init() {
	runCmd := cmd.GetRunCommand()

	runCmd.Flags().String("kafka-address", "localhost:9092", "Configure Kafka server address")
	runCmd.Flags().String("kafka-topic", "", "Configure Kafka Topic Name")

	err := viper.GetViper().BindPFlags(runCmd.Flags())
	if err != nil {
		panic(err)
	}

	framework.RegisterController(&producerController{})
}
