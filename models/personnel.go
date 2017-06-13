package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Person struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 所有者Id
	Name       string `db:"name" json:"name"`              // 名称
	Type       int8 `db:"type" json:"type"`                // 类型 1.教师   2.家长   3.学生
	Status     int8 `db:"status" json:"status"`            // 状态 1.正常   2.异常
}

type PersonRelation struct {
	Id            int64 `db:"id" json:"id"`                         // 主键
	CreateTime    int64 `db:"create_time" json:"create_time"`       // 创建时间（毫秒）
	UpdateTime    int64 `db:"update_time" json:"update_time"`       // 更新时间（毫秒）
	CompanyId     int64 `db:"company_id" json:"company_id"`         // 所有者Id
	AssociationId int64 `db:"association_id" json:"association_id"` // parent_id（关联人Id）
	AssociatedId  int64 `db:"associated_id" json:"associated_id"`   // parent_id（被关联人Id）
	Type          int8 `db:"type" json:"type"`                      // 类型
	Status        int8 `db:"status" json:"status"`                  // 状态  1.关联   2.取消关联
}

type PersonStructure struct {
	Id          int64 `db:"id" json:"id"`                     // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`   // 创建时间（毫秒）
	UpdateTime  int64 `db:"update_time" json:"update_time"`   // 更新时间（毫秒）
	CompanyId   int64 `db:"company_id" json:"company_id"`     // 所有者Id
	PersonId    int64 `db:"person_id" json:"person_id"`       // parent_id 人员Id
	StructureId int64 `db:"structure_id" json:"structure_id"` // structure_id 结构Id
	Type        int8 `db:"type" json:"type"`                  // 类型
	Status      int8 `db:"status" json:"status"`              // 状态  1.正常
}

type Structure struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 所有者Id
	ParentId   int64 `db:"parent_id" json:"parent_id"`     // parent_id
	Name       string `db:"name" json:"name"`              // 名称
	Type       int8 `db:"type" json:"type"`                // 类型 1.班级
	Status     int8 `db:"status" json:"status"`            // 状态 0.下线  1.上线
}


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("personnel", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_personnel?charset=utf8", 30)
	orm.RegisterModel(&Logger{})
}
