package city

type City struct {
	 Id int64 `db:"id" json:"id"` // 主键
	 CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	 UpdateTime int64 `db:"update_time" json:"update_time"` // 修改时间（毫秒）
	 Type string `db:"type" json:"type"` // 类型：1省、 2市、3县
	 ParentId int64 `db:"parent_id" json:"parent_id"` // 父级ID
	 Name string `db:"name" json:"name"` // 城市名称
	 Shortname string `db:"shortname" json:"shortname"` // 城市简称
	 Pinyin string `db:"pinyin" json:"pinyin"` // 拼音
	 Domain string `db:"domain" json:"domain"` // 域名
	 PositionX float64 `db:"position_x" json:"position_x"` // 经度(x) 赤道垂直
	 PositionY float64 `db:"position_y" json:"position_y"` // 纬度(y) 赤道平行
	 Pos int `db:"pos" json:"pos"` // 顺序
	
}
