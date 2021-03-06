package order

type Order struct {
	Id               int64    `db:"id" json:"id"`                              // 主键
	CreateTime       int64    `db:"create_time" json:"createTime"`             // 创建时间
	UpdateTime       int64    `db:"update_time" json:"updateTime"`             // 更新时间
	CompanyId        int64    `db:"company_id" json:"companyId"`               // 企业ID
	UserId           int64    `db:"user_id" json:"userId"`                     // 用户ID
	UserPositionType int     `db:"user_position_type" json:"userPositionType"` // 身份类型
	UserPositionId   int64    `db:"user_position_id" json:"userPositionId"`    // 身份ID
	Code             string   `db:"code" json:"code"`                          // 订单编号
	SubCode          string   `db:"sub_code" json:"subCode"`                   // 子订单编号
	Price            int64    `db:"price" json:"price"`                        // 单价
	Num              int64    `db:"num" json:"num"`                            // 数量
	TotalPrice       int64    `db:"total_price" json:"totalPrice"`             // 总价
	ProductType      int     `db:"product_type" json:"productType"`            // 产品类型
	Image            string   `db:"image" json:"image"`                        // 订单图片
	ProductId        int64    `db:"product_id" json:"productId"`               // 产品ID
	PayType          int     `db:"pay_type" json:"payType"`                    // 支付方式
	Brief            string   `db:"brief" json:"brief"`                        // 详情
	Status           int     `db:"status" json:"status"`                       // 订单状态
	TradeType        string   `db:"trade_type" json:"tradeType"`               // 交易类型
	PrepayId         string   `db:"prepay_id" json:"prepayId"`                 // 预支付交易会话标识
	CodeUrl          string   `db:"code_url" json:"codeUrl"`                   // 二维码链接
	Openid           string   `db:"openid" json:"openid"`                      // 用户标识
}
