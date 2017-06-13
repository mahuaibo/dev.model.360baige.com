package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type AttendanceGroup struct {
	Id                int64 `db:"id" json:"id"`                                   // 主键
	CreateTime        int64 `db:"create_time" json:"create_time"`                 // 创建时间
	UpdateTime        int64 `db:"update_time" json:"update_time"`                 // 更新时间
	CompanyId         int64 `db:"company_id" json:"company_id"`                   // 所有者Id
	AttendanceSetupId int64 `db:"attendance_setup_id" json:"attendance_setup_id"` // 考勤设置Id
	PersonId          int64 `db:"person_id" json:"person_id"`                     // 角色Id
	Name              string `db:"name" json:"name"`                              // 分组名称
	UserIds           string `db:"user_ids" json:"user_ids"`                      // 用户ids
	Status            int8 `db:"status" json:"status"`                            // 状态 0未启用  1已启用
}

type AttendanceRecord struct {
	Id          int64 `db:"id" json:"id"`                     // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`     // 创建时间
	UpdateTime  int64 `db:"update_time" json:"update_time"`     // 更新时间 考勤时间
	CompanyId   int64 `db:"company_id" json:"company_id"`     // 企业
	EquipmentId int64 `db:"equipment_id" json:"equipment_id"` // 设备
	UserId      int64 `db:"user_id" json:"user_id"`           // 人员id
	Type        int8 `db:"type" json:"type"`                  // 打卡:1 指纹 2 刷卡 3 人脸；
}

type AttendanceSetup struct {
	Id           int64 `db:"id" json:"id"`                     // 主键
	CreateTime   int64 `db:"create_time" json:"create_time"`     // 创建时间
	UpdateTime   int64 `db:"update_time" json:"update_time"`     // 更新时间
	CompanyId    int64 `db:"company_id" json:"company_id"`     // 所有者Id
	SigninStime  int64 `db:"signin_stime" json:"signin_stime"`   // 签到时间(起)
	SigninEtime  int64 `db:"signin_etime" json:"signin_etime"`   // 签到时间(止)
	SignoutStime int64 `db:"signout_stime" json:"signout_stime"` // 签退时间(起)
	SignoutEtime int64 `db:"signout_etime" json:"signout_etime"` // 签退时间(止)
	Type         int8 `db:"type" json:"type"`                  // 类型：1:标准考勤 2：时长考勤（在标准基础上）3：记录来走
	Datatype     int8 `db:"datatype" json:"datatype"`          // 记录类型 0 默认数据， 1用户数据
	Name         string `db:"name" json:"name"`                // 时段名称
	ShortName    string `db:"short_name" json:"short_name"`    // 简称
	Color        string `db:"color" json:"color"`              // 颜色
	Timelength   int64 `db:"timelength" json:"timelength"`       // 最小工作时长
}

type AttendanceShift struct {
	Id         int64 `db:"id" json:"id"`                  // 主键(班次Id)
	CreateTime int64 `db:"create_time" json:"create_time"`  // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"`  // 更新时间
	CompanyId  int64 `db:"company_id" json:"company_id"`  // 所有者Id
	Type       int8 `db:"type" json:"type"`               // 班次类型
	Name       string `db:"name" json:"name"`             // 班次名称
	ShortName  string `db:"short_name" json:"short_name"` // 班次简称
	Color      string `db:"color" json:"color"`           // 颜色
	Status     int8 `db:"status" json:"status"`           // 状态 -1删除
}

type AttendanceShiftItem struct {
	Id                int64 `db:"id" json:"id"`                                   // 主键(班次Id)
	CreateTime        int64 `db:"create_time" json:"create_time"`                   // 创建时间
	UpdateTime        int64 `db:"update_time" json:"update_time"`                   // 更新时间
	CompanyId         int64 `db:"company_id" json:"company_id"`                   // 所有者Id
	AttendanceShiftId int64 `db:"attendance_shift_id" json:"attendance_shift_id"` // shift id
	Pos               int8 `db:"pos" json:"pos"`                                  // 顺序
	SignIn            int64 `db:"sign_in" json:"sign_in"`                           // 标准签到时间
	SignInStart       int64 `db:"sign_in_start" json:"sign_in_start"`               // 签到时间开始
	SignInEnd         int64 `db:"sign_in_end" json:"sign_in_end"`                   // 签到时间结束
	SignBack          int64 `db:"sign_back" json:"sign_back"`                       // 标准签退时间
	SignBackStart     int64 `db:"sign_back_start" json:"sign_back_start"`           // 签退时间开始
	SignBackEnd       int64 `db:"sign_back_end" json:"sign_back_end"`               // 签退时间结束
	Type              int8 `db:"type" json:"type"`                                // 人性化 1 无 2 3
	SignInErrata      int64 `db:"sign_in_errata" json:"sign_in_errata"`             // 签到勘误时间
	SignBackErrata    int64 `db:"sign_back_errata" json:"sign_back_errata"`         // 签退勘误时间
}

type AttendanceShiftRecord struct {
	Id                    int64 `db:"id" json:"id"`                                             // 主键
	CreateTime            int64 `db:"create_time" json:"create_time"`                             // 创建时间
	UpdateTime            int64 `db:"update_time" json:"update_time"`                             // 更新时间
	DateTime              int64 `db:"date_time" json:"date_time"`                                 // 某天的时间戳
	CompanyId             int64 `db:"company_id" json:"company_id"`                             // 所有者Id
	UserId                int64 `db:"user_id" json:"user_id"`                                   // 用户id
	ShiftItemId           int64 `db:"shift_item_id" json:"shift_item_id"`                       // 班次项Id（签到）
	AttendanceInRecordId  int64 `db:"attendance_in_record_id" json:"attendance_in_record_id"`   // 签到记录Id
	AttendanceOutRecordId int64 `db:"attendance_out_record_id" json:"attendance_out_record_id"` // 签退记录Id
	AttendanceShiftItemId int64 `db:"attendance_shift_item_id" json:"attendance_shift_item_id"` // 考勤班次项
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("attendance", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_attendance?charset=utf8", 30)
	orm.RegisterModel(&User{}, &UserPosition{})
}
