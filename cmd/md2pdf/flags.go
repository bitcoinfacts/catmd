package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

type opts struct {
	orderFile string
	pathDir   string
	out       string
}

func validateFlags(c *cli.Context) (opts, error) {
	userOpts := opts{
		orderFile: c.String(orderFileFlagName),
		pathDir:   c.String(pathDirFlagName),
		out:       c.String(outFlagName),
	}

	errorMsg := "Usage: md2pdf --%s <path>"
	if userOpts.orderFile == "" {
		return opts{}, fmt.Errorf(errorMsg, orderFileFlagName)
	} else if userOpts.pathDir == "" {
		return opts{}, fmt.Errorf(errorMsg, pathDirFlagName)
	} else if userOpts.out == "" {
		return opts{}, fmt.Errorf(errorMsg, outFlagName)
	}
	return userOpts, nil
}
