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
	Status       int    `db:"status" json:"status"`                 // 状态 -1注销 0正常
	WxOpenId     string `db:"wx_open_id" json:"wxOpenId"`           // 微信openId
	QqOpenId     string `db:"qq_open_id" json:"qqOpenId"`           // qq openId
}

type UserPosition struct {
	Id              int64    `db:"id" json:"id"`                             // 主键
	CreateTime      int64    `db:"create_time" json:"createTime"`            // 创建时间
	UpdateTime      int64    `db:"update_time" json:"updateTime"`            // 更新时间
	Type            int    `db:"type" json:"type"`                           // 用户类型
	CompanyId       int64    `db:"company_id" json:"companyId"`              // 企业ID
	UserId          int64    `db:"user_id" json:"userId"`                    // 用户ID
	PersonId        int64    `db:"person_id" json:"personId"`                // 人事ID
	AccessToken     string    `db:"access_token" json:"accessToken"`         // 访问令牌
	ExpireIn        int64    `db:"expire_in" json:"expireIn"`                // 访问时效
	TransitToken    string    `db:"transit_token" json:"transitToken"`       // 过渡令牌
	TransitExpireIn int64    `db:"transit_expire_in" json:"transitExpireIn"` // 过渡时效
	Status          int    `db:"status" json:"status"`                       // 状态 -1注销 0正常
}
