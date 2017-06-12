package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Application struct {
	ID               int64 `db:"id" json:"id"`                                   // 主键
	CreateTime       int64 `db:"create_time" json:"create_time"`                   // 创建时间
	UpdateTime       int64 `db:"update_time" json:"update_time"`                   // 更新时间
	CompanyID        int64 `db:"company_id" json:"company_id"`                   // 所属公司ID
	UserID           int64 `db:"user_id" json:"user_id"`                         // 购买者ID
	ApplicationTplID float64 `db:"application_tpl_id" json:"application_tpl_id"` // 应用ID
	Name             string `db:"name" json:"name"`                              // 名称
	Type             int8 `db:"type" json:"type"`                                // 类型
	Status           int8 `db:"status" json:"status"`                            // 状态
}

type ApplicationTpl struct {
	ID         int64 `db:"id" json:"id"`                 // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	CompanyID  int64 `db:"company_id" json:"company_id"` // 开发公司ID
	UserID     int64 `db:"user_id" json:"user_id"`       // 开发者ID
	Name       string `db:"name" json:"name"`            // 名称
	Image      string `db:"image" json:"image"`          // 图片链接
	Site       string `db:"site" json:"site"`            // 访问链接
	Type       int8 `db:"type" json:"type"`              // 类型
	Desc       string `db:"desc" json:"desc"`            // 简介
	Status     int8 `db:"status" json:"status"`          // 状态
	Price      float64 `db:"price" json:"price"`         // 价格
}