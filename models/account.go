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
	UserId     int64 `db:"userId" json:"userId"`         // 所有者id
	Type       int8 `db:"type" json:"type"`                // 类型:1.余额；2.积分;
	Unit       int8 `db:"unit" json:"unit"`                // 单元:1.分；2.元
	Balance    float64 `db:"balance" json:"balance"`       // 余额
	Status     int8 `db:"status" json:"status"`            // 状态 -1删除
}

type AccountItem struct {
	Id            int64 `db:"id" json:"id"`                         // 主键自动增长id
	CreateTime    int64 `db:"create_time" json:"create_time"`         // 创建时间
	UpdateTime    int64 `db:"update_time" json:"update_time"`         // 更新时间
	TransactionId int64 `db:"transactionId" json:"transactionId"` // 交易id
	AccountId     int64 `db:"accountId" json:"accountId"`         // 账户id
	Amount        float64 `db:"amount" json:"amount"`               // 交易金额
	Balance       float64 `db:"balance" json:"balance"`             // 平账
	Remark        string `db:"remark" json:"remark"`                // 备注
}

type Transaction struct {
	Id            int64 `db:"id" json:"id"`                           // 主键
	CreateTime    int64 `db:"create_time" json:"create_time"`         // 创建时间
	UpdateTime    int64 `db:"update_time" json:"update_time"`         // 更新时间
	FromAccountId int64 `db:"from_accountId" json:"from_accountId"` // 交易人id
	ToAccountId   int64 `db:"to_accountId" json:"to_accountId"`     // 收款人id
	Amount        float64 `db:"amount" json:"amount"`                 // 交易金额
	OrderCode     string `db:"order_code" json:"order_code"`          // 订单编码
	Remark        string `db:"remark" json:"remark"`                  // 备注
	Status        int8 `db:"status" json:"status"`                    // 状态
}


/**
 * ******* Account
 */
// 新增
func (*Account) Add(args *Account, reply *Account) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.UserId = args.UserId
		reply.Type = args.Type
		reply.Unit = args.Unit
		reply.Balance = args.Balance
		reply.Status = args.Status
	}
	return err
}

// 查询 by Id
func (*Account) FindById(args *Account, reply *Account) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*Account) UpdateById(args *Account, reply *Account) error {
	o := orm.NewOrm()
	o.Using("account")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}

/**
 * ******* AccountItem
 */
// 新增
func (*AccountItem) Add(args *AccountItem, reply *AccountItem) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.TransactionId = args.TransactionId
		reply.AccountId = args.AccountId
		reply.Amount = args.Amount
		reply.Balance = args.Balance
		reply.Remark = args.Remark
	}
	return err
}

// 查询 by Id
func (*AccountItem) FindById(args *AccountItem, reply *AccountItem) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AccountItem) UpdateById(args *AccountItem, reply *AccountItem) error {
	o := orm.NewOrm()
	o.Using("account")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}

/**
 * ******* Transaction
 */
// 新增
func (*Transaction) Add(args *Transaction, reply *Transaction) error {
	o := orm.NewOrm()
	o.Using("account")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.FromAccountId = args.FromAccountId
		reply.ToAccountId = args.ToAccountId
		reply.Amount = args.Amount
		reply.OrderCode = args.OrderCode
		reply.Remark = args.Remark
		reply.Status = args.Status
	}
	return err
}

// 查询 by Id
func (*Transaction) FindById(args *Transaction, reply *Transaction) error {
	o := orm.NewOrm()
	o.Using("account")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*Transaction) UpdateById(args *Transaction, reply *Transaction) error {
	o := orm.NewOrm()
	o.Using("account")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}