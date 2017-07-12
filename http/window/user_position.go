package window

import (
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
)
type UserPositionResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    []UserPositionListItem   `json:"data,omitempty"`
}
type UserPositionPaginator struct {
	Cond        *orm.Condition
	Cols        []string
	OrderBy     []string
	List        []user.UserPosition
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
type UserPositionListPaginator struct {
	Cond        int64  //user_id
	Cols        []string
	OrderBy     []string
	List        []UserPositionListItem
	PageSize    int //每页数量
	Total       int64 //总数
}
type UserPositionListItem struct {
	Id          int64 `db:"id" json:"id"`                  //id
	UserId      int64 `db:"user_id" json:"user_id"`        //userid
	CompanyName string `db:"name" json:"company_name"`      //公司名称
	CompanyShortName string `db:"short_name" json:"company_short_name"`      //公司名称
	CompanyId   int64  `db:"company_id" json:"company_id"` //公司id
	Type        int8 `db:"type" json:"type"`               //身份类型 1：孩子 2家长 3教师 4运营商 5服务商 6 管理员
	PersonId    int64 `db:"person_id" json:"person_id"`    // 人事ID
}
