package response

type Response struct {
	Code    string        `json:"code"`
	Messgae string        `json:"messgae"`
	Data    interface{}   `json:"data,omitempty"`
}

// Response 状态码
const (
	ResponseSystemErr = "500"
	ResponseLogicErr  = "400"
	ResponseNormal    = "200"
)
