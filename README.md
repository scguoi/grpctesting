# 使用命令行进行gRPC接口测试

本仓库存存放一些本人收集的gRPC接口命令行测试工具的使用，包含以下部分

> 1. 一个example的golang gRPC服务，包括一个流式接口，一个非流式接口；
> 2. 如何使用[grpcurl](https://github.com/fullstorydev/grpcurl)进行流式和非流式接口的请求；
> 3. 如何使用[ghz](https://github.com/bojand/ghz)进行流式和非流式接口的压力测试；

## golang gRPC example

该测试工程使用go mod进行依赖管理，如果依赖包下载缓慢，可以使用下面的代理进行下载。

```shell script
GOPROXY="https://goproxy.cn,direct"
GOSUMDB="off"
```

启动服务

```shell script
go run
```