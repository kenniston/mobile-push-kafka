package main

import (
	_ "github.com/kenniston/mobile-push-kafka/golang/producer/server/controller"
	"github.com/kenniston/mobile-push-kafka/golang/restserver"
)

func main() {
	microservice.Run()
}
