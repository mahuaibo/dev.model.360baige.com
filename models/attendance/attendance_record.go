package attendance

type AttendanceRecord struct {
	Id          int64 `db:"id" json:"id"`                     // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`   // 创建时间
	UpdateTime  int64 `db:"update_time" json:"update_time"`   // 更新时间 考勤时间
	CompanyId   int64 `db:"company_id" json:"company_id"`     // 企业
	EquipmentId int64 `db:"equipment_id" json:"equipment_id"` // 设备
	UserId      int64 `db:"user_id" json:"user_id"`           // 人员id
	Type        int8 `db:"type" json:"type"`                  // 打卡:1 指纹 2 刷卡 3 人脸；
}
