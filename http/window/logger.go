package window

type LoggerPageResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    LoggerPageData   `json:"data,omitempty"`
}

type LoggerPageData struct {
	Cond        []CondValue
	Cols        []string
	OrderBy     []string
	List []Logger
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}

type LoggerResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    LoggerData   `json:"data,omitempty"`
}

type LoggerData struct {
	Detail Logger
}

type Logger struct {
	Id               int64 `json:"id"`                // 主键 id
	CreateTime       int64 `json:"create_time"`       // 创建时间
	CompanyId        int64 `json:"company_id"`        // 企业ID
	UserId           int64 `json:"user_id"`           // 操作者ID
	UserPositionId   int64 `json:"user_position_id"`  // 操作者ID
	UserPositionType int8 `json:"user_position_type"` // 类别（增、删、改、查）
	Content          string `json:"content"`          // 内容(修改前后变化)
	Remark           string `json:"remark"`           // 描述
	Type             int8 `json:"type"`               // 类别（增、删、改、查
}
