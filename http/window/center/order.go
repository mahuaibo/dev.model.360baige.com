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
	Status      int8  `json:"status"`      //订单状态：-2 全部 0:撤回 1：待审核 2：已通过 3：未通过 4：发货中 5：完成
}

type OrderValue struct {
	Id         int64 `json:"id"`          // 主键自动增长id
	CreateTime string `json:"createTime"` // 创建时间
	Code       string `json:"code"`       // 订单编号
	Price      float64 `json:"price"`     // 单价
	Type       int8 `json:"type"`         // 订单类型
	PayType    string `json:"payType"`    // 支付方式 1在线支付   2线下支付
	Brief      string `json:"brief"`      // 详情
	Status     string `json:"status"`     // 订单状态：0:撤回 1：待审核  2：已通过 3：未通过 4：发货中 5：完成
}

type OrderDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    OrderDetail   `json:"data,omitempty"`
}

type OrderDetail struct {
	CreateTime string `json:"createTime"` // 创建时间
	Code       string `json:"code"`       // 订单编号
	Price      float64 `json:"price"`     // 单价
	Type       int8 `json:"type"`         // 订单类型
	PayType    string `json:"payType"`    // 支付方式 1在线支付   2线下支付
	Brief      string `json:"brief"`      // 详情
	Status     string  `json:"status"`    // 订单状态：0:撤回 1：待审核  2：已通过 3：未通过 4：发货中 5：完成
}

type OrderAddResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
}

func GetStatus(status int8) string {
	var rStatus string
	switch  status {
	case 0:
		rStatus = "撤回"
	case 1:
		rStatus = "待审核"
	case 2:
		rStatus = "已通过"
	case 3:
		rStatus = "未通过"
	case 4:
		rStatus = "发货中"
	default:
		rStatus = "完成"
	}
	return rStatus
}
