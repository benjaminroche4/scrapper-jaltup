package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func clean(c *cli.Context) error {
	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.CleanOffers()
	if err != nil {
		return err
	}

	err = db.CleanCompanies()
	if err != nil {
		return err
	}

	err = db.CleanCategories()
	if err != nil {
		return err
	}

	fmt.Println("done") // nolint: forbidigo

	return nil
}
