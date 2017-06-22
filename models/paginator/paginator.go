package paginator

import "github.com/astaxie/beego/orm"

type Paginator struct {
	//Type      int             //分页方式 按页码:1 按ID:2
	Total       int64           //总数
	PageSize    int             //每页数量
	Current     int             //当前页码
	MarkID      int64           //最大ID 或 最小ID
	Direction   int             //翻页方向 上:-1 平:0 下:1
	Sord        string          //排序字段
	Filters     string          //过滤条件
	List        []orm.Params    //内容
}

//exact / iexact 等于
//contains / icontains 包含
//gt / gte 大于 / 大于等于
//lt / lte 小于 / 小于等于
//startswith / istartswith 以…起始
//endswith / iendswith 以…结束
//in
//isnull
