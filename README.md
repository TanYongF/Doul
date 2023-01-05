# Doul
一款go语言开发的分布式短视频后端服务
同时也是**南京邮电大学**的最后一份作业（毕业设计）~

### 接口文档
 使用ApiFox来维护接口，[接口文档在线地址](https://www.apifox.cn/apidoc/shared-b8d7c521-f55a-4b7d-84b2-cf253c111154)

### 技术选型
1. [go-zero](https://go-zero.dev/cn/)：一款开箱即用的RPC框架（基于zrpc）
2. etcd: 服务发现
3. [xid](https://github.com/rs/xid) : UUID生成器

### 关于API

每个微服务都有自己的的API网关

 - UserCenter (用户中心)
 - Comment (评论中心)
 - Friend (关系接口)
 - Video (视频接口)

由于有多个`API`网关，所以应该是使用更上一层的网关来聚合每个单独API

使用`Nginx`作为上层网关，具体配置如下：
```
#doul项目
server{
    listen 8094;
    #access_log /var/log/nginx/looklook.com_access.log;
    #error_log /var/log/nginx/looklook.com_error.log;


    location ~ /douyin/user/ {
      proxy_set_header Host $http_host;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header REMOTE-HOST $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_pass http://127.0.0.1:8889;
    }

    location ~ /douyin/comment/ {
       proxy_set_header Host $http_host;
       proxy_set_header X-Real-IP $remote_addr;
       proxy_set_header REMOTE-HOST $remote_addr;
       proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
       proxy_pass http://127.0.0.1:8888;
    }
}

```

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






## Docker下的Etcd安装（单机安装）

```shell
docker pull bitnami/etcd:latest #拉取镜像

docker network create app-tier --driver bridge #构建Docker网络

#运行etcd服务端
docker run -d --name etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
    
    
##客户端--可选
docker run -it --rm \
    --network app-tier \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest etcdctl --endpoints http://etcd-server:2379 put /message Hello

```

