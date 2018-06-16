.DEFAULT_GOAL := test

generatecert:
	certstrap --depot-path "test/certs" init --passphrase "" --common-name "snowflake ca"
	certstrap --depot-path "test/certs" request-cert --passphrase "" --domain "localhost"
	certstrap --depot-path "test/certs" sign --CA "snowflake ca" "localhost"

proto:
	protoc --proto_path=api api/*.proto --go_out=plugins=grpc:api

runserver:
	PORT=2929 go run cmd/snowflake/snowflake.go

runsecureserver:
	PORT=2929 TLSKEYPATH=test/certs/localhost.key TLSCRTPATH=test/certs/localhost.crt go run cmd/snowflake/snowflake.go

runweb:
	SERVERPORT=2929 go run cmd/snowflakeweb/snowflakeweb.go

test: testunit testintegration teste2e

teste2e:
	ginkgo -r -nodes 4 -randomizeAllSpecs -randomizeSuites -keepGoing -focus end-to-end integration

testintegration:
	ginkgo -r -nodes 4 -randomizeAllSpecs -randomizeSuites -keepGoing -skip end-to-end integration

testunit:
	ginkgo -r -nodes 4 -randomizeAllSpecs -randomizeSuites -keepGoing -skipPackage examplesuite,integration
