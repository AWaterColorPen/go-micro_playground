module gomicro-playground

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190813180838-e711354a0280
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190813060636-0c17871ad6fd
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)

go 1.13

require (
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/hashicorp/consul/api v1.2.0 // indirect
	github.com/hashicorp/memberlist v0.1.5 // indirect
	github.com/hashicorp/serf v0.8.4 // indirect
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20190725011945-5c849dd2c51d // indirect
	github.com/micro/examples v0.2.0 // indirect
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.10.0 // indirect
	github.com/miekg/dns v1.1.19 // indirect
	github.com/nlopes/slack v0.6.0 // indirect
	github.com/olivere/elastic v6.2.23+incompatible // indirect
	github.com/olivere/elastic/v7 v7.0.4
	github.com/opentracing/opentracing-go v1.1.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron/v3 v3.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/sohlich/elogrus v2.0.2+incompatible // indirect
	github.com/tebeka/strftime v0.1.3 // indirect
	golang.org/x/crypto v0.0.0-20190927123631-a832865fa7ad // indirect
	golang.org/x/net v0.0.0-20190926025831-c00fd9afed17 // indirect
	golang.org/x/sys v0.0.0-20190927073244-c990c680b611 // indirect
	google.golang.org/genproto v0.0.0-20190927181202-20e1ac93f88c
	google.golang.org/grpc v1.24.0
	gopkg.in/olivere/elastic.v5 v5.0.82 // indirect
	gopkg.in/sohlich/elogrus.v7 v7.0.0
)
