module gomicro-playground

go 1.13

replace (
	k8s.io/api => k8s.io/api v0.0.0-20191016225839-816a9b7df678
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191016225534-b1267f8c42b4
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20190725011945-5c849dd2c51d // indirect
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.11.3
	github.com/micro/go-plugins v1.3.0
	github.com/olivere/elastic/v7 v7.0.8
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron/v3 v3.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/tebeka/strftime v0.1.3 // indirect
	google.golang.org/genproto v0.0.0-20191009194640-548a555dbc03
	google.golang.org/grpc v1.24.0
	gopkg.in/sohlich/elogrus.v7 v7.0.0
)
