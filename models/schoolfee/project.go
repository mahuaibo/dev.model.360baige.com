package schoolfee

type Project struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间(毫秒)
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 更新时间(毫秒)
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 所属公司ID
	Name       string    `db:"name" json:"name"`             // 项目名称
	IsLimit    int8    `db:"is_limit" json:"isLimit"`        // 限制缴费 0：限制 1：不限制
	Desc       string    `db:"desc" json:"desc"`             // 描述
	Link       string    `db:"link" json:"link"`             // 描述链接
	Status     int8    `db:"status" json:"status"`           // 状态 -1注销 0正常
}
