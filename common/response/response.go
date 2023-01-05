package response

type ResponseSuccessBean struct {
	Code uint32      `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"data,omitempty"`
}

type ResponseErrorBean struct {
	Code uint32 `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func Success(resp interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{
		Code: 0,
		Msg:  "OK",
		Data: resp,
	}
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
