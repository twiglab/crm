package cmdb

import (
	"database/sql"
	"log"
	"time"
)

type Database struct {
	DSN     string            `json:"dsn" mapstructure:"dsn" yaml:"dsn"`
	Driver  string            `json:"driver" yaml:"driver" mapstructure:"driver"`
	Dialect string            `json:"dialect" yaml:"dialect" mapstructure:"dialect"`
	Configs map[string]string `json:"configs" mapstructure:"configs" yaml:"configs"`
}

func NewDB(c Database) *sql.DB {
	db, err := sql.Open(c.Driver, c.DSN)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(10 * time.Minute)
	return db
}

func OpenDB[C, O any](dbconf Database, newFunc func(Database) *sql.DB, openFunc func(*sql.DB, string, ...O) (C, error), o ...O) C {
	db := newFunc(dbconf)
	c, err := openFunc(db, dbconf.Dialect, o...)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
