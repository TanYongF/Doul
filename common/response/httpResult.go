package response

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go_code/Doul/common/xerr"
	"google.golang.org/grpc/status"
	"net/http"
)

// HttpResult Union Http Response Func
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功直接返回
		//r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, resp)
	} else {
		//默认错误返回，当不是自定义错误时候，返回此错误
		errCode := xerr.SERVER_COMMON_ERROR
		errMsg := xerr.Message[errCode]

		causeErr := errors.Cause(err) // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok {
			//If the error is type of Customized Error, return it to user
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			//If the error is the error from rpc, check the grpc-error's states code is the Customized Error's code
			if gStatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gStatus.Code())
				//区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
				if xerr.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMsg = gStatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
	}
}
