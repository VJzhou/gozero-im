package orm

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  int
}

type DB struct {
	*gorm.DB
}

func NewMySQL(conf *Config) (*DB, error) {
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}

	if conf.MaxOpenConns == 0 {
		conf.MaxIdleConns = 100
	}

	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 3600
	}

	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: &customLog{},
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDb.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(conf.MaxLifetime) * time.Second)

	err = db.Use(NewCustomePlugin())
	if err != nil {
		panic(err)
	}

	return &DB{DB: db}, nil
}

func MustNewMySQL(conf *Config) *DB {
	db, err := NewMySQL(conf)
	if err != nil {
		panic(err)
	}

	return db
}
