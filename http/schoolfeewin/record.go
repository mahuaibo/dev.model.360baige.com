package schoolfeewin

type ListOfRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    ListOfRecord   `json:"data,omitempty"`
}

type ClassListOfRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    []map[string]string   `json:"data,omitempty"`
}

type ListOfRecord struct {
	List []Record
}

type Record struct {
	Id         int64 `json:"id"`          // 主键
	CreateTime string `json:"create_time"` // 创建时间
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

type DeleteRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    DeleteRecord   `json:"data,omitempty"`
}

type DeleteRecord struct {
	Count int64
}

type UploadRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    UploadRecord   `json:"data,omitempty"`
}

type UploadRecord struct {
	Count int64
}

type DownloadRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    DownloadRecord   `json:"data,omitempty"`
}

type DownloadRecord struct {
	List []Record
}

type AddRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    AddRecord   `json:"data,omitempty"`
}

type AddRecord struct {
	Id int64 `json:"id"`
}

type DetailRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    DetailRecord   `json:"data,omitempty"`
}

type DetailRecord struct {
	Data Record
}

type ModifyRecordResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    ModifyRecord   `json:"data,omitempty"`
}

type ModifyRecord struct {
	Count int64 `json:"count"`
}
