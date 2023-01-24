package interceptor

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go_code/Doul/common/xerr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//
// LoggerInterceptor
//  @Description: 	when rpc-server throw errors to api-server, change the state-error type
// 				 	ErrorCode (Customized), so that api-server can know that are the errors
//				 	which from rpc-server belong type of Customized ErrorCode:
//					if yes : api-server can throw the rpc-error to user
//					if no  : api-server will not throw the error.
//  @param ctx
//  @param req
//  @param info
//  @param handler
//  @return resp
//  @return err
//
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err) // err类型
		// 如果错误类型是CoderError自定义类型，那么就会返回值给用户
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
			//转成grpc err
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		} else {
			logx.WithContext(ctx).Errorf("【RPC-SRV-ERR】 %+v", err)
		}
	}
	return resp, err
}
