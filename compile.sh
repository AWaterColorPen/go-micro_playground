protoc --go_out=paths=source_relative:. \
    proto/common/common.proto

protoc --nitro_out=paths=source_relative:. --go_out=paths=source_relative:. \
    proto/akin/akin.proto

protoc --nitro_out=paths=source_relative:. --go_out=paths=source_relative:. \
    proto/tosui/tosui.proto