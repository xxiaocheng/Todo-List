package serializers

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorGetGroupFail    = 10001
	ErrorDeleteGroupFail = 10002
	ErrorEditGroupFail   = 10003
	ErrorPutGroupFail    = 10004

	ErrorGetTaskFail    = 10101
	ErrorDeleteTaskFail = 10102
	ErrorEditTaskFail   = 10103
	ErrorPutTaskFail    = 10104

	ErrorPutUserFail        = 10204
	ErrorChangePasswordFail = 10205

	ErrorAuthCheckTokenFail    = 30001
	ErrorAuthCheckTokenTimeout = 30002
	ErrorAuthTokenGenFail      = 30003
	ErrorAuth                  = 30004
)

var ErrorMsgMap = map[int]string{
	Success:       "OK",
	Error:         "Fail",
	InvalidParams: "请求参数错误",

	ErrorGetGroupFail:    "获取分组失败",
	ErrorDeleteGroupFail: "删除分组失败",
	ErrorEditGroupFail:   "编辑分组失败",
	ErrorPutGroupFail:    "增加分组失败",

	ErrorGetTaskFail:    "获取TODO失败",
	ErrorDeleteTaskFail: "删除TODO失败",
	ErrorEditTaskFail:   "编辑TODO失败",
	ErrorPutTaskFail:    "增加TODO失败",

	ErrorPutUserFail:        "增加用户失败",
	ErrorChangePasswordFail: "修改用户密码失败",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthTokenGenFail:      "Token生成失败",
	ErrorAuth:                  "Token错误",
}

func GetMsg(code int) string {
	msg, ok := ErrorMsgMap[code]
	if ok {
		return msg
	}
	return ErrorMsgMap[Error]
}
