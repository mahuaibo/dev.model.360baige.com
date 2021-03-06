package user

// 运营账号信息
const (
	UserPositionCompanyIdAudit int64 = 100
	UserIdAudit                int64 = 100
	UserPositionIdAudit        int64 = 100
)

// 默认头像
const HEAD = "User/HeadImages/daeh1835905051.png"

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
	UserPositionTypeVisitor int = 500 // 大众 500 ~ 600
)

func UserPositionName(UserPositionType int) string {
	if UserPositionType == UserPositionTypeDev {
		return "开发者"
	} else if UserPositionType == UserPositionTypeAudit {
		return "运营商"
	} else if UserPositionType == UserPositionTypeService {
		return "服务商"
	} else if UserPositionType == UserPositionTypeSchool {
		return "管理员"
	} else if UserPositionType == UserPositionTypeStudent {
		return "学生"
	} else if UserPositionType == UserPositionTypeVisitor {
		return "大众"
	} else {
		return "其他"
	}
}

// 用户时效
const (
	UserExpireIn                int64 = 60 * 1000
	UserPositionExpireIn        int64 = 3600 * 1000 * 24 * 30
	UserPositionTransitExpireIn int64 = UserPositionExpireIn / 2
)
