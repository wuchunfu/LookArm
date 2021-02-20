FROM golang:alpine as builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

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