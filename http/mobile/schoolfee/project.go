package schoolfee

type ListOfNoLimitProjectResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    ListOfNoLimitProject   `json:"data,omitempty"`
}

type ListOfNoLimitProject struct {
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

type SearchProjectInfoResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    ListOfRecordProject   `json:"data,omitempty"`
}

type ListOfRecordProject struct {
	List []RecordProject
}

type Record struct {
	Id         int64 `json:"id"`          // 主键
	CreateTime int64 `json:"create_time"` // 创建时间
	UpdateTime int64 `json:"update_time"` // 更新时间
	CompanyId  int64 `json:"company_id"`  // 所属公司ID
	ProjectId  int64 `json:"project_id"`  // 项目ID
	Name       string `json:"name"`       // 姓名
	ClassName  string `json:"class_name"` // 班级名称
	IdCard     string `json:"id_card"`    // 身份证号
	Num        string `json:"num"`        // 编号
	Phone      string `json:"phone"`      // 联系电话
	Status     int8 `json:"status"`       // 状态 -1删除
	Price      float64 `json:"price"`     // 应缴费用
	IsFee      int8 `json:"is_fee"`       // 是否缴费
	FeeTime    int64 `json:"fee_time"`    // 缴费时间（毫秒）
	Desc       string `json:"desc"`       // 备注
}

type RecordProject struct {
	Id         int64 `json:"id"`          // 主键
	CreateTime int64 `json:"create_time"` // 创建时间
	UpdateTime int64 `json:"update_time"` // 更新时间
	CompanyId  int64 `json:"company_id"`  // 所属公司ID
	ProjectId  int64 `json:"project_id"`  // 项目ID
	Name       string `json:"name"`       // 姓名
	ClassName  string `json:"class_name"` // 班级名称
	IdCard     string `json:"id_card"`    // 身份证号
	Num        string `json:"num"`        // 编号
	Phone      string `json:"phone"`      // 联系电话
	Status     int8 `json:"status"`       // 状态 -1删除
	Price      float64 `json:"price"`     // 应缴费用
	IsFee      int8 `json:"is_fee"`       // 是否缴费
	FeeTime    int64 `json:"fee_time"`    // 缴费时间（毫秒）
	Desc       string `json:"desc"`       // 备注
	Project    Project
}
