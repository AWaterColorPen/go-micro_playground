protoc --micro_out=. --go_out=plugins=grpc:. \
    proto/common.proto \
    proto/akin.proto \
    proto/tosui.proto

protoc --grpc-gateway_out=grpc_api_configuration=proto/akin.yaml:. \
   proto/akin.proto

cp -r github.com/AWaterColorPen/go_micro_playground/tencho/. ./proto
rm -rf github.com

for f in proto/*.pb.gw.go
do
  sed -i '' 's/Register\([A-Za-z]*\)Handler(/Register\1Handler_(/g' "$f"
done

