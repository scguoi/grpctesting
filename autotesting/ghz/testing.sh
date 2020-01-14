# 获取当前目录
local_directory="$(cd "$(dirname "$0")" && pwd)"

for casefile in "${local_directory}"/oneway/*.dat; do
  ghz --insecure --call example.H3Wrapper.OneWay 127.0.0.1:18000 -c 1 -n 1 -D "${casefile}"
done

# 流式接口
for casefile in "${local_directory}"/stream/*.dat; do
  ghz --insecure --call example.H3Wrapper.Stream 127.0.0.1:18000 -c 1 -n 1  -D "${casefile}"
done