# 编译源代码
FROM golang:1.16 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

# 编译完的二进制可执行文件拷贝到镜像
FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./server", "-conf", "/data/conf"]
