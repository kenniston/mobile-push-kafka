package desktop

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kenniston/mobile-push-kafka/golang/restserver/framework/cmd"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sudo apt update && sudo apt install build-essential libxinerama-dev libxcursor-dev libx libx11-dev libgl1-mesa-dev \
// libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev xorg-dev

var desktopCmd = &cobra.Command{
	Use:     "desktop",
	Aliases: []string{"d"},
	Short:   "Starts Producer Window",
	Long: `This command start a new Producer Window with send message options`,
	RunE: func(cmd *cobra.Command, args []string) error {

		a := app.New()
		w := a.NewWindow("Push Producer")
		w.Resize(fyne.NewSize(768, 350))
		w.SetFixedSize(true)

		serverAddress := widget.NewEntry()

		w.SetContent(container.NewVBox(
			container.NewHBox(
				widget.NewLabel("Kafka Server (address:port): "),
				serverAddress,
			),
			widget.NewButton("Send", func() {
				msg := fmt.Sprintf("Sending push from Desktop. Server: %s", serverAddress.Text)
				logrus.Info(msg)
			}),
		))
		w.ShowAndRun()

		return nil
	},
}

func init() {
	cmd.GetRootCommand().AddCommand(desktopCmd)

	err := viper.GetViper().BindPFlags(desktopCmd.Flags())
	if err != nil {
		panic(err)
	}
}
