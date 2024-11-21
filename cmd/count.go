package main

import (
	"fmt"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/urfave/cli/v2"
)

func count(c *cli.Context) error {
	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	nbCategories, err := db.CountCategories()
	if err != nil {
		return err
	}
	nbCompanies, err := db.CountCompanies()
	if err != nil {
		return err
	}
	nbOffers, err := db.CountOffers()
	if err != nil {
		return err
	}

	table := simpletable.New()
	table.SetStyle(simpletable.StyleCompactLite)
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "TABLE"},
			{Align: simpletable.AlignCenter, Text: "#ROWS"},
		},
	}
	table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "categories"},
		{Align: simpletable.AlignRight, Text: strconv.Itoa(nbCategories)},
	}, []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "companies"},
		{Align: simpletable.AlignRight, Text: strconv.Itoa(nbCompanies)},
	}, []*simpletable.Cell{
		{Align: simpletable.AlignLeft, Text: "offers"},
		{Align: simpletable.AlignRight, Text: strconv.Itoa(nbOffers)},
	})

	fmt.Println(table.String()) // nolint: forbidigo

	return nil
}
