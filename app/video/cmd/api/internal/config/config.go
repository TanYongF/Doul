package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// OSSConf ths oss config
type OSSConf struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
	TargetPath      string `json:"targetPath"`
	TargetURL       string `json:"targetUrl"`
}

type Config struct {
	rest.RestConf
	UserRpcConf  zrpc.RpcClientConf
	VideoRpcConf zrpc.RpcClientConf
	OSSConf      OSSConf
}
