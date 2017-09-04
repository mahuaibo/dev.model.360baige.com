package account

// 账户初始余额

const AccountBalanceInit int64  = 0

// 账户类型
const (
	AccountTypeMoney    int8 = 1 // 现金
	AccountTypeIntegral int8 = 2 // 积分
)

// 账户状态
const (
	AccountStatusCancel = -1 // 正常
	AccountStatusNormal = 0  // 正常
)
