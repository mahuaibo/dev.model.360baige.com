package authority

type Gather struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 更新时间（毫秒）
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 企业ID
	Name       string    `db:"name" json:"name"`             // 权限名称
	Status     int8    `db:"status" json:"status"`           // 状态 -1注销、0正常
	Type       int8    `db:"type" json:"type"`               // 类型
}
