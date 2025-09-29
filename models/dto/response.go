package dto

type EmptyResponse struct {
	Status  int         `json:"status" example:"200"`
	Data    EmptyObject `json:"data"`
	Message string      `json:"message" example:"OK"`
}

const (
	MessageOK = "OK"
)

type EmptyObject struct{}
