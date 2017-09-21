package center

type UserLoginResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserLogin   `json:"data,omitempty"`
}

type UserBindResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserBind   `json:"data,omitempty"`
}

type UserBind struct {
	OpenId string `json:"openId"` // 微信Id
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
	Id       int64`json:"id"`         // 主键
	Username string `json:"username"` // 用户名
	Email    string `json:"email"`    // 邮件
	Phone    string `json:"phone"`    // 手机号
	Head     string `json:"head"`     // 头像
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
