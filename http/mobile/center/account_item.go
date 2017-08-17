package center

type AccountItemStatisticsCond struct {
	AccountId int64 `db:"account_id" json:"account_id"` // 账户id
	StartTime int64 `db:"start_time" json:"start_time"` // 开始时间
	EndTime   int64 `db:"end_time" json:"end_time"`     // 结束时间
	Income    float64 `db:"amount" json:"month_income"` // 收入
	Pay       float64 `db:"amount" json:"month_pay"`    // 支出
}
type AccountItemListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountItemList   `json:"data,omitempty"`
}
type AccountItemList struct {
	OrderBy     []string
	List        []AccountItemListValue
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
type AccountItemListValue struct {
	Id         int64 `db:"id" json:"id"`                    // 主键自动增长id
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间
	Amount     float64 `db:"amount" json:"amount"`          // 交易金额
	AmountType string `db:"amount_type" json:"amount_type"` // 交易金额类型
}

type AccountItemDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountItemDetail   `json:"data,omitempty"`
}

type AccountItemDetail struct {
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间
	ToAccount  string `db:"to_account" json:"to_account"`   // 收款人公司
	OrderCode  string `db:"order_code" json:"order_code"`   // 订单编码
	Amount     float64 `db:"amount" json:"amount"`          // 交易金额
	AmountType string `db:"amount_type" json:"amount_type"` // 交易金额类型
	Balance    float64 `db:"balance" json:"balance"`        // 平账
	Remark     string `db:"remark" json:"remark"`           // 备注
}
