package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int         `json:"status_code"`
	Msg  string      `json:"status_msg"`
	Data interface{} `json:"data,omitempty"`
}

// Response 返回请求格式
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = 0
		body.Msg = err.Error()
	} else {
		body.Code = 200
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
