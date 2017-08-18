package authority

type GatherItem struct {
	Id            int64 `db:"id" json:"id"`                         // 主键
	CreateTime    int64 `db:"create_time" json:"create_time"`       // 创建时间（毫秒）
	UpdateTime    int64 `db:"update_time" json:"update_time"`       // 更新时间（毫秒）
	GatherId      int64 `db:"gather_id" json:"gather_id"`           // 集合ID
	ApplicationId int64 `db:"application_id" json:"application_id"` // 应用ID
	Status        int8 `db:"status" json:"status"`                  // 状态 -1注销 0正常
}
