# Doul
一款go语言开发的分布式短视频后端服务
同时也是**南京邮电大学**的最后一份作业（毕业设计）~

### 接口文档
 使用ApiFox来维护接口，[接口文档在线地址](https://www.apifox.cn/apidoc/shared-b8d7c521-f55a-4b7d-84b2-cf253c111154)

### 技术选型
1. [go-zero](https://go-zero.dev/cn/)：一款开箱即用的RPC框架（基于zrpc）
2. etcd: 服务发现
3. [xid](https://github.com/rs/xid) : UUID生成器

### 关于日志处理方面

在gRPC中，当调用rpc服务发生错误时， 其返回的类型是 `status.error` , 它是通过`status code`来判断错误类型 具体可查看可查看[Status Code](https://grpc.github.io/grpc/core/md_doc_statuscodes.html)

例如，当关闭Redis时，由于rpc服务连接Redis超时，那么会返回调用层一个`Status Code = 4`的`error`来代表执行超时。

> 查看文档得知： DEADLINE_EXCEEDED  4 means that **The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully. For example, a successful response from a server could have been delayed long**

![image-20230101001757364](https://kauizhaotan.oss-accelerate.aliyuncs.com/img/image-20230101001757364.png)


我们通过自定义错误类型`Code Error` ,来生成`Customized error`返回类型, 以便更好的日志处理

```go
type CodeError struct {
	errCode uint32
	errMsg  string
}
```

在rpc服务中：

- 对于业务错误（用户密码失败等），通过添加拦截器，rpc服务执行完`logic`逻辑后，将我们在rpc服务产生的`CodeError`错误包装成`grpc error`返回到调用方，也就是API服务
- 对于grpc错误，直接返回grpc类型错误

在api中：

- 底层错误 DB等错误不应该抛出给用户
- 业务错误应该抛出给用户

几篇文章：

- [Go error 最佳实践](https://medium.com/@dche423/golang-error-handling-best-practice-cn-42982bd72672)
- https://chanjarster.github.io/post/go/err-throw-rules/