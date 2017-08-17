package machine

type Machine struct {
	Id           int64 `db:"id" json:"id"`                        // 主键
	CreateTime   int64 `db:"create_time" json:"create_time"`      // 创建时间（毫秒）
	UpdateTime   int64 `db:"update_time" json:"update_time"`      // 更新时间（毫秒）
	CompanyId    int64 `db:"company_id" json:"company_id"`        // 所有者ID
	Name         string `db:"name" json:"name"`                   // 设备名称
	MachineNo    string `db:"machine_no" json:"machine_no"`       // 设备编号
	PhysicsNo    string `db:"physics_no" json:"physics_no"`       // 物理编号
	BatchNo      int8 `db:"batch_no" json:"batch_no"`             // 批次号
	Type         int8 `db:"type" json:"type"`                     // 设备类型
	AccessSecret string `db:"access_secret" json:"access_secret"` // 访问密钥
	AccessToken  string `db:"access_token" json:"access_token"`   // 访问令牌
	ExpiresIn    int64 `db:"expires_in" json:"expires_in"`        // 访问时效（毫秒）
	Status       int8 `db:"status" json:"status"`                 // 状态 -1注销、0未出库、1已出库、2未绑定、3已绑定、4挂失
}
