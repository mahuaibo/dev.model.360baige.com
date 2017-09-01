package account

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
