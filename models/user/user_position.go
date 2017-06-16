package user

type UserPosition struct {
	Id          int64 `db:"id" json:"id"`                      // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`    // 创建时间
	UpdateTime  int64 `db:"update_time" json:"update_time"`    // 更新时间
	CompanyId   int64 `db:"company_id" json:"company_id"`      // 用户默认登录的企业
	UserId      int64 `db:"user_id" json:"user_id"`            // 用户ID
	PersonId    int64 `db:"person_id" json:"person_id"`        // 人事ID
	AccessToken string `db:"access_token" json:"access_token"` // 访问令牌
	ExpireIn    int64 `db:"expire_in" json:"expire_in"`        // 访问时效
}
