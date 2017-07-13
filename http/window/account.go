package window

type AccountStatisticsResponse struct {
	Code    string `json:"code"`
	Messgae string        `json:"messgae"`
	Data    AccountStatistics   `json:"data,omitempty"`
}
type AccountStatistics struct {
	Balance        float64 `db:"balance" json:"balance"`        // 余额
	TotalEntry     float64 `db:"amount" json:"total_entry"`     // 总入账
	TotalDischarge float64 `db:"amount" json:"total_discharge"` // 总出账
	MonthIncome    float64 `db:"amount" json:"month_income"`        // 月收入
	MonthPay       float64 `db:"amount" json:"month_pay "`        // 月支出
}

