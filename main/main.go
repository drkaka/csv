package main

import (
	"fmt"
	"os"

	"github.com/drkaka/csv"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "csv"
	app.Usage = "command line application to read csv from App Usage"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Specify the input csv file.",
		},
	}

	app.Before = func(c *cli.Context) error {
		fPath := c.GlobalString("f")
		if st, err := os.Stat(fPath); err != nil {
			return err
		} else if st.IsDir() {
			return fmt.Errorf("%s is not a file", fPath)
		}
		return nil
	}

	app.Action = func(c *cli.Context) error {
		fPath := c.GlobalString("f")
		_, err := csv.ParseFile(fPath)
		return err
	}

	app.Run(os.Args)
}
