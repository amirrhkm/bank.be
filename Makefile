dbup:
	docker-compose up -d

dbdown:
	docker-compose down -v

dbcreate:
	docker exec -it bank-db createdb --username=root --owner=root bank

dbdrop:
	docker exec -it bank-db dropdb bank

dbmigrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank?sslmode=disable" -verbose up

dbmigrateup1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank?sslmode=disable" -verbose up 1

dbmigratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank?sslmode=disable" -verbose down

dbmigratedown1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank?sslmode=disable" -verbose down 1

testdbcreate:
	docker exec -it bank-db createdb --username=root --owner=root bank-test

testdbdrop:
	docker exec -it bank-db dropdb bank-test

testdbmigrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank-test?sslmode=disable" -verbose up

testdbmigratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank-test?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

testreport:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/amirrhkm/bank.be/db/sqlc Store

.PHONY: dbup dbdown createdb dropdb dbmigrateup dbmigrateup1 dbmigratedown dbmigratedown1 sqlc test testcoverage server mock
