package personnel

type Person struct {
	Id         int64 `db:"id" json:"id"`                   // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`   // 所有者ID
	Code       string`db:"code" json:"code"`               // 编号
	Name       string `db:"name" json:"name"`              // 名称
	Sex        string `db:"sex" json:"sex"`                // 性别
	Birthday   int64 `db:"birthday" json:"birthday"`       // 生日
	Type       int8 `db:"type" json:"type"`                // 类型 1.教师   2.家长   3.学生
	Phone      string `db:"phone" json:"phone"`      // 联系方式
	Contact    string `db:"contact" json:"contact"`      // 联系人
	Status     int8 `db:"status" json:"status"`            // 状态 1.正常   2.异常
}
