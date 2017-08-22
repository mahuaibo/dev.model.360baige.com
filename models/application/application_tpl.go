package application

type ApplicationTpl struct {
	Id         int64    `db:"id" json:"id"`                  // 主键
	CreateTime int64    `db:"create_time" json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64    `db:"update_time" json:"updateTime"` // 修改时间（毫秒）
	CompanyId  int64    `db:"company_id" json:"companyId"`   // 开发公司ID
	UserId     int64    `db:"user_id" json:"userId"`         // 开发者ID
	Name       string    `db:"name" json:"name"`             // 名称
	Image      string    `db:"image" json:"image"`           // 图片链接
	Site       string    `db:"site" json:"site"`             // 访问链接
	Type       int8    `db:"type" json:"type"`               // 类型 1天气  2体育  3旅游  4工具  5视频  6社交  7参考  8效率  9摄影  10新闻  11音乐  12医学  13生活  14健康健美  15游戏  16财经  17娱乐  18教育  19商务  20图形设计  21软件开发
	Desc       string    `db:"desc" json:"desc"`             // 简介
	Status     int8    `db:"status" json:"status"`           // 状态 0下架
	Price      float64    `db:"price" json:"price"`          // 价格
	PayType    int8    `db:"pay_type" json:"payType"`        // 0:限免 1:永久免费 2:1次性收费 3:周期收费
	PayCycle   int8    `db:"pay_cycle" json:"payCycle"`      // 0无1月2季3半年4年
}
