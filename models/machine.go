package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Machine struct {
	Id         int64 `db:"id" json:"id"`                 // 主键
	CreateTime int64 `db:"create_time" json:"create_time"`           // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"`           // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"` // 所有者Id
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("machine", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_machine?charset=utf8", 30)
	orm.RegisterModel(&Logger{})
}