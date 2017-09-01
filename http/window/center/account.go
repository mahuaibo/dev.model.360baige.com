package center

type AccountStatisticsResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountStatistics   `json:"data,omitempty"`
}

type AccountStatistics struct {
	Balance    string `json:"balance"`    // 余额
	OutAccount string `json:"outAccount"` // 总出账
	InAccount  string `json:"inAccount"`  // 总入账
}
