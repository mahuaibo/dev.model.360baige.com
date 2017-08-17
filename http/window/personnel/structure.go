package personnel

type ListOfStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfStructure   `json:"data,omitempty"`
}

type ListOfStructure struct {
	List []StructureData `json:"list"`
}

type StructureData struct {
	Id       int64          `json:"id"`
	Label    string          `json:"label"`
	Type     int8           `json:"type"`
	Children []StructureData `json:"children"`
}

type Structure struct {
	Id         int64 `db:"id" json:"id"`                    // 主键
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"`  // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`    // 所有者ID
	ParentId   int64 `db:"parent_id" json:"parent_id"`      // parent_id
	Name       string `db:"name" json:"name"`               // 名称
	Type       int8 `db:"type" json:"type"`                 // 类型 1.班级
	Status     int8 `db:"status" json:"status"`             // 状态 0.下线  1.上线
}

type AddStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddStructure   `json:"data,omitempty"`
}

type AddStructure struct {
	Count int64 `json:"count"`
}

type ModifyStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ModifyStructure   `json:"data,omitempty"`
}

type ModifyStructure struct {
	Count int64 `json:"count"`
}

type DeleteStrectureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    DeleteStrecture   `json:"data,omitempty"`
}

type DeleteStrecture struct {
	Count int64
}

type GetStrectureIdsResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    GetStrectureIds   `json:"data,omitempty"`
}

type GetStrectureIds struct {
	Ids []int64        `json:"ids"`
}
