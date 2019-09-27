module gomicro-playground

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190813180838-e711354a0280
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190813060636-0c17871ad6fd
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
)

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20190725011945-5c849dd2c51d // indirect
	github.com/micro/examples v0.2.0 // indirect
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.3.0
	github.com/olivere/elastic v6.2.23+incompatible
	github.com/olivere/elastic/v7 v7.0.4
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.4.2
	github.com/sohlich/elogrus v2.0.2+incompatible
	google.golang.org/grpc v1.23.1
	gopkg.in/olivere/elastic.v5 v5.0.82 // indirect
	gopkg.in/sohlich/elogrus.v7 v7.0.0
)
