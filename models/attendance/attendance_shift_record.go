package attendance

type AttendanceShiftRecord struct {
	Id                    int64 `db:"id" json:"id"`                                             // 主键
	CreateTime            int64 `db:"create_time" json:"create_time"`                           // 创建时间
	UpdateTime            int64 `db:"update_time" json:"update_time"`                           // 更新时间
	DateTime              int64 `db:"date_time" json:"date_time"`                               // 某天的时间戳
	CompanyId             int64 `db:"company_id" json:"company_id"`                             // 所有者ID
	UserId                int64 `db:"user_id" json:"user_id"`                                   // 用户id
	ShiftItemId           int64 `db:"shift_item_id" json:"shift_item_id"`                       // 班次项ID（签到）
	AttendanceInRecordId  int64 `db:"attendance_in_record_id" json:"attendance_in_record_id"`   // 签到记录ID
	AttendanceOutRecordId int64 `db:"attendance_out_record_id" json:"attendance_out_record_id"` // 签退记录ID
	AttendanceShiftItemId int64 `db:"attendance_shift_item_id" json:"attendance_shift_item_id"` // 考勤班次项
}
