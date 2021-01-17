package db

import (
	"committees/config"
	"fmt"
	"gorm.io/driver/mysql"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DB struct {
	Conn   *gorm.DB
	logger *logrus.Logger
}

var dbSession *DB
var once sync.Once
var gormLogger *logrus.Logger

type GormLogger struct{}

func GetDB(logger *logrus.Logger) *DB {
	once.Do(func() {
		gormLogger = logger
		setupDB(logger)
	})

	return dbSession
}

func initModel(dbSession *gorm.DB) {
	dbSession.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func updateTimeStampForUpdateCallback(scope *gorm.DB) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//_ = scope.SetColumn("UpdatedAt", time.Now())
	}
}

func setupDB(logger *logrus.Logger) {
	var err error

	dbConfig := config.GetAppConfig().DBConfig
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)
	conn, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		//Logger: GormLogger{},
	})

	if err != nil {
		logger.Panic(err)
		return
	}

	dbSession = &DB{
		Conn:   conn,
		logger: logger,
	}

	dsn := fmt.Sprintf("mysql://%s", connStr)
	runDBMigrations(dsn, logger)

	initModel(dbSession.Conn)

	return
}

const migrationPath = "file://db/migrations"

func runDBMigrations(dsn string, logger *logrus.Logger) error {
	migration, err := migrate.New(migrationPath, dsn)

	if err != nil {
		logger.Panicf("Unable to connect to database. %v \n", err)
		return err
	}

	defer migration.Close()

	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		logger.Panicf("Unable to run migration. %v \n", err)
		return err
	}

	if err == migrate.ErrNoChange {
		logger.Debug("No new migrations to execute. \n")
	} else {
		logger.Debug("Successfully executed migrations. \n")
	}

	return nil
}

func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		gormLogger.WithFields(
			logrus.Fields{
				"module":  "gorm",
				"type":    "sql",
				"rows":    v[5],
				"src_ref": v[1],
				"values":  v[4],
			},
		).Debug(v[3])
	case "log":
		gormLogger.WithFields(logrus.Fields{
			"module": "gorm",
			"type":   "log",
		}).Info(v[2])
	}
}
