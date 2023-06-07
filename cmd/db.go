package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	user     string
	password string
	host     string
	port     string
	database string
}

func NewConfig(user, password, host, port, database string) *Config {
	return &Config{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

func InitDB(c *Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.password, c.host, c.port, c.database)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect server failed, err:%v\n", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("connect server failed, err:%v\n", err)
	}

	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)
	return db, nil
}
