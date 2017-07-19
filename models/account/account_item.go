package account

type AccountItem struct {
	Id            int64 `db:"id" json:"id"`                         // 主键自动增长id
	CreateTime    int64 `db:"create_time" json:"create_time"`       // 创建时间
	UpdateTime    int64 `db:"update_time" json:"update_time"`       // 更新时间
	TransactionId int64 `db:"transaction_id" json:"transaction_id"` // 交易id
	AccountId     int64 `db:"account_id" json:"account_id"`         // 账户id
	Amount        float64 `db:"amount" json:"amount"`               // 交易金额
	Balance       float64 `db:"balance" json:"balance"`             // 平账
	Remark        string `db:"remark" json:"remark"`                // 备注
}
