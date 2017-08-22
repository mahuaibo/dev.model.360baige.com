package website

type Menu struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 修改时间（毫秒）
	Icon       string    `db:"icon" json:"icon"`             // 图标
	CompanyId  string    `db:"company_id" json:"companyId"`  // 企业ID
	Name       string    `db:"name" json:"name"`             // 名称
	Sn         string    `db:"sn" json:"sn"`                 // 排序
	Status     int8    `db:"status" json:"status"`           // 状态 -1注销 0下线 1上线
}
