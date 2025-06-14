MIGRATE_DRV=postgres
POSTGRES_HOST?=0.0.0.0
MIGRATE_DSN=${PG_DSN}
MIGRATE_DIR=./migrations/migrate
GOOSE_BASE_CMD=goose -dir ${MIGRATE_DIR} ${MIGRATE_DRV} "${MIGRATE_DSN}"

migration-up:		## Migrate the DB to the most recent version available
	${GOOSE_BASE_CMD} up

migration-down: 	## Roll back the version by 1
	${GOOSE_BASE_CMD} down

migration-reset:	## Roll back all migrations
	${GOOSE_BASE_CMD} reset

migration-create:
	@read -p "Enter migration name: " MIGRATION_NAME && goose -dir ${MIGRATE_DIR} create $$MIGRATION_NAME sql


.PHONY: migration-up migration-down migration-reset migration-create

build:
	go mod tidy
	go build -o restaurant-app ./cmd/restaurant

run:
	./restaurant-app

test:
	go test ./... -v

docker-build:
	docker build -t restaurant-app .

docker-run:
	docker run -p 8000:8000 restaurant-app

deploy:
	docker compose up -d