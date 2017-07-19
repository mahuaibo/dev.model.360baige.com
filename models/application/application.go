package application

type Application struct {
	Id               int64 `db:"id" json:"id"`                                 // 主键
	CreateTime       int64 `db:"create_time" json:"create_time"`               // 创建时间
	UpdateTime       int64 `db:"update_time" json:"update_time"`               // 更新时间
	CompanyId        int64 `db:"company_id" json:"company_id"`                 // 所属公司ID
	UserId           int64 `db:"user_id" json:"user_id"`                       // 购买者ID
	UserPositionId   int64 `db:"user_position_id" json:"user_position_id"`     // 人事ID
	UserPositionType int8 `db:"user_position_type" json:"user_position_type"`  // 用户类型
	ApplicationTplId int64 `db:"application_tpl_id" json:"application_tpl_id"` // 应用ID
	Name             string `db:"name" json:"name"`                            // 名称
	Image            string `db:"image" json:"image"`                          // 图片链接
	Status           int8 `db:"status" json:"status"`                          // 状态
	StartTime        int64 `db:"start_time" json:"start_time"`                 // 开始时间
	EndTime          int64 `db:"end_time" json:"end_time"`                     // 结束时间
}
