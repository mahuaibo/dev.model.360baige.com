package company

// 初始游客公司ID
const CompanyInit = 0

// 公司状态 -1撤销 0正常 1暂停
const (
	CompanyStatusCancel int8 = -1 // 撤销
	CompanyStatusNormal int8 = 0  // 正常
	CompanyStatusPause  int8 = 1  // 暂停
)
