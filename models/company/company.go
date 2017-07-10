package company

type Company struct {
	 Id int64 `db:"id" json:"id"` // 主键自动增长id
	 CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	 UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间
	 Type int8 `db:"type" json:"type"` // 类型：10：幼儿园 11：小学 12:中学 13：高中 20：运营  30：代理商 31：经销商
	 Level int8 `db:"level" json:"level"` // 地域级别：10：省级/直辖市 20：地市级/省会 30：区县级/城区
	 Logo string `db:"logo" json:"logo"` // 企业LOGO
	 Name string `db:"name" json:"name"` // 企业全名
	 ShortName string `db:"short_name" json:"short_name"` // 企业简称（微信端显示使用）
	 SubDomain string `db:"sub_domain" json:"sub_domain"` // 企业子域名（自动分配，微信设置使用）
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
