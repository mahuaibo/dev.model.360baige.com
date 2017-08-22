package message

type MessageTemp struct {
    Id    int64    `db:"id" json:"id"`    // 主键
    CreateTime    int64    `db:"create_time" json:"createTime"`    // 创建时间（毫秒）
    UpdateTime    int64    `db:"update_time" json:"updateTime"`    // 更新时间（毫秒）
    Type    int8    `db:"type" json:"type"`    // 类型
    SendTime    int64    `db:"send_time" json:"sendTime"`    // 发送时间（毫秒）
    SendCompanyId    int64    `db:"send_company_id" json:"sendCompanyId"`    // 企业ID
    SendUserId    int64    `db:"send_user_id" json:"sendUserId"`    // 企业ID
    SendUserPositionId    int64    `db:"send_user_position_id" json:"sendUserPositionId"`    // 企业ID
    SendUserPositionType    int8    `db:"send_user_position_type" json:"sendUserPositionType"`    // 企业ID
    SendDest    string    `db:"send_dest" json:"sendDest"`    // 发送目标
    ReceiveTime    int64    `db:"receive_time" json:"receiveTime"`    // 发送时间（毫秒）
    ReceiveCompanyId    int64    `db:"receive_company_id" json:"receiveCompanyId"`    // 企业ID
    ReceiveUserId    int64    `db:"receive_user_id" json:"receiveUserId"`    // 企业ID
    ReceiveUserPositionId    int64    `db:"receive_user_position_id" json:"receiveUserPositionId"`    // 企业ID
    ReceiveUserPositionType    int8    `db:"receive_user_position_type" json:"receiveUserPositionType"`    // 企业ID
    ReceiveDest    string    `db:"receive_dest" json:"receiveDest"`    // 发送目标
    Content    string    `db:"content" json:"content"`    // 发送内容
    Total    int    `db:"total" json:"total"`    // 发送条数
    Status    int8    `db:"status" json:"status"`    // 发送状态
    
}
