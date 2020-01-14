# 获取当前目录
local_directory="$(cd "$(dirname "$0")" && pwd)"

for casefile in "${local_directory}"/stream/*.dat; do
  grpcurl <"${casefile}" -plaintext -d @ 127.0.0.1:18000 example.H3Wrapper/Stream
done

for casefile in "${local_directory}"/oneway/*.dat; do
  grpcurl <"${casefile}" -plaintext -d @ 127.0.0.1:18000 example.H3Wrapper/OneWay
done
