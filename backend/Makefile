postgres:
	docker-compose up -d

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/finance?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

.PHONY: postgres migrateup migrationdrop test server sqlc-gen