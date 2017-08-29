package center

type AccountStatisticsResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountStatistics   `json:"data,omitempty"`
}

type AccountStatistics struct {
	Balance    float64 `json:"balance"`    // 余额
	OutAccount float64 `json:"outAccount"` // 总出账
	InAccount  float64 `json:"inAccount"`  // 总入账
}
