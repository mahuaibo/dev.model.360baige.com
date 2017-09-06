package account

// 账户初始余额

const AccountBalanceInit int64 = 0

// 账户类型
const (
	AccountTypeMoney    int = 1 // 现金
	AccountTypeIntegral int = 2 // 积分
)

// 账户状态
const (
	AccountStatusCancel int = -1 // 销毁
	AccountStatusNormal int = 0  // 正常
)

// 账户记录状态
const (
	AccountItemStatusCancel int = -1 // 销毁
	AccountItemStatusNormal int = 0  // 正常
)

// 交易状态
const (
	TransactionStatusCancel int = -1 // 销毁
	TransactionStatusNormal int = 0  // 正常
)
