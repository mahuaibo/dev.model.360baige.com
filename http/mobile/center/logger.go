package center

type LoggerAddResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    LoggerAdd   `json:"data,omitempty"`
}
type LoggerAdd struct {
	Id  int64 `db:"id" json:"id"`   //ap id
}