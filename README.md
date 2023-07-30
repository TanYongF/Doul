# Doul
##  一、项目介绍

一款go语言开发的分布式短视频后端服务同时也是**南京邮电大学**的最后一份作业（毕业设计）~
此项目也是基于字节青训营项目基础上继续改进而成，项目链接：[https://github.com/TanYongF/douyin]

该系统注重的几个特点：

- 低代码，得益于`Goctl`工具
- 分布式系统，基于`grpc`协议
- 拥抱CI/CD，编码、构建、部署自动化
- 高可用（努力做到…）

## 二、 开发文档

### 0. 系统架构

包含部分技术，当然还不完善……

![未命名文件](https://kauizhaotan.oss-accelerate.aliyuncs.com/img/未命名文件.jpg?x-oss-process=style/water)

在线文档：https://www.processon.com/view/link/6279170d1efad40df02ee683

### 1. API接口

API接口是指服务端暴露给Web、客户端、Browser的接口，其是通过HTTP协议来传输的。

 该项目使用[APIFOX](https://www.apifox.cn/)来管理和维护接口，同时也可以对接口进行快速测试和代码快速生成。[接口文档在线地址](https://www.apifox.cn/apidoc/shared-b8d7c521-f55a-4b7d-84b2-cf253c111154)

### 2. 部分技术选型
1. [go-zero](https://go-zero.dev/cn/)：一款开箱即用的RPC框架（基于zrpc），快速交付是其一大特点；
2. [etcd](https://etcd.io/): 一种高性能、分布式K-V键值对存储系统，该项目作为**服务发现**和注册中心使用；
3. [xid](https://github.com/rs/xid) : UUID生成器
4. [Nginx](https://www.nginx.com/):高性能服务器，用来作为网关提供**负载均衡**以及**反向代理**；
5. [Docker](https://www.docker.com/): Devops， 用来快速部署各个微服务。
6. [Redis](http://www.redis.cn/)、[Mysql](www.mysql.com)等缓存以及数据库组件等
7. 其他待完善…

### 3. 服务划分

主要针对

系统暂定4个微服务，每个微服务及其实现功能如下：

 - UserCenter (用户中心)
   - 用户注册
   - 用户登陆
   - 用户鉴权（特定需要Auth的端口）
   - 获取用户信息等

 - Comment (评论中心)
   - 评论列表
   - 评论操作
   - 敏感词过滤

 - Friend (关系接口)
   - 粉丝列表
   - 关注列表
   - 关注操作
 - Video (视频接口)
   - 基础视频流接口（搭配**推荐算法**）
   - 点赞
   - 投稿列表
   - 发布列表


## 三、问题文档

这里主要整理开发过程中的一些问题和解决的方法思路。当然，也会参考优秀开源项目的经验。

### 1. Nginx的引入

每个微服务都有自己的的API网关，一个API后面会调用多个RPC服务，所以后期一个RPC服务的修改就需要重构整个API服务，所以为每个微服务创建自己的网关，因此上层必须使用**统一网关**来做流量分发，此项目使用`Nginx`作为统一流量入口，具体作用就是流量通过Nginx服务器分发给各个微服务的网关上。具体参考[使用Nginx作为网关](https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/doc/chinese/2-nginx%E7%BD%91%E5%85%B3.md)

常用命令：
```shell
Usage: nginx [-?hvVtTq] [-s signal] [-p prefix]
             [-e filename] [-c filename] [-g directives]
 #-t : 检测当前配置文件是否正确
 #-c : 指定配置文件路径 
 #-s signal : send signal to a master process: stop, quit, reopen, reload
nginx -s [reload/start] # 重载配置
```

本项目的Nginx配置如下：

```nginx
#doul项目
server{
    listen 8094; #监听8094端口
    #access_log /var/log/nginx/looklook.com_access.log;
    #error_log /var/log/nginx/looklook.com_error.log;


     # 将不同路径映射到不同服务器的不同端口
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

参考文档: 

1. [Nginx location配置](https://juejin.cn/post/6844903849166110733)

### 2. 日志处理方面

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

### 3. Etcd安装（单机安装）

etcd作为服务注册中心，是启动项目的前置要求，下面是在本人的`Centos 7`上通过Docker安装的方式。

```shell
docker pull bitnami/etcd:latest #拉取镜像

docker network create app-tier --driver bridge #构建Docker网络

#运行etcd服务端
docker run -d --name etcd-server \
    -p 2379:2379  -p 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest
    -advertise-client-urls http://0.0.0.0:2379 \
    -listen-client-urls http://0.0.0.0:2379 \

    docker run -d --name etcd-server \
    --network app-tier \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://81.68.239.206:2379 \
    bitnami/etcd:latest
 
##客户端--可选
docker run -it --rm \
    --env ALLOW_NONE_AUTHENTICATION=yes \

    bitnami/etcd:latest etcdctl --endpoints http://localhost:2379 put /message Hello
    
    
    
docker run -it --rm \
    --network app-tier \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest etcdctl --endpoints http://81.68.239.206:2379    get --prefix "Hello"
    

```
几个重要参数：
>--advertise-client-urls
就是客户端(etcdctl/curl等)跟etcd服务进行交互时请求的url
--listen-client-urls
这个参数是etcd服务器自己监听时用的，也就是说，监听本机上的哪个网卡，哪个端口

作者：码二哥
链接：https://www.jianshu.com/p/7bbef1ca9733
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

### 4. sqlx踩坑

注意查询多行是 `c.QueryRows` 而查询单行是`c.QueryRow`且查询出的字段数量以及名称应该和结构体字段完全匹配

参考资料



#### 5. Window安装Docker

打开`windows cmd`，安装`WSL2.0`, 执行

```shell
wsl --install #可选参数 -d 列出所有分发版本
```

然后自动安装一个分发版本，默认是ubutun

常用命令

```shell
wsl #启动Linux虚拟机
wsl --shutdown # 关闭虚拟机

```

##### 1. 登录阿里云Docker Registry

```
$ docker login --username=谭永锋nb registry.cn-heyuan.aliyuncs.com
```

用于登录的用户名为阿里云账号全名，密码为开通服务时设置的密码。

您可以在访问凭证页面修改凭证密码。

##### 2. 从Registry中拉取镜像

```
$ docker pull registry.cn-heyuan.aliyuncs.com/doul/usercenter-rpc:[镜像版本号]
```

##### 3. 将镜像推送到Registry

```
$ docker login --username=谭永锋nb registry.cn-heyuan.aliyuncs.com$ docker tag [ImageId] registry.cn-heyuan.aliyuncs.com/doul/usercenter-rpc:[镜像版本号]$ docker push registry.cn-heyuan.aliyuncs.com/doul/usercenter-rpc:[镜像版本号]
```

请根据实际镜像信息替换示例中的[ImageId]和[镜像版本号]参数。

##### 4. 选择合适的镜像仓库地址

从ECS推送镜像时，可以选择使用镜像仓库内网地址。推送速度将得到提升并且将不会损耗您的公网流量。

如果您使用的机器位于VPC网络，请使用 registry-vpc.cn-heyuan.aliyuncs.com 作为Registry的域名登录。

##### 5. 示例

使用"docker tag"命令重命名镜像，并将它通过专有网络地址推送至Registry。

```
$ docker imagesREPOSITORY                                                         TAG                 IMAGE ID            CREATED             VIRTUAL SIZEregistry.aliyuncs.com/acs/agent                                    0.7-dfb6816         37bb9c63c8b2        7 days ago          37.89 MB$ docker tag 37bb9c63c8b2 registry-vpc.cn-heyuan.aliyuncs.com/acs/agent:0.7-dfb6816
```

使用 "docker push" 命令将该镜像推送至远程。

```
$ docker push registry-vpc.cn-heyuan.aliyuncs.com/acs/agent:0.7-dfb6816
```

### MQ安装

https://juejin.cn/post/6844903970545090574

新版本的MQ需要手动打开WEB管理

```shell
docker exec -i -t <containerId> bin/bash #进入容器内部 记得替换containerId
rabbitmq-plugins enable rabbitmq_management #打开web管理服务
```

MQ的Connection和Channel区别 [阿里云](https://help.aliyun.com/document_detail/178124.html)

- 一个进程维护一条`Connection`
- 一个进程的不同线程对应一条`Channel`

### Kibana 和 ElasticSearch 安装
此部分记录一下Docker方式下，Kibana和ES工具的安装。
> Kibana version : 8.7.0
> 
> ElasticSearch version : 8.7.0


**一些文档**
- [官方文档](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)
- [安全配置](https://blog.csdn.net/The_Time_Runner/article/details/111409319)




0. 创建网络

```bash
docker network create es-net
````

1. 安装ElasticSearch
```dockerfile
docker run -d \
	--name elasticsearch \
    -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
    -e "discovery.type=single-node" \
    -v es-data:/usr/share/elasticsearch/data \
    -v es-plugins:/usr/share/elasticsearch/plugins \
    --privileged \
    --network es-net \
    -p 9200:9200 \
    -p 9300:9300 \
elasticsearch:8.7.0
```


2. 安装Kibana

```dockerfile
docker run -d \
--name kibana \
-e ELASTICSEARCH_HOSTS=http://elasticsearch:9200 \
--network es-net \
-p 5601:5601  \
kibana:8.7.0
```

3. 更改es密码

```bash
docker exec -it [es容器ID] bash #进入容器
elasticsearch-setup-passwords interactive #交互式的修改几个账号密码
exit #修改完成退出 
```
访问 http://127.0.0.1:9200 测试结果

4. 修改kibana配置

```bash
docker exec -it [kibana容器ID] bash #进入容器
vim /usr/share/kibana/config/kibana.yml #修改配置, 配置如下
exit #修改完成后退出
```

```editorconfig
server.name: kibana
server.host: "0"
elasticsearch.hosts: [ "http://ip:9200" ] # elasticsearch的ip，有时候localhost不行
monitoring.ui.container.elasticsearch.enabled: true
elasticsearch.username: "kibana_system" # 刚才提到的Kibana用的built-in user
elasticsearch.password: "111111" # 第二步里自己设置的密码

```

修改完成退出，重启，访问 http://127.0.0.1:5601 , 
账号： elastic
密码： 第3步设置的密码



