ci:
	fly -t vbox set-pipeline -p snowflake -c build/ci/pipeline.yml

proto:
	protoc --proto_path=api api/*.proto --go_out=plugins=grpc:api

runserver:
	PORT=2929 go run cmd/snowflake/snowflake.go

runweb:
	SERVERPORT=2929 go run cmd/snowflakeweb/snowflakeweb.go

test: testunit testintegration

testintegration:
	ginkgo -r integration

testunit:
	ginkgo -r -skipPackage examplesuite,integration
