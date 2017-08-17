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
	Id           int64 `json:"id"`
	Username     string `json:"username"`      // 用户名
	AccessTicket string `json:"access_ticket"` // 访问令牌
	ExpireIn     int64 `json:"expire_in"`      // 访问时效
}
type UserDetail struct {
	Id       int64`json:"id"`         // 主键
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮件
	Phone    string `json:"phone"`    // 手机号
}
type UserDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserDetail   `json:"data,omitempty"`
}
