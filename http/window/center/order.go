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
	Status      int8  `json:"status"`      //订单状态：
}

type OrderValue struct {
	Id          int64 `json:"id"`          // 主键自动增长id
	Code        string `json:"code"`       // 订单编号
	CreateTime  string `json:"createTime"` // 创建时间
	Image       string `json:"image"`      // 图片
	Price       int64 `json:"price"`       // 单价
	Num         int64 `json:"num"`         // 数量
	TotalPrice  int64 `json:"totalPrice"`  // 总价
	ProductType int8 `json:"productType"`  // 分类
	ProductId   int64 `json:"productId"`   // 分类
	Status      int8 `json:"status"`       // 状态
}

type OrderDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    OrderDetail   `json:"data,omitempty"`
}

type OrderDetail struct {
	CreateTime  string `json:"createTime"` // 创建时间
	Code        string `json:"code"`       // 订单编号
	Price       int64 `json:"price"`       // 单价
	ProductType int8 `json:"productType"`  // 订单类型
	PayType     int8 `json:"payType"`      // 支付方式
	Brief       string `json:"brief"`      // 详情
	Status      int8  `json:"status"`      // 订单状态
}

type OrderAddResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    OrderAdd  `json:"data,omitempty"`
}

type OrderAdd struct {
	Id      int64 `json:"id"`
	CodeUrl string  `json:"codeUrl"`
}

type OrderPayResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    OrderPay  `json:"data,omitempty"`
}

type OrderPay struct {
	Code      string   `json:"code"`
	Price     int64   `json:"price"`
	Brief     string   `json:"brief"`
	TradeType string   `json:"tradeType"` // 交易类型
	PrepayId  string   `json:"prepayId"`  // 预支付交易会话标识
	CodeUrl   string   `json:"codeUrl"`   // 二维码链接
	Openid    string   `json:"openid"`    // 用户标识
}

type OrderPayResultResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
