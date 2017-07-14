package window

import (
	"dev.model.360baige.com/models/order"
)

type OrderListResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    OrderList   `json:"data,omitempty"`
}
type OrderList struct {
	OrderBy     []string
	List        []OrderValue
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
	Status      int8  //订单状态：-2 全部 0:撤回 1：待审核 2：已通过 3：未通过 4：发货中 5：完成
}
type OrderListPaginator struct {
	Cond        []CondValue
	Cols        []string
	OrderBy     []string
	List        []order.Order
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
type OrderValue struct {
	Id         int64 `db:"id" json:"id"`                    // 主键自动增长id
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间
	Code       string `db:"code" json:"code"`               // 订单编号
	Price      float64 `db:"price" json:"price"`            // 单价
	Type       int8 `db:"type" json:"type"`                 // 订单类型
	PayType    string `db:"pay_type" json:"pay_type"`         // 支付方式 1在线支付   2线下支付
	Brief      string `db:"brief" json:"brief"`             // 详情
	Status     string `db:"status" json:"status"`             // 订单状态：0:撤回 1：待审核  2：已通过 3：未通过 4：发货中 5：完成
}

type OrderDetailResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    OrderDetail   `json:"data,omitempty"`
}

type OrderDetail struct {
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间
	Code       string `db:"code" json:"code"`               // 订单编号
	Price      float64 `db:"price" json:"price"`            // 单价
	Type       int8 `db:"type" json:"type"`                 // 订单类型
	PayType    string `db:"pay_type" json:"pay_type"`         // 支付方式 1在线支付   2线下支付
	Brief      string `db:"brief" json:"brief"`             // 详情
	Status     string  `db:"status" json:"status"`             // 订单状态：0:撤回 1：待审核  2：已通过 3：未通过 4：发货中 5：完成
}
