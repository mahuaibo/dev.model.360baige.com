package personnel

type PersonStructure struct {
	Id          int64 `db:"id" json:"id"`                     // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`   // 创建时间（毫秒）
	UpdateTime  int64 `db:"update_time" json:"update_time"`   // 更新时间（毫秒）
	CompanyId   int64 `db:"company_id" json:"company_id"`     // 所有者ID
	PersonId    int64 `db:"person_id" json:"person_id"`       // parent_id 人员ID
	StructureId int64 `db:"structure_id" json:"structure_id"` // structure_id 结构ID
	Type        int8 `db:"type" json:"type"`                  // 类型
	Status      int8 `db:"status" json:"status"`              // 状态  1.正常
}
