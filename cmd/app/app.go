package app

import (
	"fmt"
	"os"

	"github.com/d7561985/image-packer/spriteshits/custom"
	"github.com/urfave/cli/v2"
)

const (
	parameterFile       = "file"
	parameterManifest   = "manifest"
	parameterOutputPath = "output"
)

var generalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     parameterFile,
		Usage:    "image file",
		Aliases:  []string{"f"},
		Required: true,
	},
	&cli.StringFlag{
		Name:     parameterManifest,
		Usage:    "file contains information about how we should treat image",
		Aliases:  []string{"m"},
		Required: true,
	},
	&cli.StringFlag{
		Name:     parameterOutputPath,
		Usage:    "output directory where would be saved splited files",
		Aliases:  []string{"o"},
		Value:    "", //current directory
		Required: false,
	},
}

func split(ctx *cli.Context) error {
	img, err := os.OpenFile(ctx.String(parameterFile), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file %q error: %w", ctx.String(parameterFile), err)
	}

	defer img.Close()

	man, err := os.OpenFile(ctx.String(parameterManifest), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("open file %q error: %w", ctx.String(parameterManifest), err)
	}

	defer man.Close()

	worker, err := custom.New(img, man, ctx.String(parameterOutputPath))
	if err != nil {
		return fmt.Errorf("initiate processor: %w", err)
	}

	return worker.Process()
}
