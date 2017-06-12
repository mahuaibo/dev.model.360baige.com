package models

type Machine struct {
	ID         int64 `db:"id" json:"id"`                 // 主键
	CreateTime int64 `db:"create_time" json:"ctime"`           // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"utime"`           // 更新时间（毫秒）
	CompanyID  int64 `db:"company_id" json:"company_id"` // 所有者ID
}
