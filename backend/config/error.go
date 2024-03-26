package config

type ErrorSchema struct {
	Code    uint   `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
