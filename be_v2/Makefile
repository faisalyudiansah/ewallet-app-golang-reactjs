include .env

run-test:
	go run ./cmd/api/

test:
	go test -coverpkg=./... -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | grep total
	rm coverage.out

mock:
	mockery --all --case underscore

build:
	go build -o ./build/main ./cmd/api/main.go

migrateforce:
	@migrate -path ${CURDIR}/internal/database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose force 1

migratedown:
	@migrate -path ${CURDIR}/internal/database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose down

migrateup:
	@migrate -path ${CURDIR}/internal/database/migration/ -database "postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -verbose up

.PHONY: run-rest test mock migrateforce migratedown migrateup