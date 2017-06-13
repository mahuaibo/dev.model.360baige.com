package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Logger struct {
	Id         int64 `db:"id" json:"id"`                   // 主键 id
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 公司id
	OwnerId    int64 `db:"owner_id" json:"owner_id"`       // 操作者Id
	Remark     string `db:"remark" json:"remark"`          // 描述
	Content    string `db:"content" json:"content"`        // 内容(修改前后变化)
	Type       int `db:"type" json:"type"`                 // 类别（增、删、改、查、其他）
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("logger", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_logger?charset=utf8", 30)
	orm.RegisterModel(&Logger{})
}

// 新增
func (*Logger) AddLogger(args *Logger, reply *Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.CompanyId = args.CompanyId
		reply.OwnerId = args.OwnerId
		reply.Remark = args.Remark
		reply.Content = args.Content
		reply.Type = args.Type
	}
	return err
}

// 查询 by Id
func (*Logger) FindLoggerById(args *Logger, reply *Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}
