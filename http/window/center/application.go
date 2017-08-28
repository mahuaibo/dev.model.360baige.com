package center

type ApplicationListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ApplicationList   `json:"data,omitempty"`
}
type ApplicationList struct {
	OrderBy     []string `json:"orderBy"`
	List        []ApplicationValue `json:"list"`
	Total       int64  `json:"total"`       //总数
	PageSize    int64  `json:"pageSize"`    //每页数量
	Current     int64  `json:"current"`     //当前页码
	CurrentSize int64  `json:"currentSize"` //当前页数量
	Status      int8   `json:"status"`      //状态：0 启用 1停用
	Name        string `json:"name"`        //搜索名称
}

type ApplicationTplListResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ApplicationTplList   `json:"data,omitempty"`
}
type ApplicationTplList struct {
	OrderBy     []string  `json:"orderBy"`
	List        []ApplicationTplValue `json:"list"`
	Total       int64  `json:"total"`       //总数
	PageSize    int64  `json:"pageSize"`    //每页数量
	Current     int64  `json:"current"`     //当前页码
	CurrentSize int64  `json:"currentSize"` //当前页数量
	Status      int8   `json:"status"`      //状态：0 启用 1停用
	Name        string `json:"name"`        //搜索名称
}

type ApplicationValue struct {
	Id         int64 `json:"id"`          // 主键自动增长id apid
	CreateTime string `json:"createTime"` // 创建时间ap
	Name       string `json:"name"`       // 名称ap is null tpl
	Image      string `json:"image"`      // 图片链接ap is null tpl
	Status     int8 `json:"status"`       // 状态 0 启用 1停用
	Site       string `json:"site"`       // 访问链接tpl
}
type ApplicationTplValue struct {
	Id                 int64 `json:"id"`      // 主键自动增长id aptplid
	Name               string `json:"name"`   // 名称tpl
	Image              string `json:"image"`  // 图片链接tpl
	Desc               string `json:"desc"`   // 简介tpl
	Price              float64 `json:"price"` // 价格 tpl
	PayCycle           int8 `json:"payCycle"` // 0无1月2季3半年4年tpl
	SubscriptionStatus int8 `json:"status"`   // 状态 0 未订阅 1 已订阅
}
type ApplicationDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ApplicationDetail   `json:"data,omitempty"`
}

type ApplicationDetail struct {
	CreateTime  string `json:"createTime"`  // 创建时间ap
	Name        string `json:"name"`        // 名称ap is null tpl
	Image       string `json:"image"`       // 图片链接ap is null tpl
	Desc        string `json:"desc"`        // 简介tpl
	Price       float64 `json:"price"`      // 价格 tpl
	PayType     string `json:"payType"`     // 0:限免 1:永久免费 2:1次性收费 3:周期收费tpl
	PayCycle    string `json:"payCycle"`    // 0无1月2季3半年4年tpl
	Site        string`json:"site"`         // 访问链接
	CompanyName string `json:"companyName"` // 开发公司ID tpl
	UserName    string `json:"userName"`    // 开发者ID tpl
}
type ModifyApplicationStatusResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
type ModifyApplicationTplStatusResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ApplicationTplStatus   `json:"data,omitempty"`
}

type ApplicationTplStatus struct {
	AppId            int64 `json:"appId"`            //ap id
	ApplicationTplId int64 `json:"applicationTplId"` // 应用ID
}
