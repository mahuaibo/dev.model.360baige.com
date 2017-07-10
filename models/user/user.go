package user

type User struct {
	Id           int64 `db:"id" json:"id,omitempty"`                        // 主键
	CreateTime   int64 `db:"create_time" json:"create_time,omitempty"`      // 创建时间（毫秒）
	UpdateTime   int64 `db:"update_time" json:"update_time,omitempty"`      // 修改时间（毫秒）
	Username     string `db:"username" json:"username,omitempty"`           // 用户名
	Password     string `db:"password" json:"password,omitempty"`           // 密码
	Email        string `db:"email" json:"email,omitempty"`                 // 邮件
	Phone        string `db:"phone" json:"phone,omitempty"`                 // 手机号
	Status       int8 `db:"status" json:"status,omitempty"`                 // 状态 -1注销 0正常
	Code         string `db:"code" json:"code,omitempty"`                   // 验证码
	CodeTime     int64 `db:"code_time" json:"code_time,omitempty"`          // 验证码时效
	AccessTicket string `db:"access_ticket" json:"access_ticket,omitempty"` // 访问令牌
	ExpireIn     int64 `db:"expire_in" json:"expire_in,omitempty"`          // 访问时效
}
