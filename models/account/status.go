package account

// 账户初始余额

const AccountBalanceInit int64 = 0

// 账户类型
const (
	AccountTypeMoney    int8 = 1 // 现金
	AccountTypeIntegral int8 = 2 // 积分
)

// 账户状态
const (
	AccountStatusCancel = -1 // 销毁
	AccountStatusNormal = 0  // 正常
)

// 账户记录状态
const (
	AccountItemStatusCancel = -1 // 销毁
	AccountItemStatusNormal = 0  // 正常
)

// 交易状态
const (
	TransactionStatusCancel = -1 // 销毁
	TransactionStatusNormal = 0  // 正常
)
