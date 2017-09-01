package account

type Transaction struct {
	Id            int64    `db:"id" json:"id"`                         // 主键
	CreateTime    int64    `db:"create_time" json:"createTime"`        // 创建时间
	UpdateTime    int64    `db:"update_time" json:"updateTime"`        // 更新时间
	FromAccountId int64    `db:"from_account_id" json:"fromAccountId"` // 交易人id
	ToAccountId   int64    `db:"to_account_id" json:"toAccountId"`     // 收款人id
	Amount        int64    `db:"amount" json:"amount"`                 // 交易金额
	OrderCode     string    `db:"order_code" json:"orderCode"`         // 订单编码
	Remark        string    `db:"remark" json:"remark"`                // 备注
	Status        int8    `db:"status" json:"status"`                  // 状态
}
