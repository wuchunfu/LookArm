FROM golang:alpine as builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /app/lookarm
COPY . /app/lookarm

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./lookarm"]