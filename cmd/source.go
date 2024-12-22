package main

import (
	"fmt"
	"scrapperjaltup/matcher"
	_source "scrapperjaltup/source"

	"github.com/urfave/cli/v2"
)

func source(c *cli.Context) error {
	var src _source.Source

	command := c.Command.Name
	switch command {
	case "lba":
		src = _source.NewLBA()
	case "altpro":
		src = _source.NewAltPro()
	default:
		return fmt.Errorf("invalid enetered command '%s'", command)
	}

	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	m := matcher.New(db, src)

	err = m.Execute()
	if err != nil {
		return err
	}

	return nil
}
