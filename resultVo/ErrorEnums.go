/**
 * @Author: AF
 * @Date: 2021/8/9 17:08
 */

package resultVo

const (
	SUCCESS       = 200
	SERVER_ERROR  = 10500
	BAD_REQUEST   = 10400
	UN_AUTH       = 10401
	FORBIDDEN     = 10403
	NOT_FOUND     = 10404
	RECORD_EXISTS = 20001
)

var codeMsg = map[int]string{
	SUCCESS:       "请求成功",
	SERVER_ERROR:  "服务器异常",
	BAD_REQUEST:   "请求参数不正确",
	UN_AUTH:       "未授权",
	FORBIDDEN:     "暂无权限",
	NOT_FOUND:     "资源不存在",
	RECORD_EXISTS: "记录已存在",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
