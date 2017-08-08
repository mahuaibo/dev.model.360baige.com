package schoolfeewin

type ListOfProjectResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    ListOfProject   `json:"data,omitempty"`
}

type ListOfProject struct {
	List []Project
}

type Project struct {
	Id         int64 `json:"id"`          // ID主键
	CreateTime int64 `json:"create_time"` // 创建时间
	UpdateTime int64 `json:"update_time"` // 更新时间
	CompanyId  int64 `json:"company_id"`  // 所属公司ID
	Name       string `json:"name"`       // 项目名称
	IsLimit    int8 `json:"is_limit"`     // 是否限制缴费
	Desc       string `json:"desc"`       // 描述
	Link       string `json:"link"`       // 描述链接
	Status     int8 `json:"status"`       // 状态 -1注销 0正常
}

type AddProjectResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    AddProject   `json:"data,omitempty"`
}

type AddProject struct {
	Id int64 `json:"id"`
}


type ModifyProjectResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    ModifyProject   `json:"data,omitempty"`
}

type ModifyProject struct {
	Count int64 `json:"count"`
}