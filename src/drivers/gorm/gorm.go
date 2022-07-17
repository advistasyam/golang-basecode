package gorm

import (
	"fmt"
	"os"
	"time"

	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	models "alteacare/golang-basecode/src/drivers/gorm/models"
)

type Postgres struct {
	host                          string
	port                          string
	user                          string
	password                      string
	database                      string
	logMode                       int
	maxIdleConnection             int
	maxOpenConnection             int
	connectionMaxLifetimeInSecond int
}
type postgresOption func(*Postgres)

func Connect() (*gorm.DB, error) {
	logMode := 0
	if os.Getenv("DB_DEBUG") == "TRUE" {
		logMode = 3
	}

	dbOptions := &Postgres{
		host:                          os.Getenv("DB_HOST"),
		port:                          os.Getenv("DB_PORT"),
		user:                          os.Getenv("DB_USER"),
		password:                      os.Getenv("DB_PASSWORD"),
		database:                      os.Getenv("DB_DATABASE"),
		logMode:                       logMode,
		maxIdleConnection:             5,
		maxOpenConnection:             10,
		connectionMaxLifetimeInSecond: 60,
	}

	return connect(dbOptions)
}

func connect(param *Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		param.host, param.user, param.password, param.database, param.port)

	db, err := gorm.Open(gormPostgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(param.logMode)),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	// set configuration pooling connection
	postgresDB, _ := db.DB()
	postgresDB.SetMaxOpenConns(param.maxOpenConnection)
	postgresDB.SetConnMaxLifetime(time.Duration(param.connectionMaxLifetimeInSecond) * time.Minute)
	postgresDB.SetMaxIdleConns(param.maxIdleConnection)

	migrateAllTables(db)

	return db, nil
}

func SetMaxIdleConns(conns int) postgresOption {
	return func(c *Postgres) {
		if conns > 0 {
			c.maxIdleConnection = conns
		}
	}
}

func SetMaxOpenConns(conns int) postgresOption {
	return func(c *Postgres) {
		if conns > 0 {
			c.maxOpenConnection = conns
		}
	}
}

func SetConnMaxLifetime(conns int) postgresOption {
	return func(c *Postgres) {
		if conns > 0 {
			c.connectionMaxLifetimeInSecond = conns
		}
	}
}

func migrateAllTables(db *gorm.DB) {
	models.MigrateUser(db)
}
