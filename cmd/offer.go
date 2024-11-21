package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/jftuga/ellipsis"
	"github.com/urfave/cli/v2"
)

func selectOffers(c *cli.Context) error {
	db, err := openDB(c)
	if err != nil {
		return err
	}
	defer db.Close()

	offers, err := db.SelectOffers()
	if err != nil {
		return err
	}

	table := simpletable.New()
	table.SetStyle(simpletable.StyleCompactLite)
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "PUBLICID"},
			{Align: simpletable.AlignCenter, Text: "TITLE"},
			{Align: simpletable.AlignCenter, Text: "SLUG"},
			{Align: simpletable.AlignCenter, Text: "URL"},
			{Align: simpletable.AlignCenter, Text: "STATUS"},
			{Align: simpletable.AlignCenter, Text: "PREMIUM"},
			{Align: simpletable.AlignCenter, Text: "CREATED"},
			{Align: simpletable.AlignCenter, Text: "ENDAT"},
			{Align: simpletable.AlignCenter, Text: "SERVICE"},
			{Align: simpletable.AlignCenter, Text: "EXTERNALID"},
			{Align: simpletable.AlignCenter, Text: "CATEGORIES"},
		},
	}

	for i := range offers {
		offer := offers[i]
		categories := []string{}
		for _, category := range offer.Categories {
			categories = append(categories, category.Name)
		}

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: strconv.FormatInt(offer.ID, 10)},
			{Align: simpletable.AlignLeft, Text: offer.PublicID},
			{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(offer.Title, 30)},
			{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(offer.Slug, 30)},
			{Align: simpletable.AlignLeft, Text: ellipsis.Shorten(offer.URL, 30)},
			{Align: simpletable.AlignLeft, Text: offer.Status},
			{Align: simpletable.AlignLeft, Text: strconv.FormatBool(offer.Premium)},
			{Align: simpletable.AlignLeft, Text: offer.CreatedAt.Format(time.DateTime)},
			{Align: simpletable.AlignLeft, Text: offer.EndAt.Format(time.DateTime)},
			{Align: simpletable.AlignLeft, Text: offer.ServiceName},
			{Align: simpletable.AlignLeft, Text: offer.ExternalID},
			{Align: simpletable.AlignLeft, Text: strings.Join(categories, ",")},
		})
	}

	fmt.Println(table.String()) // nolint: forbidigo

	return nil
}
