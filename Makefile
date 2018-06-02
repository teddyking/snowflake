ci:
	fly -t vbox set-pipeline -p snowflake -c build/ci/pipeline.yml

proto:
	protoc --proto_path=api api/*.proto --go_out=plugins=grpc:api

runclient:
	PORT=2929 go run cmd/client/snfc.go

runserver:
	PORT=2929 go run cmd/server/snowflake.go

test:
	ginkgo -r -skipPackage examplesuite
