package account

type Account struct {
	Id               int64    `db:"id" json:"id"`                              // 主键
	CreateTime       int64    `db:"create_time" json:"createTime"`             // 创建时间
	UpdateTime       int64    `db:"update_time" json:"updateTime"`             // 更新时间
	CompanyId        int64    `db:"company_id" json:"companyId"`               // 所属公司ID
	UserId           int64    `db:"user_id" json:"userId"`                     // 所有者id
	UserPositionId   int64    `db:"user_position_id" json:"userPositionId"`    // 身份ID
	UserPositionType int8    `db:"user_position_type" json:"userPositionType"` // 身份类型
	Type             int8    `db:"type" json:"type"`                           // 类型:1.余额；2.积分;
	Unit             int8    `db:"unit" json:"unit"`                           // 单元:1.分；2.元
	Balance          float64    `db:"balance" json:"balance"`                  // 余额
	Status           int8    `db:"status" json:"status"`                       // 状态 -1删除
}
