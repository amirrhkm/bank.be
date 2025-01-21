dbup:
	docker-compose up -d

dbdown:
	docker-compose down -v

dbcreate:
	docker exec -it bank-db createdb --username=root --owner=root bank

dbdrop:
	docker exec -it bank-db dropdb bank

.PHONY: dbup dbdown createdb dropdb