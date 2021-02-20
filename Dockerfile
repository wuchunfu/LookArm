FROM golang:alpine as builder

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

#定义时区参数
ENV TZ=Asia/Shanghai

#设置时区
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo '$TZ' > /etc/timezone

WORKDIR /app/lookarm
COPY . /app/lookarm

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./lookarm"]