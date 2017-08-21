package center

type UserLoginResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserLogin   `json:"data,omitempty"`
}
type ModifyPasswordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

type UserLogin struct {
	Username     string `db:"username" json:"username"`           // 用户名
	AccessTicket string `db:"access_ticket" json:"access_ticket"` // 访问令牌
	ExpireIn     int64 `db:"expire_in" json:"expire_in"`          // 访问时效
}

type UserRegisterResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}
