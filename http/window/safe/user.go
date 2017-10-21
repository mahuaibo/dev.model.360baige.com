package safe

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

type UserDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserDetail   `json:"data,omitempty"`
}

type UserDetail struct {
	Email  string `json:"email"`  // 邮件
	Phone  string `json:"phone"`  // 手机号
	Qq     string `json:"qq"`     // QQopenId
	WeChat string `json:"weChat"` // 微信openId
}

type ExistKeyResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ExistKeyUserData   `json:"data,omitempty"`
}

type ExistKeyUserData struct {
	UserId int64 `json:"userId"` // userId
}

type SendMessageCodeResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

type BindOrUnbindResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}