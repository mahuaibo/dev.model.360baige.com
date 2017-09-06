package user

// 运营账号信息

const (
	UserPositionCompanyIdAudit int64 = 100
	UserIdAudit                int64 = 100
	UserPositionIdAudit        int64 = 100
)

// 用户状态
const (
	UserStatusCancel int = -1
	UserStatusNormal int = 0
)

// 游客初始公司ID
const UserPositionCompanyIdInit int64 = 0

// 用户身份类型
const (
	UserPositionTypeDev     int = 0   // 开发者 0 ~ 100
	UserPositionTypeAudit   int = 100 // 运营商 100 ~ 200
	UserPositionTypeService int = 200 // 服务商 200 ~ 300
	UserPositionTypeSchool  int = 300 // 管理员 300 ~ 400
	UserPositionTypeStudent int = 400 // 学生 400 ~ 500
	UserPositionTypeVisitor int = 500 // 游客 500 ~ 600
)
