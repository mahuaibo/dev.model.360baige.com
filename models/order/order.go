package order

type Order struct {
	Id               int64    `db:"id" json:"id"`                              // 主键
	CreateTime       int64    `db:"create_time" json:"createTime"`             // 创建时间
	UpdateTime       int64    `db:"update_time" json:"updateTime"`             // 更新时间
	CompanyId        int64    `db:"company_id" json:"companyId"`               // 企业ID
	UserId           int64    `db:"user_id" json:"userId"`                     // 用户ID
	UserPositionType int8    `db:"user_position_type" json:"userPositionType"` // 身份类型
	UserPositionId   int64    `db:"user_position_id" json:"userPositionId"`    // 身份ID
	Code             string    `db:"code" json:"code"`                         // 订单编号
	Price            float64    `db:"price" json:"price"`                      // 单价
	Type             int8    `db:"type" json:"type"`                           // 订单类型
	PayType          int8    `db:"pay_type" json:"payType"`                    // 支付方式 1在线支付   2线下支付
	Brief            string    `db:"brief" json:"brief"`                       // 详情
	Status           int8    `db:"status" json:"status"`                       // 订单状态：0:撤回 1：待审核  2：已通过 3：未通过 4：发货中 5：完成
}
