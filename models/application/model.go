package application

type Application struct {
	Id               int64    `db:"id" json:"id"`                               // 主键
	CreateTime       int64    `db:"create_time" json:"createTime"`              // 创建时间
	UpdateTime       int64    `db:"update_time" json:"updateTime"`              // 更新时间
	CompanyId        int64    `db:"company_id" json:"companyId"`                // 所属公司ID
	UserId           int64    `db:"user_id" json:"userId"`                      // 购买者ID
	ApplicationTplId int64    `db:"application_tpl_id" json:"applicationTplId"` // 应用ID
	UserPositionType int8    `db:"user_position_type" json:"userPositionType"`  // 身份类型
	UserPositionId   int64    `db:"user_position_id" json:"userPositionId"`     // 身份ID
	Name             string    `db:"name" json:"name"`                          // 名称
	Image            string    `db:"image" json:"image"`                        // 图片链接
	Status           int8    `db:"status" json:"status"`                        // 状态 0停用  1启用  2退订
	StartTime        int64    `db:"start_time" json:"startTime"`                // 开始时间
	EndTime          int64    `db:"end_time" json:"endTime"`                    // 结束时间
}

type ApplicationTpl struct {
	Id           int64    `db:"id" json:"id"`                     // 主键
	CreateTime   int64    `db:"create_time" json:"createTime"`    // 创建时间（毫秒）
	UpdateTime   int64    `db:"update_time" json:"updateTime"`    // 修改时间（毫秒）
	CompanyId    int64    `db:"company_id" json:"companyId"`      // 开发公司ID
	UserId       int64    `db:"user_id" json:"userId"`            // 开发者ID
	Name         string    `db:"name" json:"name"`                // 名称
	Image        string    `db:"image" json:"image"`              // 图片链接
	Site         string    `db:"site" json:"site"`                // 访问链接
	Type         int8    `db:"type" json:"type"`                  // 类型 1天气  2体育  3旅游  4工具  5视频  6社交  7参考  8效率  9摄影  10新闻  11音乐  12医学  13生活  14健康健美  15游戏  16财经  17娱乐  18教育  19商务  20图形设计  21软件开发
	Desc         string    `db:"desc" json:"desc"`                // 简介
	Status       int8    `db:"status" json:"status"`              // 状态 0下架
	Price        int64    `db:"price" json:"price"`               // 价格
	PayType      int8    `db:"pay_type" json:"payType"`           // 0:限免 1:永久免费 2:1次性收费 3:周期收费
	PayCycle     int8    `db:"pay_cycle" json:"payCycle"`         // 0无1月2季3半年4年
	Subscription int64    `db:"subscription" json:"subscription"` // 订阅量
}
