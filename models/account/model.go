package account

type Account struct {
	Id               int64    `db:"id" json:"id"`                              // 主键
	CreateTime       int64    `db:"create_time" json:"createTime"`             // 创建时间
	UpdateTime       int64    `db:"update_time" json:"updateTime"`             // 更新时间
	CompanyId        int64    `db:"company_id" json:"companyId"`               // 所属公司ID
	UserId           int64    `db:"user_id" json:"userId"`                     // 所有者id
	UserPositionId   int64    `db:"user_position_id" json:"userPositionId"`    // 身份ID
	UserPositionType int8     `db:"user_position_type" json:"userPositionType"` // 身份类型
	Type             int8     `db:"type" json:"type"`                           // 类型:1.余额；2.积分;
	Balance          int64    `db:"balance" json:"balance"`                    // 余额
	Status           int8     `db:"status" json:"status"`                       // 状态 -1删除
}

type AccountItem struct {
	Id            int64    `db:"id" json:"id"`                        // 主键自动增长id
	CreateTime    int64    `db:"create_time" json:"createTime"`       // 创建时间（毫秒）
	UpdateTime    int64    `db:"update_time" json:"updateTime"`       // 修改时间（毫秒）
	TransactionId int64    `db:"transaction_id" json:"transactionId"` // 交易id
	AccountId     int64    `db:"account_id" json:"accountId"`         // 账户id
	Amount        int64    `db:"amount" json:"amount"`                // 交易金额
	Balance       int64    `db:"balance" json:"balance"`              // 平账
	Remark        string    `db:"remark" json:"remark"`               // 备注
	Status        int8    `db:"status" json:"status"`                 // 状态
}

type Transaction struct {
	Id            int64    `db:"id" json:"id"`                         // 主键
	CreateTime    int64    `db:"create_time" json:"createTime"`        // 创建时间
	UpdateTime    int64    `db:"update_time" json:"updateTime"`        // 更新时间
	FromAccountId int64    `db:"from_account_id" json:"fromAccountId"` // 交易人id
	ToAccountId   int64    `db:"to_account_id" json:"toAccountId"`     // 收款人id
	Amount        int64    `db:"amount" json:"amount"`                 // 交易金额
	OrderCode     string    `db:"order_code" json:"orderCode"`         // 订单编码
	Remark        string    `db:"remark" json:"remark"`                // 备注
	Status        int8      `db:"status" json:"status"`                  // 状态
}
