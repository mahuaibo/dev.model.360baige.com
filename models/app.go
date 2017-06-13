package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Application struct {
	Id               int64 `db:"id" json:"id"`                                   // 主键
	CreateTime       int64 `db:"create_time" json:"create_time"`                   // 创建时间
	UpdateTime       int64 `db:"update_time" json:"update_time"`                   // 更新时间
	CompanyId        int64 `db:"company_id" json:"company_id"`                   // 所属公司Id
	UserId           int64 `db:"user_id" json:"user_id"`                         // 购买者Id
	ApplicationTplId float64 `db:"application_tpl_id" json:"application_tpl_id"` // 应用Id
	Name             string `db:"name" json:"name"`                              // 名称
	Type             int8 `db:"type" json:"type"`                                // 类型
	Status           int8 `db:"status" json:"status"`                            // 状态
}

type ApplicationTpl struct {
	Id         int64 `db:"id" json:"id"`                 // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	CompanyId  int64 `db:"company_id" json:"company_id"` // 开发公司Id
	UserId     int64 `db:"user_id" json:"user_id"`       // 开发者Id
	Name       string `db:"name" json:"name"`            // 名称
	Image      string `db:"image" json:"image"`          // 图片链接
	Site       string `db:"site" json:"site"`            // 访问链接
	Type       int8 `db:"type" json:"type"`              // 类型
	Desc       string `db:"desc" json:"desc"`            // 简介
	Status     int8 `db:"status" json:"status"`          // 状态
	Price      float64 `db:"price" json:"price"`         // 价格
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("application", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_application?charset=utf8", 30)
	orm.RegisterModel(&User{}, &UserPosition{})
}