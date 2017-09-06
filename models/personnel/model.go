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

type PersonRelation struct {
	Id            int64    `db:"id" json:"id"`                        // 主键
	CreateTime    int64    `db:"create_time" json:"createTime"`       // 创建时间（毫秒）
	UpdateTime    int64    `db:"update_time" json:"updateTime"`       // 更新时间（毫秒）
	CompanyId     int64    `db:"company_id" json:"companyId"`         // 所有者ID
	AssociationId int64    `db:"association_id" json:"associationId"` // parent_id（关联人ID）
	AssociatedId  int64    `db:"associated_id" json:"associatedId"`   // parent_id（被关联人ID）
	Type          int8    `db:"type" json:"type"`                     // 类型
	Status        int8    `db:"status" json:"status"`                 // 状态  1.关联   2.取消关联
}

type Person struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 更新时间（毫秒）
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 所有者ID
	Code       string    `db:"code" json:"code"`             // 编号
	Name       string    `db:"name" json:"name"`             // 名称
	Sex        string    `db:"sex" json:"sex"`               // 性别
	Birthday   int64    `db:"birthday" json:"birthday"`      // 生日
	Type       int8    `db:"type" json:"type"`               // 类型 1.教师   2.学生   3.开发者
	Phone      string    `db:"phone" json:"phone"`           // 联系电话
	Contact    string    `db:"contact" json:"contact"`       // 联系人
	Status     int8    `db:"status" json:"status"`           // 状态 -1.删除 1.正常   2.异常
}
