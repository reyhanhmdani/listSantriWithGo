package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigration "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Db(ctx context.Context) (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	//
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
	})
	if err != nil {
		fmt.Println(err)
		panic("Cannot Connect to database")
		//return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	//ping database to make sure connection is established successfully
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	logrus.Info("Connect to Database log")
	return db, err

	// migrate -database "mysql://root:Pastibisa@tcp(localhost:3306)/goSantri" -path database/migrations up
	// migrate create -ext sql -dir database/migrations create_table_wikis
}

func Migrate(db *gorm.DB) error {
	logrus.Info("running database migration")

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	driver, err := mysqlMigration.WithInstance(sqlDB, &mysqlMigration.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"mysql", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if !(err == nil || !errors.Is(err, migrate.ErrNoChange)) {
		logrus.Info("No schema changes to apply")
		return nil
	}

	return err
}
