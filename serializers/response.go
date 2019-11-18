package serializers

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BuildResponse(errorCode int, data interface{}) interface{} {
	return CommonResponse{
		Code: errorCode,
		Msg:  GetMsg(errorCode),
		Data: data,
	}
}
