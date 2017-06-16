package personnel

type Structure struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 所有者ID
	ParentId   int64 `db:"parent_id" json:"parent_id"`     // parent_id
	Name       string `db:"name" json:"name"`              // 名称
	Type       int8 `db:"type" json:"type"`                // 类型 1.班级
	Status     int8 `db:"status" json:"status"`            // 状态 0.下线  1.上线
}
