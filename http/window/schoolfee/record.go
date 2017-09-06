package schoolfee

type ListOfRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfRecord   `json:"data,omitempty"`
}

type ClassListOfRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    []map[string]string   `json:"data,omitempty"`
}

type ListOfRecord struct {
	List     []Record `json:"list"`
	Total    int64 `json:"total"`
	PageSize int64 `json:"pageSize"`
	Current  int64 `json:"current"`
}

type Record struct {
	Id         int64 `json:"id"`          // 主键
	CreateTime string `json:"createTime"` // 创建时间
	UpdateTime int64 `json:"updateTime"`  // 更新时间
	CompanyId  int64 `json:"companyId"`   // 所属公司ID
	ProjectId  int64 `json:"projectId"`   // 项目ID
	Name       string `json:"name"`       // 姓名
	ClassName  string `json:"className"`  // 班级名称
	IdCard     string `json:"idCard"`     // 身份证号
	Num        string `json:"num"`        // 编号
	Phone      string `json:"phone"`      // 联系电话
	Status     int `json:"status"`       // 状态 -1删除
	Price      int64 `json:"price"`       // 应缴费用
	IsFee      int `json:"isFee"`        // 是否缴费
	FeeTime    string `json:"feeTime"`    // 缴费时间（毫秒）
	Desc       string `json:"desc"`       // 备注
}

type DeleteRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    DeleteRecord   `json:"data,omitempty"`
}

type DeleteRecord struct {
	Count int64 `json:"count"`
}

type UploadRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UploadRecord   `json:"data,omitempty"`
}

type UploadRecord struct {
	Count int64 `json:"count"`
}

type DownloadRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    DownloadRecord   `json:"data,omitempty"`
}

type DownloadRecord struct {
	List []Record `json:"list"`
}

type AddRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddRecord   `json:"data,omitempty"`
}

type AddRecord struct {
	Id int64 `json:"id"`
}

type DetailRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    DetailRecord   `json:"data,omitempty"`
}

type DetailRecord struct {
	Data Record `json:"data"`
}

type ModifyRecordResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ModifyRecord   `json:"data,omitempty"`
}

type ModifyRecord struct {
	Count int64 `json:"count"`
}
