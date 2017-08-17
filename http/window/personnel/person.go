package personnel

type ListOfPersonResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ListOfPerson   `json:"data,omitempty"`
}

type ListOfPerson struct {
	List []Person `json:"list"`
}

type Person struct {
	Id         int64 `db:"id" json:"id"`                    // 主键
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	UpdateTime int64 `db:"update_time" json:"update_time"`  // 更新时间（毫秒）
	CompanyId  int64 `db:"company_id" json:"company_id"`    // 所有者ID
	Code       string`db:"code" json:"code"`                // 编号
	Name       string `db:"name" json:"name"`               // 名称
	Sex        string `db:"sex" json:"sex"`                 // 性别
	Birthday   int64 `db:"birthday" json:"birthday"`        // 生日
	Type       int8 `db:"type" json:"type"`                 // 类型 1.教师   2.家长   3.学生
	Phone      string `db:"phone" json:"phone"`             // 联系方式
	Contact    string `db:"contact" json:"contact"`         // 联系人
	Status     int8 `db:"status" json:"status"`             // 状态 1.正常   2.异常  3.学生
}

type AddPersonResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AddPerson   `json:"data,omitempty"`
}

type AddPerson struct {
	Id int64 `json:"id"`
}

type ModifyPersonResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ModifyPerson   `json:"data,omitempty"`
}

type ModifyPerson struct {
	Count int64 `json:"count"`
}

type DeletePersonResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    DeletePerson   `json:"data,omitempty"`
}

type DeletePerson struct {
	Count int64
}

type UploadPersonResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UploadPerson   `json:"data,omitempty"`
}

type UploadPerson struct {
	Count int64
}

type DownloadPersonResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    DownloadPerson   `json:"data,omitempty"`
}

type DownloadPerson struct {
	List []Person
}