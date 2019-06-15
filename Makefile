OS := $(shell uname -s |tr '[A-Z]' '[a-z]')

build: deps
	GOOS=$(OS) GOARCH=amd64 go build -o dist/itar_$(OS) cmd/cmd.go

deps:
	# dep ensure
	# protoc --proto_path=pkg/proto/sdp_exchange --go_out=plugins=grpc:pkg/proto/sdp_exchange pkg/proto/sdp_exchange/sdp_exchange.proto
	mkdir -p dist
