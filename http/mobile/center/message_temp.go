package center

type MessageListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    MessageList   `json:"data,omitempty"`
}
type MessageList struct {
	OrderBy     []string
	List        []MessageListValue
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
type MessageListValue struct {
	Id         int64 `db:"id" json:"id"`                    // 主键自动增长id
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间
	Content string `db:"content" json:"content"` // 内容
}

type MessageDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    MessageDetail   `json:"data,omitempty"`
}

type MessageDetail struct {
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间
	ToAccount  string `db:"to_account" json:"to_account"`   // 收款人公司
	OrderCode  string `db:"order_code" json:"order_code"`   // 订单编码
	Amount     float64 `db:"amount" json:"amount"`          // 交易金额
	AmountType string `db:"amount_type" json:"amount_type"` // 交易金额类型
	Balance    float64 `db:"balance" json:"balance"`        // 平账
	Remark     string `db:"remark" json:"remark"`           // 备注
}
