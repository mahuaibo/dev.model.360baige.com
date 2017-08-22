package personnel

type Structure struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 更新时间（毫秒）
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 所有者ID
	ParentId   int64    `db:"parent_id" json:"parentId"`     // parent_id 上级部门ID
	Name       string    `db:"name" json:"name"`             // 名称
	Type       int8    `db:"type" json:"type"`               // 类型 1.班级 2.部门
	Status     int8    `db:"status" json:"status"`           // 状态 0.下线  1.上线
}
