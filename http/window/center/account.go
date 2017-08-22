package center

type AccountStatisticsResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    AccountStatistics   `json:"data,omitempty"`
}

type AccountStatistics struct {
	Balance        float64 `json:"balance"`        // 余额
	TotalEntry     float64 `json:"totalEntry"`     // 总入账
	TotalDischarge float64 `json:"totalDischarge"` // 总出账
	MonthIncome    float64 `json:"monthIncome"`    // 月收入
	MonthPay       float64 `json:"monthPay"`       // 月支出
}
