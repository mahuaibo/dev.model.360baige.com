package card

type Card struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 企业ID
	PersonId   int64 `db:"person_id" json:"person_id"`     // 身份ID
	CardNo     string `db:"card_no" json:"card_no"`        // IC卡编号
	PhysicsNo  string `db:"physics_no" json:"physics_no"`  // 物理编号
	BatchNo    string `db:"batch_no" json:"batch_no"`      // 批次号
	Status     int8 `db:"status" json:"status"`            // 状态 -1注销、0未出库、1已出库、2未绑定、3已绑定、4挂失
}
