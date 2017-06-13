package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("account", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_account?charset=utf8", 30)
	orm.RegisterModel(&Account{})
}

type Account struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	UserId     int64 `db:"user_id" json:"user_id"`         // 所有者id
	Type       int8 `db:"type" json:"type"`                // 类型:1.余额；2.积分;
	Unit       int8 `db:"unit" json:"unit"`                // 单元:1.分；2.元
	Balance    float64 `db:"balance" json:"balance"`       // 余额
	Status     int8 `db:"status" json:"status"`            // 状态 -1删除
}

type AccountItem struct {
	Id            int64 `db:"id" json:"id"`                         // 主键自动增长id
	CreateTime    int64 `db:"create_time" json:"create_time"`         // 创建时间
	UpdateTime    int64 `db:"update_time" json:"update_time"`         // 更新时间
	TransactionId int64 `db:"transaction_id" json:"transaction_id"` // 交易id
	AccountId     int64 `db:"account_id" json:"account_id"`         // 账户id
	Amount        float64 `db:"amount" json:"amount"`               // 交易金额
	Balance       float64 `db:"balance" json:"balance"`             // 平账
	Remark        string `db:"remark" json:"remark"`                // 备注
}

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
