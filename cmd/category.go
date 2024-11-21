package main

import (
	"fmt"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/urfave/cli/v2"
)

func selectCategories(c *cli.Context) error {
	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	categories, err := db.SelectCategories()
	if err != nil {
		return err
	}

	table := simpletable.New()
	table.SetStyle(simpletable.StyleCompactLite)
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "PUBLICID"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "SLUG"},
		},
	}

	for _, category := range categories {
		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: strconv.FormatInt(category.ID, 10)},
			{Align: simpletable.AlignLeft, Text: category.PublicID},
			{Align: simpletable.AlignLeft, Text: category.Name},
			{Align: simpletable.AlignLeft, Text: category.Slug},
		})
	}

	fmt.Println(table.String()) // nolint: forbidigo

	return nil
}
