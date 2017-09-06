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
	Children []StructureData `json:"children"`
}

type Structure struct {
	Id         int64 `json:"id"`          // 主键
	CreateTime string `json:"createTime"` // 创建时间（毫秒）
	UpdateTime int64 `json:"updateTime"`  // 更新时间（毫秒）
	CompanyId  int64 `json:"companyId"`   // 所有者ID
	ParentId   int64 `json:"parentId"`    // parent_id
	Name       string `json:"name"`       // 名称
	Type       int `json:"type"`         // 类型 1.班级
	Status     int `json:"status"`       // 状态 0.下线  1.上线
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
