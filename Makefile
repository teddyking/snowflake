.DEFAULT_GOAL := test

generatecert:
	certstrap --depot-path "test/certs" init --passphrase "" --common-name "snowflake ca"
	certstrap --depot-path "test/certs" request-cert --passphrase "" --ip "0.0.0.0"
	certstrap --depot-path "test/certs" sign --CA "snowflake ca" "0.0.0.0"

proto:
	protoc --proto_path=api api/*.proto --go_out=plugins=grpc:api

runsecureserver:
	TLSKEYPATH=test/certs/0.0.0.0.key \
  TLSCRTPATH=test/certs/0.0.0.0.crt \
  go run cmd/snowflake/snowflake.go

runsecureweb:
	TLSCRTPATH=test/certs/0.0.0.0.crt \
  go run cmd/snowflakeweb/snowflakeweb.go

runserver:
	go run cmd/snowflake/snowflake.go

runweb:
	go run cmd/snowflakeweb/snowflakeweb.go

test: testunit testintegration teste2e

teste2e:
	ginkgo -r -nodes 4 -randomizeAllSpecs -randomizeSuites -keepGoing -focus end-to-end integration

testintegration:
	ginkgo -r -nodes 4 -randomizeAllSpecs -randomizeSuites -keepGoing -skip end-to-end integration

testunit:
	ginkgo -r -nodes 4 -randomizeAllSpecs -randomizeSuites -keepGoing -skipPackage examplesuite,integration
