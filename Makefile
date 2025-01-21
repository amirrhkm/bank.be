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

sqlc:
	sqlc generate

.PHONY: dbup dbdown createdb dropdb dbmigrateup dbmigratedown sqlc
