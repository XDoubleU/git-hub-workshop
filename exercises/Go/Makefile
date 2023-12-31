db ?= postgres://postgres@localhost/postgres

init:
	go install github.com/pressly/goose/v3/cmd/goose@v3.15.0
	go install github.com/segmentio/golines@v0.11.0
	go install github.com/daixiang0/gci@v0.11.2
	go install github.com/securego/gosec/v2/cmd/gosec@v2.17.0
	go install github.com/swaggo/swag/cmd/swag@v1.16.2

lint:
	golangci-lint run
	sqlfluff lint --dialect postgres ./migrations

lint/fix:
	swag fmt
	golines . -m 88 -w
	golangci-lint run --fix
	gci write --skip-generated -s standard -s default -s "prefix(check-in/api)" .
	sqlfluff fix --dialect postgres ./migrations

build:
	go build -o=./bin/api ./cmd/api

run/api:
	go run ./cmd/api

test:
	go test ./cmd/api

test/v:
	go test -v ./cmd/api

test/cov/report:
	go test -covermode=set -coverprofile=coverage.out ./cmd/api -coverpkg ./cmd/api,./internal/...

test/cov:
	go test -covermode=set -coverprofile=coverage.out ./cmd/api -coverpkg ./cmd/api,./internal/...
	go tool cover -html=coverage.out -o=coverage.html
	make test/cov/open

test/cov/open:
	CMD /C start chrome /new-tab %CD%/coverage.html

db/migrations/new:
	goose -s -dir=./migrations create ${name} sql

db/migrations/up:
	goose -dir=./migrations postgres ${db} up

db/migrations/down:
	goose -dir=./migrations postgres ${db} down
