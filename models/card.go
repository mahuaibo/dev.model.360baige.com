package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type Card struct {
	ID         int64 `db:"id" json:"id"`                // 主键自动增长id
	CreateTime int64 `db:"create_time" json:"create_time"`            // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"`            // 更新时间
	PersonID   string `db:"person_id" json:"person_id"` // 身份ID
	Cardno     string `db:"cardno" json:"cardno"`       // 卡号
	Physicsno  string `db:"physicsno" json:"physicsno"` // 物理号
	Status     int8 `db:"status" json:"status"`         // 状态: -1删除  0：无效卡 1:入库  2：出库  3：未绑定  4：已绑定
}
