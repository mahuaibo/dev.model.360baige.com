package app

type MessageReminder struct {
	Id int64 `db:"id" json:"id"` // 主键自动增长id
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
}
