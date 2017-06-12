package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type User struct {
	ID int64 `db:"id" json:"id"` // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"` // 修改时间（毫秒）
	CompanyID int64 `db:"company_id" json:"company_id"` // 学校ID
	Username string `db:"username" json:"username"` // 用户名
	Password string `db:"password" json:"password"` // 密码
	Email string `db:"email" json:"email"` // 邮件
	Phone string `db:"phone" json:"phone"` // 手机号
	Status int8 `db:"status" json:"status"` // 0正常  1注销
	Code string `db:"code" json:"code"` // 验证码
	CodeTime int64 `db:"code_time" json:"code_time"` // 验证码时效
}

type UserPosition struct {
	ID int64 `db:"id" json:"id"` // 主键
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	CompanyID int64 `db:"company_id" json:"company_id"` // 用户默认登录的企业
	UserID int64 `db:"user_id" json:"user_id"` // 用户ID
	PersonID int64 `db:"person_id" json:"person_id"` // 人事ID
	AccessToken string `db:"access_token" json:"access_token"` // 访问令牌
	ExpireIn int64 `db:"expire_in" json:"expire_in"` // 访问时效
}

