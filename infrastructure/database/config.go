package database

import "github.com/jmoiron/sqlx"

type DatabaseConfig struct {
	Dialect  string
	Host     string
	Name     string
	Username string
	Password string
	Port     string

	SetMaxIdleConn string
	SetMaxOpenConn string
	SetMaxIdleTime string
	SetMaxLifeTime string
}

type Database struct {
	*sqlx.DB
}
