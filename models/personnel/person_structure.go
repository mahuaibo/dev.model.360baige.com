package personnel

type PersonStructure struct {
	Id          int64    `db:"id" json:"id"`                    // 主键
	CreateTime  int64    `db:"create_time" json:"createTime"`   // 创建时间（毫秒）
	UpdateTime  int64    `db:"update_time" json:"updateTime"`   // 更新时间（毫秒）
	CompanyId   int64    `db:"company_id" json:"companyId"`     // 所有者ID
	PersonId    int64    `db:"person_id" json:"personId"`       // parent_id 人员ID
	StructureId int64    `db:"structure_id" json:"structureId"` // structure_id 结构ID
	Type        int8    `db:"type" json:"type"`                 // 类型
	Status      int8    `db:"status" json:"status"`             // 状态  1.正常
}
