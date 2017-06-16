package application

type Application struct {
	Id               int64 `db:"id" json:"id"`                                 // 主键
	CreateTime       int64 `db:"create_time" json:"create_time"`               // 创建时间
	UpdateTime       int64 `db:"update_time" json:"update_time"`               // 更新时间
	CompanyId        int64 `db:"company_id" json:"company_id"`                 // 所属公司ID
	UserId           int64 `db:"user_id" json:"user_id"`                       // 购买者ID
	ApplicationTplId int64 `db:"application_tpl_id" json:"application_tpl_id"` // 应用ID
	Name             string `db:"name" json:"name"`                            // 名称
	Type             int8 `db:"type" json:"type"`                              // 类型
	Status           int8 `db:"status" json:"status"`                          // 状态
}
