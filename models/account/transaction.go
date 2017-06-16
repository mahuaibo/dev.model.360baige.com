package account

type Transaction struct {
	Id            int64 `db:"id" json:"id"`                           // 主键
	CreateTime    int64 `db:"create_time" json:"create_time"`         // 创建时间
	UpdateTime    int64 `db:"update_time" json:"update_time"`         // 更新时间
	FromAccountId int64 `db:"from_account_id" json:"from_account_id"` // 交易人id
	ToAccountId   int64 `db:"to_account_id" json:"to_account_id"`     // 收款人id
	Amount        float64 `db:"amount" json:"amount"`                 // 交易金额
	OrderCode     string `db:"order_code" json:"order_code"`          // 订单编码
	Remark        string `db:"remark" json:"remark"`                  // 备注
	Status        int8 `db:"status" json:"status"`                    // 状态
}
