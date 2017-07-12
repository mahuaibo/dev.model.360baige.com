package window

import (
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/company"
)

type CompanyPaginator struct {
	Cond        *orm.Condition
	Cols        []string
	OrderBy     []string
	List        []company.Company
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}
