package mobile

// 注册
type UserRegisterResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    UserRegisterData   `json:"data,omitempty"`
}

type UserRegisterData struct {
	Body UserRegister
}
type UserRegister struct {
	Username string `json:"username"` // 用户名
}

// 登录
type UserLoginResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    UserLoginData   `json:"data,omitempty"`
}

type UserLoginData struct {
	Body UserLogin
}

type UserLogin struct {
	Username     string `json:"username"`      // 用户名
	AccessTicket string `json:"access_ticket"` // 访问令牌
	ExpireIn     int64 `json:"expire_in"`      // 访问时效
}

//
//type ModifyPasswordResponse struct {
//	Code    string `json:"code"`
//	Messgae string        `json:"messgae"`
//}
//
//type UserLogin struct {
//	Username     string `db:"username" json:"username"`           // 用户名
//	AccessTicket string `db:"access_ticket" json:"access_ticket"` // 访问令牌
//	ExpireIn     int64 `db:"expire_in" json:"expire_in"`          // 访问时效
//}
//type UserDetail struct {
//	Username string `db:"username" json:"username"` // 用户名
//	Email    string `db:"email" json:"email"`       // 邮件
//	Phone    string `db:"phone" json:"phone"`       // 手机号
//}
//type UserDetailResponse struct {
//	Code    string `json:"code"`
//	Messgae string        `json:"messgae"`
//	Data    UserDetail   `json:"data,omitempty"`
//}
