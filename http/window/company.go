package window

import (
	"dev.model.360baige.com/models/company"
)

type CompanyDetailResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    CompanyDetail   `json:"data,omitempty"`
}
type CompanyDetail struct {
	Id         int64 `json:"id"`           // 企业Id
	Logo       string `json:"logo"`        // 企业LOGO
	Name       string `json:"name"`        // 企业全名
	ShortName  string `json:"short_name"`  // 企业简称（微信端显示使用）
	ProvinceId int64 `json:"province_id"`  // 省
	CityId     int64 `json:"city_id"`      // 市
	DistrictId int64 `json:"district_id"`  // 县
	Address    string `json:"address"`     // 企业地址
	PositionX  float64 `json:"position_x"` // x
	PositionY  float64 `json:"position_y"` // y
	Remark     string `json:"remark"`      // 备注
	Brief      string `json:"brief"`       // 简介
	Status     int8 `json:"status"`        // 1 启用 0 停用
}
type CompanyModifyResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    CompanyDetail   `json:"data,omitempty"`
}

type CompanyPaginator struct {
	Cond        []CondValue
	Cols        []string
	OrderBy     []string
	List        []company.Company
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
