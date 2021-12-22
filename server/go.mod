module old-danmaku

go 1.16

require (
	github.com/asim/go-micro/plugins/server/http/v3 v3.0.0-20210924081004-8c39b1e1204d
	github.com/asim/go-micro/v3 v3.5.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.4
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
