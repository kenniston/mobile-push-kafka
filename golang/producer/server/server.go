package server

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"net/http"
)

type Service struct {
	Endpoint string
	Port string
}

//===============================================================================
// Server is a interface for a Producer Server
//
//
type Server interface {
	// ConfigureServer initialize server port, Kafka Topic and Partition
	ConfigureServer(v *viper.Viper)

	// Run 'run' the Producer Server
	Run() error
}

//===============================================================================
// ProducerServer is a server which uses REST endpoints
// to send messages to the Kafka Server.
//
type ProducerServer struct {
	app *iris.Application
	Service
}

func (s *ProducerServer) ConfigureServer(v *viper.Viper) {
	s.app = iris.New()

	//pushAPI := s.app.Party("/v1/push")
	//{
	//	pushAPI.Get("/", server)
	//}

	s.Port = v.GetString("ms-port")
}


func (s *ProducerServer) Run() error {
	fmt.Printf("Starting GraphQL Server with Microservices on port %s...\n", s.Port)
	return s.app.Run(iris.Addr(fmt.Sprintf(":%s", s.Port)))
}

//===============================================================================
// disableCors disable CORS for Producer Server
//
//
func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}