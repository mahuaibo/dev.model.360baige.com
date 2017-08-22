package card

type Card struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 更新时间（毫秒）
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 企业ID
	PersonId   int64    `db:"person_id" json:"personId"`     // 身份ID
	CardNo     string    `db:"card_no" json:"cardNo"`        // IC卡编号
	PhysicsNo  string    `db:"physics_no" json:"physicsNo"`  // 物理编号
	BatchNo    string    `db:"batch_no" json:"batchNo"`      // 批次号
	Status     int8    `db:"status" json:"status"`           // 状态 -1注销、0未出库、1已出库、2未绑定、3已绑定、4挂失
}
