FROM golang:alpine as builder

RUN  apk add --update --no-cache yarn make g++

ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
    GO111MODULE=on \
    CGO_ENABLED=1

RUN apk add --no-cache  gettext tzdata   && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" >  /etc/timezone && \
    date && \
    apk del tzdata

WORKDIR /app/lookarm
COPY . /app/lookarm

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./lookarm"]