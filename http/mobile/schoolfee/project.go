package schoolfee

type ListOfNoLimitProjectResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfNoLimitProject   `json:"data,omitempty"`
}

type ListOfNoLimitProject struct {
	List []Project `json:"list"`
}

type Project struct {
	Id         int64 `json:"id"`         // ID主键
	CreateTime int64 `json:"createTime"` // 创建时间
	UpdateTime int64 `json:"updateTime"` // 更新时间
	CompanyId  int64 `json:"companyId"`  // 所属公司ID
	Name       string `json:"name"`      // 项目名称
	IsLimit    int `json:"isLimit"`     // 是否限制缴费
	Desc       string `json:"desc"`      // 描述
	Link       string `json:"link"`      // 描述链接
	Status     int `json:"status"`      // 状态 -1注销 0正常
}

type SearchProjectInfoResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfRecordProject   `json:"data,omitempty"`
}
type ProjectDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    Project   `json:"data,omitempty"`
}
type RecordHistoryResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfRecordProject   `json:"data,omitempty"`
}
type AddRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddRecord   `json:"data,omitempty"`
}
type AddMultipleRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddMultipleRecord   `json:"data,omitempty"`
}
type AddMultipleRecord struct {
	Num int64 `json:"num"`
}
type AddRecord struct {
	Id int64 `json:"id"`
}

type ListOfRecordProject struct {
	List        []RecordProject `json:"list"`
	OrderBy     []string `json:"orderBy"`
	Total       int64  `json:"total"`       //总数
	PageSize    int64  `json:"pageSize"`    //每页数量
	Current     int64  `json:"current"`     //当前页码
	CurrentSize int64  `json:"currentSize"` //当前页数量
	SearchKey   string `json:"search"`      //搜索key
}

type Record struct {
	Id         int64 `json:"id"`         // 主键
	CreateTime int64 `json:"createTime"` // 创建时间
	UpdateTime int64 `json:"updateTime"` // 更新时间
	CompanyId  int64 `json:"companyId"`  // 所属公司ID
	ProjectId  int64 `json:"projectId"`  // 项目ID
	Name       string `json:"name"`      // 姓名
	ClassName  string `json:"className"` // 班级名称
	IdCard     string `json:"idCard"`    // 身份证号
	Num        string `json:"num"`       // 编号
	Phone      string `json:"phone"`     // 联系电话
	Status     int `json:"status"`      // 状态 -1删除
	Price      int64 `json:"price"`      // 应缴费用
	IsFee      int `json:"isFee"`       // 是否缴费
	FeeTime    int64 `json:"feeTime"`    // 缴费时间（毫秒）
	Desc       string `json:"desc"`      // 备注
}

type RecordProject struct {
	Id         int64 `json:"id"`         // 主键
	CreateTime int64 `json:"createTime"` // 创建时间
	UpdateTime int64 `json:"updateTime"` // 更新时间
	CompanyId  int64 `json:"companyId"`  // 所属公司ID
	ProjectId  int64 `json:"projectId"`  // 项目ID
	Name       string `json:"name"`      // 姓名
	ClassName  string `json:"className"` // 班级名称
	IdCard     string `json:"idCard"`    // 身份证号
	Num        string `json:"num"`       // 编号
	Phone      string `json:"phone"`     // 联系电话
	Status     int `json:"status"`      // 状态 -1删除
	Price      int64 `json:"price"`      // 应缴费用
	IsFee      int `json:"isFee"`       // 是否缴费
	FeeTime    int64 `json:"feeTime"`    // 缴费时间（毫秒）
	Desc       string `json:"desc"`      // 备注
	Project    Project  `json:"project"`
}
