GOPATH=$(shell go env GOPATH)
SCHEMAS_SRC=$(GOPATH)/src/github.com/marcos-wz/capstone/proto

gen-base:
	@echo "Generating BASE scheme..."
	protoc \
		--proto_path=$(SCHEMAS_SRC)/basepb/ \
		--go_out=$(GOPATH)/src/ \
		base.proto

gen-filter:
	@echo "Generating FILTER scheme..."
	protoc \
		--proto_path=$(SCHEMAS_SRC)/filterpb/ \
		--proto_path=$(GOPATH)/src/ \
		--go_out=$(GOPATH)/src/ \
		filter.proto

gen-loader:
	@echo "Generating LOADER scheme..."
	protoc \
		--proto_path=$(SCHEMAS_SRC)/loaderpb/ \
		--go_out=$(GOPATH)/src/ \
		loader.proto

gen-filtercc:
	@echo "Generating FILTER CC scheme..."
	protoc \
		--proto_path=$(SCHEMAS_SRC)/filterccpb/ \
		--go_out=$(GOPATH)/src/ \
		filtercc.proto

gen-fruit:
	@echo "Generating FRUIT scheme..."
	protoc \
		--go_out=$(SCHEMAS_SRC)/fruitpb/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(SCHEMAS_SRC)/fruitpb/ \
		--go-grpc_opt=paths=source_relative \
		--proto_path=$(GOPATH)/src/ \
		--proto_path=$(SCHEMAS_SRC)/fruitpb/ \
		 fruit.proto

mocks:
	mockery --name=FruitRepo --srcpkg=./internal/service --output=./internal/service/mocks
	mockery --name=FruitService --srcpkg=./internal/server --output=./internal/server/mocks

generate: gen-base gen-filter gen-loader gen-filtercc gen-fruit
