package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Order struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 企业Id
	UserId     int64 `db:"user_id" json:"user_id"`         // 用户Id
	Code       string `db:"code" json:"code"`              // 订单编号
	Price      float64 `db:"price" json:"price"`           // 单价
	Type       int8 `db:"type" json:"type"`                // 订单类型
	PayType    int8 `db:"pay_type" json:"pay_type"`        // 支付方式 1在线支付   2线下支付
	Brief      string `db:"brief" json:"brief"`            // 详情
	Status     int8 `db:"status" json:"status"`            // 订单状态：0:撤回 1：待审核  2：已通过 3：未通过 4：发货中 5：完成
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("order", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_order?charset=utf8", 30)
	orm.RegisterModel(&Logger{})
}