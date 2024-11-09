package config 

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const baseDSN = "%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4"

func createDatabaseConnection(cfg *Database) (*gorm.DB, error) {
	mysqlConfig := mysql.Config{
		DSN: buildConnectionString(cfg),
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		return nil, err
	}

	db.Logger = createLogger()

	return db, nil
}

func buildConnectionString(cfg *Database) string {
	return fmt.Sprintf(baseDSN, cfg.Username, cfg.Password, cfg.Hostname, cfg.Schema)
}

func createLogger() logger.Interface {
	rawLogger := log.New(os.Stdout, "\n", log.LstdFlags)

	cfg := logger.Config{
		Colorful:                  false,
		LogLevel:                  logger.Warn,
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: true,
	}

	return logger.New(rawLogger, cfg)
}