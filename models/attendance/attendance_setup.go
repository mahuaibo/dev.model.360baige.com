package attendance

type AttendanceSetup struct {
	Id           int64 `db:"id" json:"id"`                       // 主键
	CreateTime   int64 `db:"create_time" json:"create_time"`     // 创建时间
	UpdateTime   int64 `db:"update_time" json:"update_time"`     // 更新时间
	CompanyId    int64 `db:"company_id" json:"company_id"`       // 所有者ID
	SigninStime  int64 `db:"signin_stime" json:"signin_stime"`   // 签到时间(起)
	SigninEtime  int64 `db:"signin_etime" json:"signin_etime"`   // 签到时间(止)
	SignoutStime int64 `db:"signout_stime" json:"signout_stime"` // 签退时间(起)
	SignoutEtime int64 `db:"signout_etime" json:"signout_etime"` // 签退时间(止)
	Type         int8 `db:"type" json:"type"`                    // 类型：1:标准考勤 2：时长考勤（在标准基础上）3：记录来走
	Datatype     int8 `db:"datatype" json:"datatype"`            // 记录类型 0 默认数据， 1用户数据
	Name         string `db:"name" json:"name"`                  // 时段名称
	ShortName    string `db:"short_name" json:"short_name"`      // 简称
	Color        string `db:"color" json:"color"`                // 颜色
	Timelength   int64 `db:"timelength" json:"timelength"`       // 最小工作时长
}
