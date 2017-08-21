package center

import (
	"dev.model.360baige.com/models/user"
)

type UserPositionResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    []UserPositionListItem   `json:"data,omitempty"`
}

type UserPositionPaginator struct {
	Cond        []CondValue
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
	Id          int64 `json:"user_position_id"`  //id
	Type        int8 `json:"user_position_type"` //身份类型 11：孩子 12家长 13教师 24运营商 35服务商 46 管理员
	CompanyLogo string `json:"company_logo"`     //公司名称
	CompanyName string `json:"company_name"`     //公司名称
}
type UserPositionTokenResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserPositionToken   `json:"data,omitempty"`
}
type UserPositionToken struct {
	AccessToken string `json:"access_token"` // 访问令牌
	ExpireIn    int64 `json:"expire_in"`     // 访问时效
}
