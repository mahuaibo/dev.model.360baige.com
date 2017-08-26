package schoolfee

import (
	"strconv"
)

const (
	Normal      = "200"
	ErrorPower  = "300"
	ErrorLogic  = "400"
	ErrorSystem = "500"
)

var (
	ListOfNormal      map[int64]string
	ListOfErrorPower  map[int64]string
	ListOfErrorSystem map[int64]string
	ListOfErrorLogic  map[int64]string
)

func Message(key int64, def ...string) (string) {
	str := ""
	for _, d := range def {
		str += "(" + d + ")"
	}
	if 20000 <= key && key < 30000 {
		return "[" + strconv.FormatInt(key, 10) + "]" + ListOfNormal[key] + str
	} else if key < 40000 {
		return "[" + strconv.FormatInt(key, 10) + "]" + ListOfErrorPower[key] + str
	} else if key < 50000 {
		return "[" + strconv.FormatInt(key, 10) + "]" + ListOfErrorSystem[key] + str
	} else if key < 60000 {
		return "[" + strconv.FormatInt(key, 10) + "]" + ListOfErrorLogic[key] + str
	} else {
		return "[" + strconv.FormatInt(key, 10) + "]" + str
	}
}

func init() {
	ListOfNormal = map[int64]string{
		20000: "获取校园收费列表成功",
		20001: "",
	}
	ListOfErrorPower = map[int64]string{
		30000: "访问令牌无效请重新登录",
	}
	ListOfErrorSystem = map[int64]string{
		40000: "请求参数不能为空",
	}
	ListOfErrorLogic = map[int64]string{
		50000: "访问令牌验证服务异常",
		50001: "数据通讯服务异常",
		50002: "JSON解析服务异常",
	}
}
