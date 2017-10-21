package authority

type UserLoginResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserLogin   `json:"data,omitempty"`
}

type UserBindResponse struct {
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    UserBind   `json:"data,omitempty"`
}

type UserBind struct {
	OpenType int `json:"openType"`  // 第三方类型
	OpenId   string `json:"openId"` // 第三方openId
}

type ModifyPasswordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

type UserLogin struct {
	Username     string `json:"username"`     // 用户名
	Head         string `json:"head"`         // 头像
	AccessTicket string `json:"accessTicket"` // 访问令牌
	ExpireIn     int64 `json:"expireIn"`      // 访问时效
}

type UserDetail struct {
	Email  string `json:"email"`  // 邮件
	Phone  string `json:"phone"`  // 手机号
	Qq     string `json:"qq"`     // QQopenId
	WeChat string `json:"weChat"` // 微信openId
}

type UserDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserDetail   `json:"data,omitempty"`
}

type UserRegisterResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

type ExistKeyResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

type SendMessageCodeResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

type UploadHeadResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    string   `json:"head"`
}