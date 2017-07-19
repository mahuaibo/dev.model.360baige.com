package personnel

type Person struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 所有者ID
	Name       string `db:"name" json:"name"`              // 名称
	Type       int8 `db:"type" json:"type"`                // 类型 1.教师   2.家长   3.学生
	Status     int8 `db:"status" json:"status"`            // 状态 1.正常   2.异常
}
