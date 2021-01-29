module github.com/kenniston/mobile-push-kafka/golang/producer

go 1.15

require (
	github.com/kataras/iris/v12 v12.2.0-alpha2
	github.com/kenniston/mobile-push-kafka/golang/restserver v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
)

replace github.com/kenniston/mobile-push-kafka/golang/restserver => ../restserver
