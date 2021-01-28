package framework

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
	"reflect"
	"strings"
)

type BaseRepository interface {
	GetClient() *resty.Client
	GetConfig() *viper.Viper
	CheckRequestError(resp *resty.Response, err error, level logrus.Level) (bool, error)
	ChecksInitialized()
	Post(endpoint string, headers map[string]string, pathParams map[string]string, queryParams map[string]string,
		body interface{}, result interface{}) error
	Patch(endpoint string, headers map[string]string, pathParams map[string]string, queryParams map[string]string,
		body interface{}, result interface{}) error
	Get(endpoint string, headers map[string]string, pathParams map[string]string, queryParams map[string]string,
		result interface{}) error
	Delete(endpoint string, headers map[string]string, pathParams map[string]string, queryParams map[string]string,
		result interface{}) error
}

func NewBaseRepository(baseUrl string, apiKey string, name string, config *viper.Viper) BaseRepository {
	client := resty.New()
	client.SetHostURL(baseUrl)
	client.SetHeader("user_key", apiKey)

	if logLevel, err := logrus.ParseLevel(config.GetString("log-level")); err == nil {
		if logLevel == logrus.DebugLevel {
			client.SetLogger(logrus.New().Out)
			client.SetDebug(true)
		}
	}

	return &baseRepository{
		client:       client,
		serverConfig: config,
		name:         name,
		built:        true,
	}
}

//===============================================================================
// BaseRepository is a structure with common methods for use with
// Web Framework
//
type baseRepository struct {
	client       *resty.Client
	serverConfig *viper.Viper
	name         string
	built        bool
}

func (r *baseRepository) GetClient() *resty.Client {
	return r.client
}

func (r *baseRepository) GetConfig() *viper.Viper {
	return r.serverConfig
}

func (r *baseRepository) CheckRequestError(resp *resty.Response, err error, level logrus.Level) (bool, error) {
	desc := ""
	statusCode := 0
	body := make(map[string]interface{})

	if resp != nil {
		desc = resp.Status()
		statusCode = resp.StatusCode()
		if b := resp.Body(); b != nil {
			if parseError := json.Unmarshal(b, &body); parseError == nil {
				if msgs := body["message"]; msgs != nil {
					if tp := reflect.TypeOf(msgs); tp.Kind() == reflect.Slice || tp.Kind() == reflect.Array {
						arr := msgs.([]interface{})
						message := ""
						for _, m := range arr {
							message = fmt.Sprintf("%s%v - ", message, m)
						}
						desc = strings.TrimRight(message, " - ")
					}
				}
			}
		}
	}

	if err != nil || (resp != nil && resp.StatusCode() >= 400) {
		return false, NewIntegrationApiError(desc, statusCode, err, body, level)
	}

	return true, nil
}

func (r *baseRepository) ChecksInitialized() {
	if !r.built || r.name == "" {
		logrus.Fatal("(Repository) %s has not been initialized", r.name)
	}
}

func (r *baseRepository) Post(endpoint string, headers map[string]string, pathParams map[string]string,
	queryParams map[string]string, body interface{}, result interface{}) error {

	request := r.client.R()
	if headers != nil {
		request.SetHeaders(headers)
	}
	if pathParams != nil {
		request.SetPathParams(pathParams)
	}
	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}

	request.SetBody(body)

	resp, err := request.Post(endpoint)
	if ok, err := r.CheckRequestError(resp, err, logrus.ErrorLevel); !ok {
		return err
	}

	if parseError := json.Unmarshal(resp.Body(), &result); parseError != nil {
		return NewJsonParserError(fmt.Sprintf("Error parsing result for %v", result), parseError, logrus.ErrorLevel)
	}

	return nil
}

func (r *baseRepository) Patch(endpoint string, headers map[string]string, pathParams map[string]string,
	queryParams map[string]string, body interface{}, result interface{}) error {

	request := r.client.R()
	if headers != nil {
		request.SetHeaders(headers)
	}
	if pathParams != nil {
		request.SetPathParams(pathParams)
	}
	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}

	request.SetBody(body)

	resp, err := request.Patch(endpoint)
	if ok, err := r.CheckRequestError(resp, err, logrus.ErrorLevel); !ok {
		return err
	}

	if parseError := json.Unmarshal(resp.Body(), &result); parseError != nil {
		return NewJsonParserError(fmt.Sprintf("Error parsing result for %v", result), parseError, logrus.ErrorLevel)
	}

	return nil
}

func (r *baseRepository) Get(endpoint string, headers map[string]string, pathParams map[string]string,
	queryParams map[string]string, result interface{}) error {

	request := r.client.R()
	if headers != nil {
		request.SetHeaders(headers)
	}
	if pathParams != nil {
		request.SetPathParams(pathParams)
	}
	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}

	resp, err := request.Get(endpoint)
	if ok, err := r.CheckRequestError(resp, err, logrus.ErrorLevel); !ok {
		return err
	}

	if parseError := json.Unmarshal(resp.Body(), &result); parseError != nil {
		return NewJsonParserError(fmt.Sprintf("Error parsing result for %v", result), parseError, logrus.ErrorLevel)
	}

	return nil
}

func (r *baseRepository) Delete(endpoint string, headers map[string]string, pathParams map[string]string,
	queryParams map[string]string, result interface{}) error {

	request := r.client.R()
	if headers != nil {
		request.SetHeaders(headers)
	}
	if pathParams != nil {
		request.SetPathParams(pathParams)
	}
	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}

	resp, err := request.Delete(endpoint)
	if ok, err := r.CheckRequestError(resp, err, logrus.ErrorLevel); !ok {
		return err
	}

	if parseError := json.Unmarshal(resp.Body(), &result); parseError != nil {
		return NewJsonParserError(fmt.Sprintf("Error parsing result for %v", result), parseError, logrus.ErrorLevel)
	}

	return nil
}
