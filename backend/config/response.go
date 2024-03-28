package config

type Response struct {
	Code    uint        `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Row     int64       `json:"row,omitempty"`
}
