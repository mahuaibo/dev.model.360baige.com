package center

import (
	"dev.model.360baige.com/models/account"
)

type AccountItemStatisticsCond struct {
	AccountId int64 `json:"account_id"`     // 账户id
	StartTime int64 `json:"start_time"`     // 开始时间
	EndTime   int64 `json:"end_time"`       // 结束时间
	Income    float64 `json:"month_income"` // 收入
	Pay       float64 `json:"month_pay"`    // 支出
}
type AccountItemListResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
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
type AccountItemListPaginator struct {
	Cond        []CondValue
	Cols        []string
	OrderBy     []string
	List        []account.AccountItem
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
type AccountItemListValue struct {
	Id         int64 `json:"id"`           // 主键自动增长id
	CreateTime string `json:"create_time"` // 创建时间
	Amount     float64 `json:"amount"`     // 交易金额
	AmountType string `json:"amount_type"` // 交易金额类型
}

type AccountItemDetailResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    AccountItemDetail   `json:"data,omitempty"`
}

type AccountItemDetail struct {
	CreateTime string `json:"create_time"` // 创建时间
	ToAccount  string `json:"to_account"`  // 收款人公司
	OrderCode  string `json:"order_code"`  // 订单编码
	Amount     float64 `json:"amount"`     // 交易金额
	AmountType string `json:"amount_type"` // 交易金额类型
	Balance    float64 `json:"balance"`    // 平账
	Remark     string `json:"remark"`      // 备注
}
