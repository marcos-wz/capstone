GOPATH=$(shell go env GOPATH)
SCHEMAS_SRC=$(GOPATH)/src/github.com/marcos-wz/capstone/proto

# PROTO GENERATION  *********************************************************************
generate: proto/basepb/base.proto proto/filterpb/filter.proto proto/loaderpb/loader.proto proto/filterccpb/filtercc.proto proto/fruitpb/fruit.proto
	protoc --proto_path=$(SCHEMAS_SRC)/basepb/ --go_out=$(GOPATH)/src/ base.proto
	protoc --proto_path=$(SCHEMAS_SRC)/filterpb/ --proto_path=$(GOPATH)/src/ --go_out=$(GOPATH)/src/ filter.proto
	protoc --proto_path=$(SCHEMAS_SRC)/loaderpb/ --go_out=$(GOPATH)/src/ loader.proto
	protoc --proto_path=$(SCHEMAS_SRC)/filterccpb/ --go_out=$(GOPATH)/src/ filtercc.proto
	protoc --proto_path=$(SCHEMAS_SRC)/fruitpb/ --proto_path=$(GOPATH)/src/ --go_out=$(SCHEMAS_SRC)/fruitpb/ --go_opt=paths=source_relative --go-grpc_out=$(SCHEMAS_SRC)/fruitpb/ --go-grpc_opt=paths=source_relative fruit.proto

# MOCKS GENERATION  *********************************************************************
mocks:
	mockery --name=FruitRepo --srcpkg=./internal/service --output=./internal/service/mocks
	mockery --name=FruitService --srcpkg=./internal/server --output=./internal/server/mocks


# SSL CERTIFICATES GENERATION  **********************************************************
# ca.key: Certificate Authority private key file (this shouldn't be shared).
# ca.crt: Certificate Authority trust certificate (this should be shared with users)
# server.key: Server private key, password protected (this shouldn't be shared)
# server.csr: Server certificate signing request (this should be shared with the CA owner)
# server.crt: Server certificate signed by the CA (this would be sent back by the CA owner) - keep on server
# server.pem: Conversion of server.key into a format gRPC likes (this shouldn't be shared)
#
# Private files: ca.key, server.key, server.pem, server.crt
# "Share" files: ca.crt (needed by the client), server.csr (needed by the CA)
#
# Changes these CN's to match your hosts in your environment if needed. Ex: SERVER_CN=myapi.example.com
SERVER_CN=localhost
SSL_PWD=c4pSt0n3Fru1t5
certificates: certs config/ssl.cnf
	openssl genrsa -passout pass:$(SSL_PWD) -des3 -out ./certs/ca.key 4096
	openssl req -passin pass:$(SSL_PWD) -new -x509 -days 3650 -key ./certs/ca.key -out ./certs/ca.crt -subj "/CN=${SERVER_CN}"
	openssl genrsa -passout pass:$(SSL_PWD) -des3 -out ./certs/server.key 4096
	openssl req -passin pass:$(SSL_PWD) -new -key ./certs/server.key -out ./certs/server.csr -subj "/CN=${SERVER_CN}" -config ./config/ssl.cnf
	openssl x509 -req -passin pass:$(SSL_PWD) -days 3650 -in ./certs/server.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key -set_serial 01 -out ./certs/server.crt -extensions req_ext -extfile ./config/ssl.cnf
	openssl pkcs8 -topk8 -nocrypt -passin pass:$(SSL_PWD) -in ./certs/server.key -out ./certs/server.pem
