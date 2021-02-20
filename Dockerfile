FROM golang:alpine as builder

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /app/lookarm
COPY . /app/lookarm

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./lookarm"]