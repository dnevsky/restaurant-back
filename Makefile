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
	docker build --tag dnevsky/restaurant-app .