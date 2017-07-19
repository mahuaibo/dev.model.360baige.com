package window
import (
	"dev.model.360baige.com/models/logger"
)
type LoggerAddResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    LoggerAdd   `json:"data,omitempty"`
}
type LoggerAdd struct {
	Id  int64 `db:"id" json:"id"`   //ap id
}
type LoggerListResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    LoggerList   `json:"data,omitempty"`
}
type LoggerList struct {
	OrderBy     []string
	List        []LoggerValue
	Total       int64  //总数
	PageSize    int64  //每页数量
	Current     int64  //当前页码
	CurrentSize int64  //当前页数量
	Status      int8   //状态：0 启用 1停用
	Name        string //搜索名称
}
type LoggerListPaginator struct {
	Cond        []CondValue
	Cols        []string
	OrderBy     []string
	List        []logger.Logger
	Total       int64  //总数
	PageSize    int64  //每页数量
	Current     int64  //当前页码
	CurrentSize int64  //当前页数量
	Name        string //搜索名称
}
type LoggerValue struct {
	CreateTime string `db:"create_time" json:"create_time"` // 创建时间ap
	//CompanyName string `db:"company_name" json:"company_name"` // 企业ID
	//UserName string `db:"user_name" json:"user_name"` // 操作者ID
	//UserPosition string `db:"user_position_id" json:"user_position_id"` // 操作者ID
	Content string `db:"content" json:"content"` // 内容(修改前后变化)
	Remark string `db:"remark" json:"remark"` // 描述
	Type string `db:"type" json:"type"` // 类别（1增、2删、3改、4查）
}
