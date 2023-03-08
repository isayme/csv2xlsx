package main

import (
	"fmt"
	"os"

	"github.com/isayme/csv2xlsx/csv2xlsx"
	"github.com/isayme/go-logger"
	"github.com/urfave/cli/v2"
)

var (
	APP_NAME    string = "csv2xlsx"
	APP_VERSION string = "1.0.0"
)

func main() {
	options := csv2xlsx.Options{}

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "input",
			Usage:       "source csv file",
			Destination: &options.InputFilePath,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "comma",
			Usage:       "csv value separator",
			Value:       ",",
			Destination: &options.Comma,
			Action: func(ctx *cli.Context, s string) error {
				if len(s) > 1 || len(s) == 0 {
					return fmt.Errorf("comma should have one charactor")
				}

				return nil
			},
		},
		&cli.StringFlag{
			Name:        "output",
			Usage:       "destination file",
			Destination: &options.OutputFilePath,
		},
		&cli.StringFlag{
			Name:        "password",
			Usage:       "destination file password",
			Destination: &options.OutputFilePassword,
		},
	}

	app := &cli.App{
		Name:    APP_NAME,
		Version: APP_VERSION,
		Usage:   "convert csv file to xlsx format",
		Flags:   flags,
		Action: func(ctx *cli.Context) error {
			return csv2xlsx.Convert(&options)
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
	}
}
