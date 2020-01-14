# 获取当前目录
local_directory="$(cd "$(dirname "$0")" && pwd)"

# 获取所有服务
grpcurl -plaintext 127.0.0.1:18000 list

# 获取服务的方法
grpcurl -plaintext 127.0.0.1:18000 list example.H3Wrapper

# 获取方法的声明
grpcurl -plaintext 127.0.0.1:18000 describe example.H3Wrapper.Stream

# 获取message的定义
grpcurl -plaintext 127.0.0.1:18000 describe .example.H3Req

for casefile in "${local_directory}"/stream/*.dat; do
  grpcurl <"${casefile}" -plaintext -d @ 127.0.0.1:18000 example.H3Wrapper/Stream
done

for casefile in "${local_directory}"/oneway/*.dat; do
  grpcurl <"${casefile}" -plaintext -d @ 127.0.0.1:18000 example.H3Wrapper/OneWay
done
