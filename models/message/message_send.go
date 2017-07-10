package message

type MessageSend struct {
	 Id int64 `db:"id" json:"id"` // 主键
	 CreateTime int64 `db:"create_time" json:"create_time"` // 创建时间（毫秒）
	 UpdateTime int64 `db:"update_time" json:"update_time"` // 更新时间（毫秒）
	 Type int8 `db:"type" json:"type"` // 类型
	 SendTime int64 `db:"send_time" json:"send_time"` // 发送时间（毫秒）
	 SendCompanyId int64 `db:"send_company_id" json:"send_company_id"` // 企业ID
	 SendUserId int64 `db:"send_user_id" json:"send_user_id"` // 企业ID
	 SendUserPositionId int64 `db:"send_user_position_id" json:"send_user_position_id"` // 企业ID
	 SendUserPositionType int8 `db:"send_user_position_type" json:"send_user_position_type"` // 企业ID
	 SendDest string `db:"send_dest" json:"send_dest"` // 发送目标
	 ReceiveTime int64 `db:"receive_time" json:"receive_time"` // 发送时间（毫秒）
	 ReceiveCompanyId int64 `db:"receive_company_id" json:"receive_company_id"` // 企业ID
	 ReceiveUserId int64 `db:"receive_user_id" json:"receive_user_id"` // 企业ID
	 ReceiveUserPositionId int64 `db:"receive_user_position_id" json:"receive_user_position_id"` // 企业ID
	 ReceiveUserPositionType int8 `db:"receive_user_position_type" json:"receive_user_position_type"` // 企业ID
	 ReceiveDest string `db:"receive_dest" json:"receive_dest"` // 发送目标
	 MessageTempId int64 `db:"message_temp_id" json:"message_temp_id"` // 企业ID
	 SplitContent string `db:"split_content" json:"split_content"` // 内容
	 Total int `db:"total" json:"total"` // 发送条数
	 Status int8 `db:"status" json:"status"` // 发送状态
	 Pos int `db:"pos" json:"pos"` // 第几条
	
}
