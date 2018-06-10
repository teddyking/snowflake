proto:
	protoc --proto_path=api api/*.proto --go_out=plugins=grpc:api

runserver:
	PORT=2929 go run cmd/snowflake/snowflake.go

runweb:
	SERVERPORT=2929 go run cmd/snowflakeweb/snowflakeweb.go

test: testunit testintegration teste2e

teste2e:
	ginkgo -r -randomizeAllSpecs -randomizeSuites -focus end-to-end integration

testintegration:
	ginkgo -r -randomizeAllSpecs -randomizeSuites -skip end-to-end integration

testunit:
	ginkgo -r -randomizeAllSpecs -randomizeSuites -skipPackage examplesuite,integration
