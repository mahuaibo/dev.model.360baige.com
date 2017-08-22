package action

type Num struct {
	Value int64
}

type DeleteByIdCond struct {
	Value []int64
}

type UpdateByCond struct {
	CondList   []CondValue   // 更新条件
	UpdateList []UpdateValue // 更新内容
}

type CondValue struct {
	Type string
	Key  string
	Val  interface{}
}

type UpdateValue struct {
	Key string
	Val interface{}
}

type UpdateByIdCond struct {
	Id         []int64       // 更新条件
	UpdateList []UpdateValue // 更新内容
}

type FindByCond struct {
	CondList []CondValue // 更新条件
	Fileds   []string
}

type DeleteByCond struct {
	CondList []CondValue // 更新条件
}

type ListByCond struct {
	CondList []CondValue
	Cols     []string
	OrderBy  []string
	PageSize int64 //每页数量
}

type PageByCond struct {
	CondList    []CondValue
	Cols        []string
	OrderBy     []string
	Json        []byte
	Total       int64 //总数
	PageSize    int64 //每页数量
	Current     int64 //当前页码
	CurrentSize int64 //当前页数量
}

type CountByCond struct {
	CondList []CondValue // 更新条件
}
