package attendance

type AttendanceGroup struct {
	Id                int64 `db:"id" json:"id"`                                   // 主键
	CreateTime        int64 `db:"create_time" json:"create_time"`                 // 创建时间
	UpdateTime        int64 `db:"update_time" json:"update_time"`                 // 更新时间
	CompanyId         int64 `db:"company_id" json:"company_id"`                   // 所有者ID
	AttendanceSetupId int64 `db:"attendance_setup_id" json:"attendance_setup_id"` // 考勤设置ID
	PersonId          int64 `db:"person_id" json:"person_id"`                     // 角色ID
	Name              string `db:"name" json:"name"`                              // 分组名称
	UserIds           string `db:"user_ids" json:"user_ids"`                      // 用户ids
	Status            int8 `db:"status" json:"status"`                            // 状态 0未启用  1已启用
}
