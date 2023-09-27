package model

type PublishReq struct {
	Topic   string `json:"topic" validate:"required"`
	Message string `json:"message" validate:"required"`
}
