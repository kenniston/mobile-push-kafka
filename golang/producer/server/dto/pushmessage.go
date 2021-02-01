package dto

type PushMessage struct {
	Message string `json:"message" yaml:"message" validate:"required"`
}
