package window

type AccountStatisticsResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    AccountStatistics   `json:"data,omitempty"`
}

type AccountStatistics struct {
	Balance        float64 `json:"balance"`        // 余额
	TotalEntry     float64 `json:"total_entry"`     // 总入账
	TotalDischarge float64 `json:"total_discharge"` // 总出账
	MonthIncome    float64 `json:"month_income"`    // 月收入
	MonthPay       float64 `json:"month_pay"`       // 月支出
}
