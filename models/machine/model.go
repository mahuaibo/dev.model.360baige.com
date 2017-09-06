package machine

type Machine struct {
	Id           int64    `db:"id" json:"id"`                       // 主键
	CreateTime   int64    `db:"create_time" json:"createTime"`      // 创建时间（毫秒）
	UpdateTime   int64    `db:"update_time" json:"updateTime"`      // 更新时间（毫秒）
	CompanyId    int64    `db:"company_id" json:"companyId"`        // 所有者ID
	Name         string    `db:"name" json:"name"`                  // 设备名称
	MachineNo    string    `db:"machine_no" json:"machineNo"`       // 设备编号
	PhysicsNo    string    `db:"physics_no" json:"physicsNo"`       // 物理编号
	BatchNo      int    `db:"batch_no" json:"batchNo"`              // 批次号
	Type         int    `db:"type" json:"type"`                     // 设备类型
	AccessSecret string    `db:"access_secret" json:"accessSecret"` // 访问密钥
	AccessToken  string    `db:"access_token" json:"accessToken"`   // 访问令牌
	ExpiresIn    int64    `db:"expires_in" json:"expiresIn"`        // 访问时效（毫秒）
	Status       int    `db:"status" json:"status"`                 // 状态 -1注销、0未出库、1已出库、2未绑定、3已绑定、4挂失
}
