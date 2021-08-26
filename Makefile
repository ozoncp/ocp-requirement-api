.PHONY: build
build: vendor-proto .generate .build .generate_mocks

.PHONY: .generate
.generate:
		mkdir -p swagger
		mkdir -p pkg/ocp-requirement-api
		protoc -I vendor.protogen \
				--go_out=pkg/ocp-requirement-api --go_opt=paths=import \
				--go-grpc_out=pkg/ocp-requirement-api --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/ocp-requirement-api \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/ocp-requirement-api \
				--swagger_out=allow_merge=true,merge_file_name=api:swagger \
				api/ocp-requirement-api/ocp-requirement-api.proto
		mv pkg/ocp-requirement-api/github.com/ozoncp/ocp-requirement-api/pkg/ocp-requirement-api/* pkg/ocp-requirement-api/
		rm -rf pkg/ocp-requirement-api/github.com
		mkdir -p cmd/ocp-requirement-api

.PHONY: .build
.build:
		CGO_ENABLED=0 GOOS=linux go build -o bin/ocp-requirement-api cmd/ocp-requirement-api/main.go

.PHONY: install
install: build .install

.PHONY: .install
install:
		go install cmd/grpc-server/main.go

.PHONY: vendor-proto
vendor-proto: .vendor-proto

.PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		mkdir -p vendor.protogen/api/ocp-requirement-api
		cp api/ocp-requirement-api/ocp-requirement-api.proto vendor.protogen/api/ocp-requirement-api
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/github.com/envoyproxy &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
		fi


.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get -u github.com/golang/mock/mockgen
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate
		go install github.com/golang/mock/mockgen

.PHONY: migrate
migrate: .install-go-deps .migrate

.PHONY: .migrate
.migrate:
	 go get -U github.com/pressly/goose/v3/cmd/goose
	 go install github.com/pressly/goose/v3/cmd/goose
	 goose -s -dir ./migrations postgres "postgres://user:password@127.0.0.1:5444/requirement_db?sslmode=disable" up

.PHONY: generate_mocks
generate_mocks: .install-go-deps .generate_mocks

.PHONY: .generate_mocks
.generate_mocks:
	mockgen  -destination ./internal/mocks/repository.go -package mocks github.com/ozoncp/ocp-requirement-api/internal/repository Repo
	mockgen  -destination ./internal/mocks/flusher.go -package mocks github.com/ozoncp/ocp-requirement-api/internal/flusher Flusher

.PHONY: tests
tests: .generate_mocks .tests

.PHONY: .tests
.tests:
	go test ./... -v

.PHONY: up
up: .up

.PHONY: .up
.up:
	docker-compose -f docker-compose.only-env.yaml -f docker-compose.app.yaml down
	docker-compose -f docker-compose.only-env.yaml -f docker-compose.app.yaml up --build

.PHONY: up-env
up-env: .up-env

.PHONY: .up-env
.up-env:
	docker-compose -f docker-compose.only-env.yaml down
	docker-compose -f docker-compose.only-env.yaml up
