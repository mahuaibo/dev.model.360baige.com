package personnel

type ListOfPersonStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfPersonStructure   `json:"data,omitempty"`
}

type ListOfPersonStructure struct {
	List []PersonStructure
}

type PersonStructure struct {
	Id          int64 `json:"id"`          // 主键
	CreateTime  int64 `json:"createTime"`  // 创建时间（毫秒）
	UpdateTime  int64 `json:"updateTime"`  // 更新时间（毫秒）
	CompanyId   int64 `json:"companyId"`   // 所有者ID
	PersonId    int64 `json:"personId"`    // parent_id 人员ID
	StructureId int64 `json:"structureId"` // structure_id 结构ID
	Type        int `json:"type"`         // 类型
	Status      int `json:"status"`       // 状态  1.正常
}

type AddPersonStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddPersonStructure   `json:"data,omitempty"`
}

type AddPersonStructure struct {
	Id int64 `json:"id"`
}

type ModifyPersonStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ModifyPersonStructure   `json:"data,omitempty"`
}

type DetailPersonStructureResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    PersonStructure   `json:"data,omitempty"`
}

type ModifyPersonStructure struct {
	Count int64 `json:"count"`
}
