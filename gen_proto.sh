cd pb
rm *.go
cd ..
protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb