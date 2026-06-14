include .env
export

start:
	go run main.go

migrate-prepare:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# n is for number of migration step (up) to apply. If n is not specified, all migrations will be applied.
migrate-up:
	migrate -database $$(grep '^DATABASE_URL=' .env | cut -d= -f2-) -path=migrations up ${n}

# n is for number of migration step (down) to apply. If n is not specified, all migrations will be applied.
migrate-down:
	migrate -database $$(grep '^DATABASE_URL=' .env | cut -d= -f2-) -path=migrations down ${n}