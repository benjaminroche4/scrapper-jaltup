package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// DRIVERNAME the driver name.
	DRIVERNAME = "mysql"
	// DBURL the database url connect string.
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// .
	DBURL = "%s:%s@(%s:%d)/%s"
)

type DB struct {
	hostname string
	port     int
	username string
	password string
	database string
	timeout  time.Duration
	db       *sql.DB
}

func New(hostname string, port int, username, password, database string) Database {
	return &DB{
		hostname: hostname,
		port:     port,
		username: username,
		password: password,
		database: database,
		timeout:  10 * time.Second,
		db:       nil,
	}
}

func (thiz *DB) Open() error {
	var err error
	connectionString := fmt.Sprintf(DBURL, thiz.username, thiz.password, thiz.hostname, thiz.port, thiz.database)
	thiz.db, err = sql.Open(DRIVERNAME, connectionString)
	if err != nil {
		return fmt.Errorf("DB(Open): %w", err)
	}
	return nil
}

func (thiz *DB) Close() error {
	if err := thiz.db.Close(); err != nil {
		return fmt.Errorf("DB(Close): %w", err)
	}
	return nil
}

func (thiz *DB) Ping() error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelfunc()

	err := thiz.db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("DB(Ping): %w", err)
	}

	return nil
}
