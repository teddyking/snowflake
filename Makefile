ci:
	fly -t vbox set-pipeline -p snowflake -c build/ci/pipeline.yml

dockerimage:
	docker build .

proto:
	protoc api/api.proto --go_out=plugins=grpc:.

runclient:
	PORT=80 go run cmd/client/snfc.go

runserver:
	PORT=2929 go run cmd/server/snowflake.go

test:
	ginkgo -r -skipPackage examplesuite
