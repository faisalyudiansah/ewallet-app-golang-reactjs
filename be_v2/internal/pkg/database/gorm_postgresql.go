package database

import (
	"fmt"
	"time"

	"ewallet-server-v2/internal/config"
	"ewallet-server-v2/internal/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// in this file, we initialize the Gorm database definition
func InitGorm(cfg *config.Config) *gorm.DB {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Port,
		dbCfg.Sslmode,
	)

	logCfg := gormlogger.Config{
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      false,
		Colorful:                  true,
		LogLevel:                  gormlogger.Info,
	}

	newLogger := gormlogger.New(
		logger.Log,
		logCfg,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Log.Fatal("error initializing database: ", err.Error())
	}

	dbinstance, err := db.DB()
	if err != nil {
		logger.Log.Fatal("error getting generic database instance: ", err.Error())
	}

	dbinstance.SetMaxIdleConns(dbCfg.MaxIdleConn)
	dbinstance.SetMaxOpenConns(dbCfg.MaxOpenConn)
	dbinstance.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetimeMinute) * time.Minute)

	return db
}
