createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/must-have" -verbose up

migratedown:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/must-have" -verbose down

wire:
	wire ./...

sqlc:
	sqlc generate

migrate-sqlc:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/must-have" -verbose up
	sqlc generate

swagger:
	swag init -g cmd/server/main.go

.PHONY: migrate migratedown createmigration wire sqlc migrate-sqlc