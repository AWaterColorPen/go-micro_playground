protoc --plugin=protoc-gen-go=/Users/slyao/go/bin/protoc-gen-go --plugin=protoc-gen-micro=/Users/slyao/go/bin/protoc-gen-micro --micro_out=. --go_out=. \
    proto/common.proto \
    proto/akin.proto \
    proto/tosui.proto
