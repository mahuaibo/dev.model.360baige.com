package user

type User struct {
	Id           int64    `db:"id" json:"id"`                       // 主键
	CreateTime   int64    `db:"create_time" json:"createTime"`      // 创建时间（毫秒）
	UpdateTime   int64    `db:"update_time" json:"updateTime"`      // 修改时间（毫秒）
	Head         string    `db:"head" json:"head"`                  // 头像
	Username     string    `db:"username" json:"username"`          // 用户名
	Password     string    `db:"password" json:"password"`          // 密码
	Email        string    `db:"email" json:"email"`                // 邮件
	Phone        string    `db:"phone" json:"phone"`                // 手机号
	AccessTicket string    `db:"access_ticket" json:"accessTicket"` // 访问令牌
	ExpireIn     int64    `db:"expire_in" json:"expireIn"`          // 访问时效
	Status       int8    `db:"status" json:"status"`                // 状态 -1注销 0正常
}
