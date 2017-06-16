package attendance

type AttendanceShift struct {
	Id         int64 `db:"id" json:"id"`                   // 主键(班次ID)
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 所有者ID
	Type       int8 `db:"type" json:"type"`                // 班次类型
	Name       string `db:"name" json:"name"`              // 班次名称
	ShortName  string `db:"short_name" json:"short_name"`  // 班次简称
	Color      string `db:"color" json:"color"`            // 颜色
	Status     int8 `db:"status" json:"status"`            // 状态 -1删除
}
