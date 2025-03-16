package postgres

import (
	"fmt"
	config "github.com/dnevsky/restaurant-back/internal/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func NewPostgres() (*gorm.DB, error) {
	dbLogger := gormLogger.Default.LogMode(gormLogger.Silent)
	if config.Config.Env == config.EnvDev {
		dbLogger = gormLogger.Default.LogMode(gormLogger.Info)
	}
	pg := postgres.New(postgres.Config{
		PreferSimpleProtocol: true,
		DSN:                  config.Config.PgDsn,
	})
	ormDB, err := gorm.Open(pg, &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              false,
	})
	if err != nil {
		return nil, err
	}

	ormDB.Statement.RaiseErrorOnNotFound = true

	return ormDB, nil
}

func CloseDB(ormDB *gorm.DB) {
	db, err := ormDB.DB()
	if err != nil {
		log.Fatalf(fmt.Sprintf("cant close the connector: %s", err))
	}

	err = db.Close()
	if err != nil {
		log.Fatalf(fmt.Sprintf("cant close the connector: %s", err))
	}
}
