package center

type OrderListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    OrderList   `json:"data,omitempty"`
}
type OrderList struct {
	OrderBy     []string `json:"orderBy"`
	List        []OrderValue `json:"list"`
	Total       int64 `json:"total"`       //总数
	PageSize    int64 `json:"pageSize"`    //每页数量
	Current     int64 `json:"current"`     //当前页码
	CurrentSize int64 `json:"currentSize"` //当前页数量
	Status      int  `json:"status"`       //订单状态：
}

type OrderValue struct {
	Id          int64 `json:"id"`         // 主键自动增长id
	Code        string `json:"code"`      // 订单编号
	CreateTime  int64 `json:"createTime"` // 创建时间
	Image       string `json:"image"`     // 图片
	Price       int64 `json:"price"`      // 单价
	Num         int64 `json:"num"`        // 数量
	TotalPrice  int64 `json:"totalPrice"` // 总价
	ProductType int `json:"productType"`  // 分类
	ProductId   int64 `json:"productId"`  // 分类
	Brief       string `json:"brief"`     // 图片
	Status      int `json:"status"`       // 状态
}

type OrderDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    OrderDetail   `json:"data,omitempty"`
}

type OrderDetail struct {
	Price   int64 `json:"price"`     // 单价
	Num     int64 `json:"num"`       // 数量
	Status  int  `json:"status"`     // 订单状态
	PayType int  `json:"pay_type"`   // 支付类型
	CodeUrl string  `json:"codeUrl"` //
}

type OrderAddResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    OrderAdd  `json:"data,omitempty"`
}

type OrderAdd struct {
	Id      int64 `json:"id"`
	PayType int `json:payType`
	CodeUrl string  `json:"codeUrl"`
}

type OrderPayResult struct {
	TradeState string  `json:"tradeState"`
}

type OrderPayResultResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    OrderPayResult  `json:"data,omitempty"`
}

type OrderCancelResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}
