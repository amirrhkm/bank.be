

### Database
-> DBML
-> Postgres
-> SQLC

## API
-> Gin

## Config management
-> Viper

## CICD
-> Github Actions

## Unit Testing
-> testify
-> gomock

run test
`make test`

run test with report coverage
`make testreport`

run server
`make server`

generate model and type-safe queries with boilerplate
`make sqlc`

create new migration
`migrate create -ext sql -dir db/migration -seq <migration_name>`

```
Deadlock
-> prioritise ACID principles
-> query order
-> transaction isolation level
```
