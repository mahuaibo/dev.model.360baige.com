package center

type LoggerAddResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    LoggerAdd   `json:"data,omitempty"`
}

type LoggerAdd struct {
	Id int64 `db:"id" json:"id"` //
}

type LoggerListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    LoggerList   `json:"data,omitempty"`
}

type LoggerList struct {
	OrderBy     []string `json:"orderBy"`
	List        []LoggerValue `json:"list"`
	Total       int64 `json:"total"`       //总数
	PageSize    int64 `json:"pageSize"`    //每页数量
	Current     int64 `json:"current"`     //当前页码
	CurrentSize int64 `json:"currentSize"` //当前页数量
	Status      int8  `json:"status"`      //状态：0 启用 1停用
	Name        string `json:"name"`       //搜索名称
}

type LoggerValue struct {
	CreateTime string `json:"createTime"` // 创建时间
	Content    string `json:"content"`    // 内容(修改前后变化)
	Remark     string `json:"remark"`     // 描述
	Type       string `json:"type"`       // 类别（1增、2删、3改、4查）
}
