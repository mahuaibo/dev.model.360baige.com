package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = false
	orm.RegisterDataBase("user", "mysql", "demo2015:baige.2016@tcp(182.92.163.192:3306)/db_passport?charset=utf8", 30)
	orm.RegisterModel(&User{}, &UserPosition{})
}

type User struct {
	Id          int64 `db:"id" json:"id"`                      // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`    // 创建时间（毫秒）
	UpdateTime  int64 `db:"update_time" json:"update_time"`    // 修改时间（毫秒）
	Username    string `db:"username" json:"username"`         // 用户名
	Password    string `db:"password" json:"password"`         // 密码
	Email       string `db:"email" json:"email"`               // 邮件
	Phone       string `db:"phone" json:"phone"`               // 手机号
	Status      int8 `db:"status" json:"status"`               // 0正常  1注销
	Code        string `db:"code" json:"code"`                 // 验证码
	CodeTime    int64 `db:"code_time" json:"code_time"`        // 验证码时效
	AccessToken string `db:"access_token" json:"access_token"` // 访问令牌
	ExpireIn    int64 `db:"expire_in" json:"expire_in"`        // 访问时效
}

type UserPosition struct {
	Id          int64 `db:"id" json:"id"`                      // 主键
	CreateTime  int64 `db:"create_time" json:"create_time"`    // 创建时间
	UpdateTime  int64 `db:"update_time" json:"update_time"`    // 更新时间
	CompanyId   int64 `db:"company_id" json:"company_id"`      // 用户默认登录的企业
	UserId      int64 `db:"user_id" json:"user_id"`            // 用户Id
	PersonId    int64 `db:"person_id" json:"person_id"`        // 人事Id
	AccessToken string `db:"access_token" json:"access_token"` // 访问令牌
	ExpireIn    int64 `db:"expire_in" json:"expire_in"`        // 访问时效
}

/**
 * User
 */
// 新增
func (*User) AddUser(args *User, reply *User) error {
	o := orm.NewOrm()
	o.Using("user")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.Username = args.Username
		reply.Password = args.Password
		reply.Email = args.Email
		reply.Phone = args.Phone
		reply.Code = args.Code
		reply.CodeTime = args.CodeTime
		reply.Status = args.Status
	}
	return err
}

// 查询 by Id
func (*User) FindUserById(args *User, reply *User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 查询 by Username
func (*User) FindUserByUsername(args *User, reply *User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Username = args.Username
	err := o.Read(reply, "Username")
	return err
}

// 查询 by Id
func (*User) UpdateUserById(args *User, reply *User) error {
	o := orm.NewOrm()
	o.Using("user")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}

// 删除
func (*User) Delete(args *User, reply *User) error {
	o := orm.NewOrm()
	o.Using("user")
	num, err := o.Delete(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		} else {
			return errors.New("没有找到相关信息")
		}
	}
	return err
}

/**
 * UserPosition
 */
//
func (*UserPosition) FindUserPositionById(args *UserPosition, reply *UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

func (*UserPosition) FindUserPositionByUserId(args *UserPosition, reply *UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.UserId = args.UserId
	err := o.Read(reply, "UserId")
	return err
}
