package personnel

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
