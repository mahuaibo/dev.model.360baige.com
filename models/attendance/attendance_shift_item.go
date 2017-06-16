package attendance

type AttendanceShiftItem struct {
	Id                int64 `db:"id" json:"id"`                                   // 主键(班次ID)
	CreateTime        int64 `db:"create_time" json:"create_time"`                 // 创建时间
	UpdateTime        int64 `db:"update_time" json:"update_time"`                 // 更新时间
	CompanyId         int64 `db:"company_id" json:"company_id"`                   // 所有者ID
	AttendanceShiftId int64 `db:"attendance_shift_id" json:"attendance_shift_id"` // shift id
	Pos               int8 `db:"pos" json:"pos"`                                  // 顺序
	SignIn            int64 `db:"sign_in" json:"sign_in"`                         // 标准签到时间
	SignInStart       int64 `db:"sign_in_start" json:"sign_in_start"`             // 签到时间开始
	SignInEnd         int64 `db:"sign_in_end" json:"sign_in_end"`                 // 签到时间结束
	SignBack          int64 `db:"sign_back" json:"sign_back"`                     // 标准签退时间
	SignBackStart     int64 `db:"sign_back_start" json:"sign_back_start"`         // 签退时间开始
	SignBackEnd       int64 `db:"sign_back_end" json:"sign_back_end"`             // 签退时间结束
	Type              int8 `db:"type" json:"type"`                                // 人性化 1 无 2 3
	SignInErrata      int64 `db:"sign_in_errata" json:"sign_in_errata"`           // 签到勘误时间
	SignBackErrata    int64 `db:"sign_back_errata" json:"sign_back_errata"`       // 签退勘误时间
}
