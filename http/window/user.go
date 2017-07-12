package window

type UserLoginResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    UserLogin   `json:"data,omitempty"`
}

type UserLogin struct {
	Username     string `db:"username" json:"username"`           // 用户名
	AccessTicket string `db:"access_ticket" json:"access_ticket"` // 访问令牌
	ExpireIn     int64 `db:"expire_in" json:"expire_in"`          // 访问时效
}
