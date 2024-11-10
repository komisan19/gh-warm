package config

import (
	"fmt"

	"github.com/komisan19/gh-warm/warm"
	"github.com/urfave/cli/v2"
)

func Execute() *cli.App {
	app := cli.NewApp()
	config(app)
	app.Commands = menu()
	return app
}

func config(app *cli.App) {
	app.Name = "gh warm"
	app.Usage = "searching Good First Issue"
	app.Version = "1.0.0"
}

func menu() []*cli.Command {
	warm := warm.NewWarm()
	return []*cli.Command{
		{
			Name:    "type",
			Aliases: []string{"t"},
			Usage:   "set language(ex: go, rust, python etc.)",
			Action: func(cCtx *cli.Context) error {
				sout, eout, err := warm.SearchRepo(cCtx.Args().First())
				if err != nil {
					fmt.Println(eout)
					return err
				}
				fmt.Println(sout)
				return nil
			},
		},
	}
}
