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

dbmigratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/bank?sslmode=disable" -verbose down

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

sqlc:
	sqlc generate

.PHONY: dbup dbdown createdb dropdb dbmigrateup dbmigratedown sqlc
