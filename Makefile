proto:
	protoc api/api.proto --go_out=plugins=grpc:.

test:
	ginkgo -r
