package window

type LoggerResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    UserLogin   `json:"data,omitempty"`
}
