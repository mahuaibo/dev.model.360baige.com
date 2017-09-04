package logger

type Logger struct {
	Id               int64    `db:"id" json:"id"`                              // 主键 id
	CreateTime       int64    `db:"create_time" json:"createTime"`             // 创建时间
	CompanyId        int64    `db:"company_id" json:"companyId"`               // 企业ID
	UserId           int64    `db:"user_id" json:"userId"`                     // 操作者ID
	UserPositionId   int64    `db:"user_position_id" json:"userPositionId"`    // 操作者ID
	UserPositionType int8    `db:"user_position_type" json:"userPositionType"` // 操作者type
	Content          string    `db:"content" json:"content"`                   // 内容(修改前后变化)
	Remark           string    `db:"remark" json:"remark"`                     // 描述
	Type             int8    `db:"type" json:"type"`                           // 类别（增、删、改、查）
}
