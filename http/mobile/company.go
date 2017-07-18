package mobile

import (
	"dev.model.360baige.com/models/company"
)
type CompanyDetailResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    CompanyDetail   `json:"data,omitempty"`
}
type CompanyDetail struct {
	Logo string `db:"logo" json:"logo"` // 企业LOGO
	Name string `db:"name" json:"name"` // 企业全名
	ShortName string `db:"short_name" json:"short_name"` // 企业简称（微信端显示使用）
	ProvinceId int64 `db:"province_id" json:"province_id"` // 省
	CityId int64 `db:"city_id" json:"city_id"` // 市
	DistrictId int64 `db:"district_id" json:"district_id"` // 县
	Address string `db:"address" json:"address"` // 企业地址
	PositionX float64 `db:"position_x" json:"position_x"` // x
	PositionY float64 `db:"position_y" json:"position_y"` // y
	Remark string `db:"remark" json:"remark"` // 备注
	Brief string `db:"brief" json:"brief"` // 简介
	Status int8 `db:"status" json:"status"` // 1 启用 0 停用
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
