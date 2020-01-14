# 使用命令行进行gRPC接口测试

本仓库存存放一些本人收集的gRPC接口命令行测试工具的使用，包含以下部分

> 1. 一个example的golang gRPC服务，包括一个流式接口，一个非流式接口；
> 2. 如何使用[grpcurl](https://github.com/fullstorydev/grpcurl)进行流式和非流式接口的请求；
> 3. 如何使用[ghz](https://github.com/bojand/ghz)进行流式和非流式接口的压力测试；
> 4. 如何使用命令行工具进行自动化测试；

## golang gRPC example

该测试工程使用go mod进行依赖管理，如果依赖包下载缓慢，可以使用下面的代理进行下载。

```shell script
GOPROXY="https://goproxy.cn,direct"
GOSUMDB="off"
```

启动服务

```shell script
go run main.go
```

接口声明

```proto
syntax = "proto3";

package example;

message H3Req {
    string msgType = 1;
    map<string, string> headers = 2;
    bytes body = 3;
    bool fin = 4;
}

message H3Resp {
    string msgType = 1;
    map<string, string> headers = 2;
    bytes body = 3;
    bool fin = 4;
}

// wrapper service
service H3Wrapper {
    rpc Stream (stream H3Req) returns (stream H3Resp);
    rpc OneWay (H3Req) returns (H3Resp);
}
```

## 整体说明

本文提到的两个工具都是通过json发起调用的，json和pb之间的转化关系可以参考[这里](https://developers.google.com/protocol-buffers/docs/proto3#json)

强调一下pb中的二进制bytes对应着json中的base64字符串。

## grpcurl

### 获取服务信息

```shell script
# 获取所有服务
grpcurl -plaintext 127.0.0.1:18000 list

# 获取服务的方法
grpcurl -plaintext 127.0.0.1:18000 list example.H3Wrapper

# 获取方法的声明
grpcurl -plaintext 127.0.0.1:18000 describe example.H3Wrapper.Stream

# 获取message的定义
grpcurl -plaintext 127.0.0.1:18000 describe .example.H3Req
```

### OneWay接口请求

```shell script
grpcurl -plaintext -d '{"msgType":"header","headers":{"content-type":"application/json"},"fin":false}' 127.0.0.1:18000 example.H3Wrapper/OneWay
```

### Stream接口请求

对于流式接口，grpcurl的处理逻辑是每行一个json，使用控制台输入
```
➜  ~ grpcurl -plaintext -d @ 127.0.0.1:18000 example.H3Wrapper/Stream
# 输入
{"msgType":"header","headers":{"content-type":"application/json"},"fin":false}
# 返回
{
  "msgType": "header"
}
# 输入
{"msgType":"data","body":"Z3Vvc29uZ2NodQ==","fin":false}
# 返回
{
  "msgType": "data",
  "body": "c2VydmVyIHJldHVybg=="
}
# 输入
{"msgType":"finish","body":"Z3Vvc29uZ2NodQ==","fin":true}
# 返回
{
  "msgType": "data",
  "body": "c2VydmVyIHJldHVybg==",
  "fin": true
}
```

## ghz

### OneWay接口请求

### Stream接口请求

## 自动化测试

### grpcurl

