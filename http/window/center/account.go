package center

type AccountStatisticsResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountStatistics   `json:"data,omitempty"`
}

type AccountStatistics struct {
	Balance    int64 `json:"balance"`    // 余额
	OutAccount int64 `json:"outAccount"` // 总出账
	InAccount  int64 `json:"inAccount"`  // 总入账
}
