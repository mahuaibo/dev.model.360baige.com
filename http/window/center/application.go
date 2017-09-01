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
	Price              string `json:"price"`  // 价格 tpl
	PayCycle           int8 `json:"payCycle"` // 0无1月2季3半年4年tpl
	SubscriptionStatus int8 `json:"status"`   // 状态 0 未订阅 1 已订阅
}
type ApplicationTalDetailResponse struct {
	Code    string `json:"code"`
	Message string        `json:"message"`
	Data    ApplicationTalDetail   `json:"data,omitempty"`
}

type ApplicationTalDetail struct {
	Name        string `json:"name"`             // 名称
	Image       string `json:"image"`            // 图片链接
	Desc        string `json:"desc"`             // 产品说明
	PriceDesc   string `json:"priceDesc"`        // 费用说明
	PayType     string `json:"payType"`          // 0:限免 1:永久免费 2:1次性收费 3:周期收费
	Price       float64 `json:"price"`           // 价格
	PayCycle    string `json:"payCycle"`         // 0无1月2季3半年4年
	Application Application `json:"application"` // 0无1月2季3半年4年
}

type Application struct {
	Id        int64 `json:"id"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
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

func GetPayTypeName(ptype int8) string {
	// 0:限免 1:永久免费 2:1次性收费 3:周期收费tpl
	var rPtype string
	switch ptype {
	case 0:
		rPtype = "限免"
	case 1:
		rPtype = "永久免费"
	case 2:
		rPtype = "1次性收费"
	case 3:
		rPtype = "周期收费"
	}
	return rPtype
}

func GetPayCycleName(ptype int8) string {
	// 0无1月2季3半年4年tpl
	var rPtype string
	switch ptype {
	case 0:
		rPtype = "无"
	case 1:
		rPtype = "月"
	case 2:
		rPtype = "季"
	case 3:
		rPtype = "半年"
	case 4:
		rPtype = "年"
	}
	return rPtype
}
