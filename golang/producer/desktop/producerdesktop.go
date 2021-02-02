package desktop

import (
	"context"
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"github.com/kataras/i18n"
	"github.com/kenniston/mobile-push-kafka/golang/producer/server/dto"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework/cmd"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"image/color"
	"time"
)

var desktopCmd = &cobra.Command{
	Use:     "desktop",
	Aliases: []string{"d"},
	Short:   "Starts Producer Window",
	Long: `This command start a new Producer Window with send message options`,
	RunE: func(cmd *cobra.Command, args []string) error {
		window()
		return nil
	},
}

func window() {
	a := app.New()
	w := a.NewWindow("Push Producer")
	w.Resize(fyne.NewSize(768, 326))
	w.SetFixedSize(true)

	header := canvas.NewText("Push", color.White)
	header.TextSize = 25
	header.TextStyle.Bold = true

	kafkaAddress := widget.NewEntry()
	kafkaAddress.Text = "localhost:9092"
	kafkaAddress.SetPlaceHolder("localhost:9092")
	kafkaAddress.Validator = validation.NewRegexp(`\w{1,}:\d{1,4}`, "Not a valid server address")

	kafkaTopic := widget.NewEntry()
	kafkaTopic.Text = "MobileSendPush"
	kafkaTopic.SetPlaceHolder("MobileSendPush")
	kafkaTopic.Validator = validation.NewRegexp(`\w{1,}`, "Not a valid Topic name")

	messageText := widget.NewMultiLineEntry()
	kafkaAddress.SetPlaceHolder("json / yaml / text / xml")
	messageText.Validator = validation.NewRegexp(`^(\s|\S)*(\S)+(\s|\S)*$`, "Message is required!")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Kafka Server", Widget: kafkaAddress, HintText: "Kafka Server Address and Port"},
			{Text: "Kafka Topic", Widget: kafkaTopic, HintText: "Kafka Topic Name"},
			{Text: "Message", Widget: messageText},
		},
		OnSubmit: func() {
			if err := sendKafkaMessage(messageText.Text, kafkaAddress.Text, kafkaTopic.Text); err != nil {
				title := i18n.Tr("en", "message-send-error")
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "Push Message",
					Content: title,
				})
				logrus.Error(title)
				return
			}
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Push Message",
				Content: "Message sent to the Kafka Server!",
			})
		},
		SubmitText: "Send Push",
	}

	w.SetContent(container.NewVBox(
		header,
		widget.NewSeparator(),
		form,
	))
	w.ShowAndRun()
}

func sendKafkaMessage(message string, server string, topic string) error {
	pushMsg := dto.PushMessage{
		Message: message,
	}

	kafkaWriter := &kafka.Writer{
		Addr:     kafka.TCP(server),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer kafkaWriter.Close()

	msg, err := json.Marshal(pushMsg)
	if err != nil {
		return err
	}
	logrus.Debug("Message: %s", string(msg))

	kakfaMsg := kafka.Message{
		Key:   []byte(fmt.Sprintf("push-%d", time.Now().Unix())),
		Value: msg,
	}
	e := kafkaWriter.WriteMessages(context.Background(), kakfaMsg)
	return e
}

func init() {
	cmd.GetRootCommand().AddCommand(desktopCmd)

	err := viper.GetViper().BindPFlags(desktopCmd.Flags())
	if err != nil {
		panic(err)
	}
}
