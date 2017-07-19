package user

type User struct {
	Id           int64 `db:"id" json:"id"`                        // 主键
	CreateTime   int64 `db:"create_time" json:"create_time"`      // 创建时间（毫秒）
	UpdateTime   int64 `db:"update_time" json:"update_time"`      // 修改时间（毫秒）
	Username     string `db:"username" json:"username"`           // 用户名
	Password     string `db:"password" json:"password"`           // 密码
	Email        string `db:"email" json:"email"`                 // 邮件
	Phone        string `db:"phone" json:"phone"`                 // 手机号
	Status       int8 `db:"status" json:"status"`                 // 状态 -1注销 0正常
	Code         string `db:"code" json:"code"`                   // 验证码
	CodeTime     int64 `db:"code_time" json:"code_time"`          // 验证码时效
	AccessTicket string `db:"access_ticket" json:"access_ticket"` // 访问令牌
	ExpireIn     int64 `db:"expire_in" json:"expire_in"`          // 访问时效
}
