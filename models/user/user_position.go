package user

type UserPosition struct {
	Id          int64    `db:"id" json:"id"`                     // 主键
	CreateTime  int64    `db:"create_time" json:"createTime"`    // 创建时间
	UpdateTime  int64    `db:"update_time" json:"updateTime"`    // 更新时间
	Type        int8    `db:"type" json:"type"`                  // 用户类型
	CompanyId   int64    `db:"company_id" json:"companyId"`      // 企业ID
	UserId      int64    `db:"user_id" json:"userId"`            // 用户ID
	PersonId    int64    `db:"person_id" json:"personId"`        // 人事ID
	AccessToken string    `db:"access_token" json:"accessToken"` // 访问令牌
	ExpireIn    int64    `db:"expire_in" json:"expireIn"`        // 访问时效
	Status      int8    `db:"status" json:"status"`              // 状态 -1注销 0正常
}
