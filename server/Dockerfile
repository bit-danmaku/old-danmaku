FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
#ENV GOPROXY="https://mirrors.aliyun.com/goproxy/"
WORKDIR /go/src/old-danmaku
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY Makefile go.mod go.sum ./
#RUN make init && go mod download
RUN GOPROXY=https://goproxy.cn,direct go mod download
COPY . .
RUN make build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/old-danmaku/old-danmaku /old-danmaku
ENTRYPOINT ["/old-danmaku"]
CMD []
