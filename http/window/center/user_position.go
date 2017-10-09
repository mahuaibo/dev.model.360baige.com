package center

type UserPositionResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    []UserPositionListItem   `json:"data,omitempty"`
}

type UserPositionListItem struct {
	UserPositionId   int64 `json:"userPositionId"`    //id
	UserPositionName string `json:"userPositionName"` //身份类型 11：孩子 12家长 13教师 24运营商 35服务商 46 管理员
	CompanyLogo      string `json:"companyLogo"`      //公司名称
	CompanyName      string `json:"companyName"`      //公司名称
}
type UserPositionTokenResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserPositionToken   `json:"data,omitempty"`
}
type UserPositionToken struct {
	AccessToken string `json:"accessToken"` // 访问令牌
	ExpireIn    int64 `json:"expireIn"`     // 访问时效
}
type ChangeAccessTokenResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    UserPositionToken   `json:"data,omitempty"`
}
