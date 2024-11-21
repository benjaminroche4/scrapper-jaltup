package main

import (
	_db "scrapperjaltup/db"

	"github.com/urfave/cli/v2"
)

func openDB(c *cli.Context) (_db.Database, error) {
	host := c.String("host")
	port := c.Int("port")
	user := c.String("user")
	pass := c.String("pass")
	dbname := c.String("dbname")

	db := _db.New(host, port, user, pass, dbname)
	err := db.Open()
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
