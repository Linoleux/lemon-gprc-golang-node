gen: 
	./gen_proto.sh

server:
	go run ./cmd/server/main.go --port 50051

client:
	go run ./cmd/client/main.go --address 127.0.0.1:50051 --name Lino

node-install:
	./nodejs/install.sh

node:
	node ./nodejs/client.js --addr 127.0.0.0:50051
