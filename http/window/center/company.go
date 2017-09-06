package center

type CompanyDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    CompanyDetail   `json:"data,omitempty"`
}
type CompanyDetail struct {
	Id         int64 `json:"id"`          // 企业Id
	Logo       string `json:"logo"`       // 企业LOGO
	Name       string `json:"name"`       // 企业全名
	ShortName  string `json:"shortName"`  // 企业简称（微信端显示使用）
	ProvinceId int64 `json:"provinceId"`  // 省
	CityId     int64 `json:"cityId"`      // 市
	DistrictId int64 `json:"districtId"`  // 县
	Address    string `json:"address"`    // 企业地址
	PositionX  float64 `json:"positionX"` // x
	PositionY  float64 `json:"positionY"` // y
	Remark     string `json:"remark"`     // 备注
	Brief      string `json:"brief"`      // 简介
	Status     int8 `json:"status"`       // 1 启用 0 停用
}
type CompanyModifyResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    CompanyDetail   `json:"data,omitempty"`
}

type UploadLogoResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    string   `json:"head"`
}