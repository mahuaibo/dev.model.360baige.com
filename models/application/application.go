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
