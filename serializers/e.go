package serializers

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	ErrorGetGroupFail    = 10001
	ErrorDeleteGroupFail = 10002
	ErrorEditGroupFail   = 10003
	ErrorPutGroupFail    = 10004

	ErrorGetTodoFail    = 10101
	ErrorDeleteTodoFail = 10102
	ErrorEditTodoFail   = 10103
	ErrorPutTodoFail    = 10104

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

	ErrorGetTodoFail:    "获取TODO失败",
	ErrorDeleteTodoFail: "删除TODO失败",
	ErrorEditTodoFail:   "编辑TODO失败",
	ErrorPutTodoFail:    "增加TODO失败",

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
