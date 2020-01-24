package models

import (
	"database/sql"
	"fmt"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// ConnectionConfig is the parameters to open connection pool
type ConnectionConfig struct {
	User, Password, Protocol, Address, DbName string
	Port                                      int
}

// OpenConnection opens the connection pool
func OpenConnection(cfg ConnectionConfig) (connection *sql.DB, err error) {
	// See https://github.com/Go-SQL-Driver/MySQL/#dsn-data-source-name
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%v)/%s?collation=utf8_general_ci&parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Protocol,
		cfg.Address,
		cfg.Port,
		cfg.DbName,
	)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open connection")
	}
	// See http://go-database-sql.org/connection-pool.html
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Second * 25)
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to ping database connection")
	}
	return db, nil
}
