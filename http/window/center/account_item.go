package center

type AccountItemStatisticsCond struct {
	AccountId int64 `json:"accountId"`     // 账户id
	StartTime int64 `json:"startTime"`     // 开始时间
	EndTime   int64 `json:"endTime"`       // 结束时间
	Income    float64 `json:"monthIncome"` // 收入
	Pay       float64 `json:"monthPay"`    // 支出
}
type AccountItemListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountItemList   `json:"data,omitempty"`
}

type AccountItemList struct {
	OrderBy     []string `json:"orderBy"`            // 排序
	List        []AccountItemListValue `json:"list"` // 集合
	Total       int64 `json:"total"`                 // 总数
	PageSize    int64 `json:"pageSize"`              // 每页数量
	Current     int64 `json:"current"`               // 当前页码
	CurrentSize int64 `json:"currentSize"`           // 当前页数量
	InAccount   float64 `json:"inAccount"`           // 进账
	OutAccount  float64 `json:"outAccount"`          // 出账
}

type AccountItemListValue struct {
	Id         int64 `json:"id"`          // 主键自动增长id
	CreateTime string `json:"createTime"` // 创建时间
	Amount     float64 `json:"amount"`    // 交易金额
	AmountType string `json:"amountType"` // 交易金额类型
}

type AccountItemDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountItemDetail   `json:"data,omitempty"`
}

type AccountItemDetail struct {
	CreateTime string `json:"createTime"` // 创建时间
	ToAccount  string `json:"toAccount"`  // 收款人公司
	OrderCode  string `json:"orderCode"`  // 订单编码
	Amount     float64 `json:"amount"`    // 交易金额
	AmountType string `json:"amountType"` // 交易金额类型
	Balance    float64 `json:"balance"`   // 平账
	Remark     string `json:"remark"`     // 备注
}
