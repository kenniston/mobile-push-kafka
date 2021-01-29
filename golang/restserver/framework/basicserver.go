package framework

import (
	"fmt"
	"github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
	"reflect"
)

// Defines Basic Server Methods
type BasicServer interface {
	GetApp() *iris.Application
	GetConfig() *viper.Viper
	Run() error
}

var server BasicServer

func GetServer() BasicServer {
	return server
}

var controllers []BaseController

// Store the controller for future initialization
func RegisterController(c BaseController) {
	controllers = append(controllers, c)
}

// Builds the Server from a series of steps.
func NewBasicServer(v *viper.Viper) BasicServer {
	validateConfig(v)

	// Configure Logs
	var logLevel logrus.Level
	logLevel, err := logrus.ParseLevel(v.GetString("log-level"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
	if logLevel == logrus.DebugLevel {
		resty.SetLogger(logrus.New().Out)
		resty.SetDebug(true)
	}

	// Configure Graylog
	var useGraylog = v.GetBool("graylog")
	var graylogIP = v.GetString("graylog-ip")
	var graylogPort = v.GetInt32("graylog-port")
	if useGraylog {
		hook := graylog.NewGraylogHook(fmt.Sprintf("%s:%d", graylogIP, graylogPort), nil)
		logrus.AddHook(hook)
		logrus.Infof("Log messages are now sent to Graylog (udp://%s:%d)", graylogIP, graylogPort)
	}

	// Create Basic Server Instance
	server := &basicServer{}
	server.config = v

	// Create Iris Server to serve
	server.app = iris.New()

	// Configure Logrus Middleware
	if logLevel == logrus.DebugLevel {
		requestLogger := logger.New(logger.Config{
			Status:             true,
			IP:                 true,
			Method:             true,
			Path:               true,
			Query:              true,
			MessageContextKeys: []string{"logger_message"},
			MessageHeaderKeys:  []string{"User-Agent"},
		})
		server.app.Use(requestLogger)
	}

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	server.app.Use(crs)
	server.app.AllowMethods(iris.MethodOptions)

	server.app.Get("/health", func(ctx iris.Context) {
		_, err := ctx.JSON(iris.Map{"message": "The Server is running..."})
		if err != nil {
			logrus.Fatal(err)
		}
	})

	// Setup Controllers
	for _, controller := range controllers {
		if controller.GetControllerName() == "" {
			logrus.Info("Controller with noname can't be initialized. Controller Type: %v", reflect.TypeOf(controller))
			continue
		}
		logrus.Infof("Initializing controller \"%s\"...", controller.GetControllerName())
		controller.Setup(server.app, server.config)
		logrus.Infof("\"%s\" Controller initialized!", controller.GetControllerName())
	}

	return server
}

//===============================================================================
// BasicServer is a struct which uses Microservices endpoints
// to integrate with other internal systems and external systems.
//
// This BasicServer has a set of Microservices endpoints which are
// called through HTTP protocol.
//
// The Microservices on this BasicServer mode communicate through API Gateway
// with other internal systems.
//
type basicServer struct {
	app    *iris.Application
	config *viper.Viper
}

// Returns Iris Application for this server
func (s *basicServer) GetApp() *iris.Application {
	return s.app
}

// Returns params loaded for this server
func (s *basicServer) GetConfig() *viper.Viper {
	return s.config
}

// Run Basic Service.
func (s *basicServer) Run() error {
	port := s.config.GetInt("server-port")
	logrus.Infof("Starting Server on port %d...\n", port)
	return s.app.Run(iris.Addr(fmt.Sprintf(":%d", port)))
}

// Validate the Basic Server parameters
func validateConfig(v *viper.Viper) {
	if _, err := logrus.ParseLevel(v.GetString("log-level")); err != nil {
		panic("Invalid Server Log Level.")
	}

	port := v.GetInt("server-port")
	if port <= 0 {
		panic("Invalid Server Port.")
	}
}
