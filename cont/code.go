package cont

// 状态码

// 成功状态
const SUCCESS = 111

// 外部错误:2x
const (
	MISSING_PARAMS = 20
	WRONG_PARAMS   = 21
)

// 内部错误:3x
const INTERNAL_ERROR = 30

// 查询不到
const NOT_FOUND = 40
