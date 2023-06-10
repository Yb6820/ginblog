package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	EMAIL_ERROR = 11
	//redis错误
	REDIS_SET_VERIFY_ERROR = 601
	REDIS_GET_VERIFY_ERROR = 602
	//code = 1000...用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004 //token不存在
	ERROR_TOKEN_RUNTIME    = 1005 //token过时
	ERROR_TOKEN_WRONG      = 1006 //token解析错误
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NOT_RIGHT   = 1008
	ERROR_VERIFY_NOT_SAME  = 1009 //验证码不一致
	//code = 2000...文章模块的错误
	ERROR_ART_NOT_EXIST = 2001
	//code = 3000...分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在!",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKNE已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
	ERROR_CATENAME_USED:    "该分类已存在!",
	ERROR_ART_NOT_EXIST:    "文章不存在",
	ERROR_CATE_NOT_EXIST:   "该分类不存在",
	ERROR_USER_NOT_RIGHT:   "该用户无权限",
	EMAIL_ERROR:            "验证码发送失败",
	REDIS_SET_VERIFY_ERROR: "set验证码出错",
	REDIS_GET_VERIFY_ERROR: "get验证码出错",
	ERROR_VERIFY_NOT_SAME:  "验证码不一致",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
