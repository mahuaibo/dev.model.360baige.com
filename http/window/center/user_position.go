package center

import (
	"dev.model.360baige.com/models/user"
)

type UserPositionResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    []UserPositionListItem   `json:"data,omitempty"`
}

type UserPositionPaginator struct {
	Cond        []CondValue
			  //Cond        orm.Condition
			  //Cond        []byte
			  //Cond        string
	Cols        []string
	OrderBy     []string
	List        []user.UserPosition
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
type UserPositionListPaginator struct {
	Cond     int64 //user_id
	Cols     []string
	OrderBy  []string
	List     []UserPositionListItem
	PageSize int   //每页数量
	Total    int64 //总数
}
type UserPositionListItem struct {
	Id               int64 `json:"id"`                  //id
	CompanyName      string `json:"company_name"`       //公司名称
	CompanyShortName string `json:"company_short_name"` //公司名称
	CompanyId        int64  `json:"company_id"`         //公司id
	Type             int8 `json:"type"`                 //身份类型 11：孩子 12家长 13教师 24运营商 35服务商 46 管理员
	PersonId         int64 `json:"person_id"`           // 人事ID
	CompanyStatus    int8 `json:"status"`               // 1 启用 0 停用
}
type UserPositionTokenResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    UserPositionToken   `json:"data,omitempty"`
}
type UserPositionToken struct {
	AccessToken string `json:"access_token"` // 访问令牌
	ExpireIn    int64 `json:"expire_in"`     // 访问时效
}
