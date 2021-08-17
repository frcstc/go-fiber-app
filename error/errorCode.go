package businessError

const (
	SUCCESS       = 200
	UNKNOW_ERR = 201
	SERVER_ERROR  = 10500
	BAD_REQUEST   = 10400
	UN_AUTH       = 10401
	FORBIDDEN     = 10403
	NOT_FOUND     = 10404
	PARAMS_ERROR = 10405
	RECORD_EXISTS = 10406
	TOKEN_CREATE_ERROR = 30001
	TOKEN_PARSE_ERROR = 30002
	TOKEN_NOT_EXISTS = 30003
	TOKEN_EXPIRED = 30004
	USER_NOT_FOUND = 40404
	USER_LOCKED = 40405
	USER_LOGIN_ERROR = 40406
)

var codeMsg = map[int]string{
	SUCCESS:       "请求成功",
	UNKNOW_ERR: "未知的业务异常",
	SERVER_ERROR:  "服务器异常",
	BAD_REQUEST: "请求错误",
	PARAMS_ERROR:   "请求参数不正确",
	UN_AUTH:       "未授权",
	FORBIDDEN:     "暂无权限",
	NOT_FOUND:     "资源不存在",
	RECORD_EXISTS: "记录已存在",
	TOKEN_CREATE_ERROR: "登录凭证生成失败",
	TOKEN_PARSE_ERROR: "登录凭证不存在或已过期",
	TOKEN_NOT_EXISTS: "登录凭证不存在",
	TOKEN_EXPIRED: "登陆凭证已失效，请重新登录",
	USER_NOT_FOUND: "用户不存在",
	USER_LOCKED: "用户已被锁定",
	USER_LOGIN_ERROR: "账号或密码错误",
}

func GetMsg(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	} else {
		return codeMsg[UNKNOW_ERR]
	}

}
