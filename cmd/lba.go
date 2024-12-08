package main

import (
	"scrapperjaltup/matcher"
	"scrapperjaltup/source"

	"github.com/urfave/cli/v2"
)

func sourceLba(c *cli.Context) error {
	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	source := source.NewLBA()

	m := matcher.New(db, source)

	err = m.Execute()
	if err != nil {
		return err
	}

	return nil
}
