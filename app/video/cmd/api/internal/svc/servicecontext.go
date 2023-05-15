package svc

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go_code/Doul/app/usercenter/cmd/rpc/userclient"
	"go_code/Doul/app/video/cmd/api/internal/config"
	"go_code/Doul/app/video/cmd/rpc/video"
	"go_code/Doul/common/middleware"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	UserRpc        userclient.User
	VideoRpc       video.Video
	OSSEngine      *oss.Bucket
}

func NewServiceContext(c config.Config) *ServiceContext {
	userClient := userclient.NewUser(zrpc.MustNewClient(c.UserRpcConf))
	videoClient := video.NewVideo(zrpc.MustNewClient(c.VideoRpcConf))

	//获取oss bucket
	client, _ := oss.New(c.OSSConf.Endpoint, c.OSSConf.AccessKeyId, c.OSSConf.AccessKeySecret)
	bucket, _ := client.Bucket(c.OSSConf.BucketName)

	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(userClient).Handle,
		UserRpc:        userClient,
		VideoRpc:       videoClient,
		OSSEngine:      bucket,
	}
}
