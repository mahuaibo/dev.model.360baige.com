package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Card struct {
	Id         int64 `db:"id" json:"id"`                // 主键自动增长id
	CreateTime int64 `db:"create_time" json:"create_time"`            // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"`            // 更新时间
	PersonId   string `db:"person_id" json:"person_id"` // 身份Id
	Cardno     string `db:"cardno" json:"cardno"`       // 卡号
	Physicsno  string `db:"physicsno" json:"physicsno"` // 物理号
	Status     int8 `db:"status" json:"status"`         // 状态: -1删除  0：无效卡 1:入库  2：出库  3：未绑定  4：已绑定
}



func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("card", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_card?charset=utf8", 30)
	orm.RegisterModel(&User{}, &UserPosition{})
}