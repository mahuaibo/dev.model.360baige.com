package batch

type BatchModify struct {
	Ids        string //,分隔的ids
	UpdateTime int64  //更新时间
	Status     int8   //更新状态
}
