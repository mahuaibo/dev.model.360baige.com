package models

type Person struct {
	ID        int64 `db:"id" json:"id"`                 // 主键
	CreateTime     int64 `db:"ctime" json:"ctime"`           // 创建时间（毫秒）
	UpdateTime     int64 `db:"utime" json:"utime"`           // 更新时间（毫秒）
	CompanyID int64 `db:"company_id" json:"company_id"` // 所有者ID
	Name      string `db:"name" json:"name"`            // 名称
	Type      int8 `db:"type" json:"type"`              // 类型 1.教师   2.家长   3.学生
	Status    int8 `db:"status" json:"status"`          // 状态 1.正常   2.异常
}

type PersonRelation struct {
	ID            int64 `db:"id" json:"id"`                         // 主键
	CreateTime          int64 `db:"cime" json:"cime"`                     // 创建时间（毫秒）
	UpdateTime          int64 `db:"uime" json:"uime"`                     // 更新时间（毫秒）
	CompanyID     int64 `db:"company_id" json:"company_id"`         // 所有者ID
	AssociationID int64 `db:"association_id" json:"association_id"` // parent_id（关联人ID）
	AssociatedID  int64 `db:"associated_id" json:"associated_id"`   // parent_id（被关联人ID）
	Type          int8 `db:"type" json:"type"`                      // 类型
	Status        int8 `db:"status" json:"status"`                  // 状态  1.关联   2.取消关联
}

type PersonStructure struct {
	ID          int64 `db:"id" json:"id"`                     // 主键
	CreateTime       int64 `db:"ctime" json:"ctime"`               // 创建时间（毫秒）
	UpdateTime       int64 `db:"utime" json:"utime"`               // 更新时间（毫秒）
	CompanyID   int64 `db:"company_id" json:"company_id"`     // 所有者ID
	PersonID    int64 `db:"person_id" json:"person_id"`       // parent_id 人员ID
	StructureID int64 `db:"structure_id" json:"structure_id"` // structure_id 结构ID
	Type        int8 `db:"type" json:"type"`                  // 类型
	Status      int8 `db:"status" json:"status"`              // 状态  1.正常
}

type Structure struct {
	ID        int64 `db:"id" json:"id"`                 // 主键
	CreateTime     int64 `db:"ctime" json:"ctime"`           // 创建时间（毫秒）
	UpdateTime     int64 `db:"utime" json:"utime"`           // 更新时间（毫秒）
	CompanyID int64 `db:"company_id" json:"company_id"` // 所有者ID
	ParentID  int64 `db:"parent_id" json:"parent_id"`   // parent_id
	Name      string `db:"name" json:"name"`            // 名称
	Type      int8 `db:"type" json:"type"`              // 类型 1.班级
	Status    int8 `db:"status" json:"status"`          // 状态 0.下线  1.上线
}
