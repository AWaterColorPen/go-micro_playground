module gomicro-playground

replace (
	github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.7.1-0.20190913061013-f15a82d3fdc3
	k8s.io/api => k8s.io/api v0.0.0-20190813180838-e711354a0280
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190813060636-0c17871ad6fd
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.1 // indirect
	github.com/hashicorp/consul/api v1.2.0 // indirect
	github.com/hashicorp/memberlist v0.1.5 // indirect
	github.com/hashicorp/serf v0.8.4 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20190725011945-5c849dd2c51d // indirect
	github.com/micro/examples v0.2.0 // indirect
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
	github.com/micro/micro v1.10.0 // indirect
	github.com/miekg/dns v1.1.19 // indirect
	github.com/nlopes/slack v0.6.0 // indirect
	github.com/olivere/elastic v6.2.23+incompatible
	github.com/olivere/elastic/v7 v7.0.4
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.4.2
	github.com/sohlich/elogrus v2.0.2+incompatible
	golang.org/x/crypto v0.0.0-20190927123631-a832865fa7ad // indirect
	golang.org/x/net v0.0.0-20190926025831-c00fd9afed17 // indirect
	golang.org/x/sys v0.0.0-20190927073244-c990c680b611 // indirect
	google.golang.org/genproto v0.0.0-20190927181202-20e1ac93f88c // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/olivere/elastic.v5 v5.0.82 // indirect
	gopkg.in/sohlich/elogrus.v7 v7.0.0
)
