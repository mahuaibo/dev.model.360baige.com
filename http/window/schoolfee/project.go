package schoolfee

type ListOfProjectResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfProject   `json:"data,omitempty"`
}

type ListOfProject struct {
	List []Project `json:"list"`
}

type Project struct {
	Id         int64 `json:"id"`          // ID主键
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime int64 `json:"updateTime"`  // 更新时间
	CompanyId  int64 `json:"companyId"`   // 所属公司ID
	Name       string `json:"name"`       // 项目名称
	IsLimit    int8 `json:"isLimit"`      // 是否限制缴费
	Desc       string `json:"desc"`       // 描述
	Link       string `json:"link"`       // 描述链接
	Status     int8 `json:"status"`       // 状态 -1注销 0正常
}

type AddProjectResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddProject   `json:"data,omitempty"`
}

type AddProject struct {
	Id int64 `json:"id"`
}

type ModifyProjectResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ModifyProject   `json:"data,omitempty"`
}

type DetailProjectResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    Project   `json:"data,omitempty"`
}

type ModifyProject struct {
	Count int64 `json:"count"`
}
