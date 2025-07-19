package cont

// 状态码

// 成功状态
const SUCCESS = 111

// 外部错误:2x
const (
	MISSING_PARAMS = 20
	WRONG_PARAMS   = 21
	// 操作步骤出错,一般出现在用户注册时未获取验证码时就点击注册,但应该先获取验证码然后点击注册
	MISSING_STEPS = 22
)

// 内部错误:3x
const INTERNAL_ERROR = 30

// 查询问题
const (
	NOT_FOUND      = 40
	ALREADY_EXISTS = 41
)
