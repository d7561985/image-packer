package app

import (
	"os"

	"log"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:  "Image Packer",
		Usage: "CLI application",

		Commands: []*cli.Command{
			{
				Action:  split,
				Name:    "split",
				Usage:   `split sprite-shit image file according with manifest file`,
				Aliases: []string{"s"},
				Flags:   generalFlags,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
