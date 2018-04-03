package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

const (
	orderFileFlagName = "order-file"
	pathDirFlagName   = "path-dir"
	outFlagName       = "out"
)

func main() {
	app := cli.NewApp()
	app.Name = "md2pdf"
	app.Usage = "Concatenates multiple jekyll sources into a standard markdown file"
	app.UsageText = "md2pdf --order-file docs_toc.yml --base-dir _docs --out bin/glacier.md"
	app.Action = runAction
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  orderFileFlagName,
			Usage: "yaml file with import order",
		},
		cli.StringFlag{
			Name:  pathDirFlagName,
			Usage: "the directory with markdown files",
		},
		cli.StringFlag{
			Name:  outFlagName,
			Usage: "where to store the resulting markdown file",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runAction(c *cli.Context) error {
	userOpts, err := validateFlags(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SUCCESS: %v", userOpts)

	return nil
}
