package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"catmd/reader"
	"catmd/writer"
)

const (
	orderFileFlagName = "order-file"
	pathDirFlagName   = "path-dir"
	outFlagName       = "out"
)

func main() {
	app := cli.NewApp()
	app.Name = "catmd"
	app.Usage = "Concatenates multiple jekyll sources into a standard markdown file"
	app.UsageText = "catmd --order-file docs_toc.yml --base-dir _docs --out bin/glacier.md"
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
	book := reader.Read(userOpts.orderFile, userOpts.pathDir)
	writer.Write(book, userOpts.out)
	return nil
}
