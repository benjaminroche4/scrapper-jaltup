// nolint:forbidigo
package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

// parseCLI parses command lines arguments.
func parseCLI() error {
	cliapp := cli.NewApp()
	cliapp.Name = "scrapper"
	cliapp.Usage = ""
	cliapp.UsageText = "scrapper <command> [options]"
	cliapp.Description = "Jaltup Scrapper"
	cliapp.Version = version

	cliapp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "host",
			Value:    "localhost",
			Usage:    "database host name",
			Aliases:  []string{"H"},
			Required: false,
			EnvVars:  []string{"DATABASE_HOST"},
		},
		&cli.IntFlag{
			Name:     "port",
			Value:    3306,
			Usage:    "database port number",
			Aliases:  []string{"P"},
			Required: false,
			EnvVars:  []string{"DATABASE_PORT"},
		},
		&cli.StringFlag{
			Name:     "user",
			Value:    "jaltup",
			Usage:    "database username",
			Aliases:  []string{"u"},
			Required: false,
			EnvVars:  []string{"DATABASE_USERNAME"},
		},
		&cli.StringFlag{
			Name:     "pass",
			Value:    "",
			Usage:    "database password",
			Aliases:  []string{"p"},
			Required: false,
			EnvVars:  []string{"DATABASE_PASSWORD"},
		},
		&cli.StringFlag{
			Name:     "dbname",
			Value:    "jaltup",
			Usage:    "database name",
			Aliases:  []string{},
			Required: false,
			EnvVars:  []string{"DATABASE_NAME"},
		},
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println("Version:\t", c.App.Version)
	}

	cliapp.Commands = []*cli.Command{
		cmdSourceLba(cliapp.Flags),
		cmdSourceAltPro(cliapp.Flags),
		cmdCount(cliapp.Flags),
		cmdSelectCategories(cliapp.Flags),
		cmdSelectCompanies(cliapp.Flags),
		cmdSelectOffers(cliapp.Flags),
		cmdClean(cliapp.Flags),
	}

	sort.Sort(cli.FlagsByName(cliapp.Flags))
	sort.Sort(cli.CommandsByName(cliapp.Commands))

	if err := cliapp.Run(os.Args); err != nil {
		return fmt.Errorf("failed to parse command line arguments: %w", err)
	}
	return nil
}

func cmdSourceLba(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "lba",
		Usage:     "fetch all offers from 'la bonne alternance' source",
		UsageText: "scrapper lba <options>",
		Action:    source,
		Flags:     flags,
	}
}

func cmdSourceAltPro(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "altpro",
		Usage:     "fetch all offers from 'alternance professionnelle' source",
		UsageText: "scrapper altpro <options>",
		Action:    source,
		Flags:     flags,
	}
}

func cmdCount(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "count",
		Usage:     "count rows for all tables",
		UsageText: "scrapper count <options>",
		Action:    count,
		Flags:     flags,
	}
}

func cmdSelectCategories(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "categories",
		Usage:     "retrieve all categories from database",
		UsageText: "scrapper categories <options>",
		Action:    selectCategories,
		Flags:     flags,
	}
}

func cmdSelectCompanies(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "companies",
		Usage:     "retrieve all companies from database",
		UsageText: "scrapper companies <options>",
		Action:    selectCompanies,
		Flags:     flags,
	}
}

func cmdSelectOffers(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "offers",
		Usage:     "retrieve all offers from database",
		UsageText: "scrapper offers <options>",
		Action:    selectOffers,
		Flags:     flags,
	}
}

func cmdClean(flags []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:      "clean",
		Usage:     "clean(empty) all tables, this action is irreversible",
		UsageText: "scrapper clean <options>",
		Action:    clean,
		Flags:     flags,
	}
}
