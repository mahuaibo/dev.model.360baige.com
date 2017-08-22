package schoolfee

type Record struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间(毫秒)
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 更新时间(毫秒)
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 所属公司ID
	ProjectId  int64    `db:"project_id" json:"projectId"`   // 项目ID
	Name       string    `db:"name" json:"name"`             // 姓名
	ClassName  string    `db:"class_name" json:"className"`  // 班级名称
	IdCard     string    `db:"id_card" json:"idCard"`        // 身份证号
	Num        string    `db:"num" json:"num"`               // 编号
	Phone      string    `db:"phone" json:"phone"`           // 联系电话
	Price      float64    `db:"price" json:"price"`          // 应缴费用
	IsFee      int8    `db:"is_fee" json:"isFee"`            // 是否缴费 0：未缴费 1：缴费
	FeeTime    int64    `db:"fee_time" json:"feeTime"`       // 缴费时间（毫秒）
	Desc       string    `db:"desc" json:"desc"`             // 备注
	Status     int8    `db:"status" json:"status"`           // 状态 -1删除
}
