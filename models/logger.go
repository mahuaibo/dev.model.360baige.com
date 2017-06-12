package models

type Logger struct {
	ID         int64 `db:"id" json:"id"`                   // 主键 id
	CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间
	CompanyID  int64 `db:"company_id" json:"company_id"`   // 公司id
	OwnerID    int64 `db:"owner_id" json:"owner_id"`       // 操作者ID
	Remark     string `db:"remark" json:"remark"`          // 描述
	Content    string `db:"content" json:"content"`        // 内容(修改前后变化)
	Type       int `db:"type" json:"type"`                 // 类别（增、删、改、查）
}
