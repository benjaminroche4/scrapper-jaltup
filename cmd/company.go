package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/jftuga/ellipsis"
	"github.com/urfave/cli/v2"
)

func selectCompanies(c *cli.Context) error {
	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	companies, err := db.SelectCompanies()
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
			{Align: simpletable.AlignCenter, Text: "SIRET"},
			{Align: simpletable.AlignCenter, Text: "VERFIIED"},
			{Align: simpletable.AlignCenter, Text: "CREATED"},
		},
	}

	for i := range companies {
		company := companies[i]
		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: strconv.FormatInt(company.ID, 10)},
			{Align: simpletable.AlignLeft, Text: company.PublicID},
			{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(company.Name, 40)},
			{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(company.Slug, 40)},
			{Align: simpletable.AlignLeft, Text: company.Siret},
			{Align: simpletable.AlignLeft, Text: strconv.FormatBool(company.Verified)},
			{Align: simpletable.AlignLeft, Text: company.CreatedAt.Format(time.DateTime)},
		})
	}

	fmt.Println(table.String()) // nolint: forbidigo

	return nil
}
