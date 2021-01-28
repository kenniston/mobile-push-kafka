package framework

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

func printError(err error, level logrus.Level) {
	switch level {
	case logrus.DebugLevel:
		logrus.Debug(err)
	case logrus.InfoLevel:
		logrus.Info(err)
	case logrus.ErrorLevel:
		logrus.Error(err)
	case logrus.FatalLevel:
		logrus.Fatal(err)
	case logrus.PanicLevel:
		logrus.Panic(err)
	case logrus.TraceLevel:
		logrus.Trace(err)
	case logrus.WarnLevel:
		logrus.Warn(err)
	}
}

//===============================================================================
// IntegrationApiError is a structure API Integration Errors
//
type IntegrationApiError struct {
	ErrorCode   int
	Description string
	StatusCode  int
	Body        map[string]interface{}
	Err         error
}

func (e *IntegrationApiError) Error() string {
	var body string
	if jsonBody, parseError := json.Marshal(e.Body); parseError == nil {
		body = string(jsonBody)
	} else {
		body = "null"
	}

	return fmt.Sprintf("{ ErrorCode: %d, Description: \"%s\", StatusCode: %d, Body: %v, Error: %v}",
		e.ErrorCode, e.Description, e.StatusCode, body, e.Err)
}

func NewIntegrationApiError(desc string, status int, err error, body map[string]interface{}, logLevel logrus.Level) *IntegrationApiError {
	ex := &IntegrationApiError{-9999, desc, status, body, err}
	printError(ex, logLevel)
	return ex
}

//===============================================================================
// JsonParserError is a structure json parse errors (Marshal and Unmarshal)
//
type JsonParserError struct {
	ErrorCode   int
	Description string
	Detail      string
	Err         error
}

func (e *JsonParserError) Error() string {
	return fmt.Sprintf("{ ErrorCode: %d, Description: %s, Error: %v}",
		e.ErrorCode, e.Description, e.Err)
}

func NewJsonParserError(detail string, err error, logLevel logrus.Level) *JsonParserError {
	ex := &JsonParserError{-9998, "Error Parsing JSON.", detail, err}
	printError(ex, logLevel)
	return ex
}

//===============================================================================
// RequiredError is a structure required params
//
type RequiredError struct {
	ErrorCode   int
	Description string
	Detail      string
	Err         error
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("{ ErrorCode: %d, Description: %s, Error: %v}",
		e.ErrorCode, e.Description, e.Err)
}

func NewRequiredError(field string, detail string, err error, logLevel logrus.Level) *RequiredError {
	ex := &RequiredError{-9997, fmt.Sprintf("%s is required!", field), detail, err}
	printError(ex, logLevel)
	return ex
}
