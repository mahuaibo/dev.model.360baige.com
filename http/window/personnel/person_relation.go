package personnel

type ListOfPersonRelationRelationResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfPersonRelation   `json:"data,omitempty"`
}

type ListOfPersonRelation struct {
	List []PersonRelation
}

type PersonRelation struct {
	Id            int64 `db:"id" json:"id"`                         // 主键
	CreateTime    int64 `db:"create_time" json:"create_time"`       // 创建时间（毫秒）
	UpdateTime    int64 `db:"update_time" json:"update_time"`       // 更新时间（毫秒）
	CompanyId     int64 `db:"company_id" json:"company_id"`         // 所有者ID
	AssociationId int64 `db:"association_id" json:"association_id"` // parent_id（关联人ID）
	AssociatedId  int64 `db:"associated_id" json:"associated_id"`   // parent_id（被关联人ID）
	Type          int8 `db:"type" json:"type"`                      // 类型
	Status        int8 `db:"status" json:"status"`                  // 状态  1.关联   2.取消关联
}

type AddPersonRelationResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddPersonRelation   `json:"data,omitempty"`
}

type AddPersonRelation struct {
	Id int64 `json:"id"`
}


type ModifyPersonRelationResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ModifyPersonRelation   `json:"data,omitempty"`
}

type DetailPersonRelationResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    PersonRelation   `json:"data,omitempty"`
}

type ModifyPersonRelation struct {
	Count int64 `json:"count"`
}