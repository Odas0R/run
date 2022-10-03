package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// var (
// 	conf = c.Conf{
// 		Id:   "",
// 		Dir:  "/home/odas0r/github.com/odas0r/configs",
// 		File: "config.json",
// 	}
// )

func App() *cli.App {
	app := &cli.App{
		Name:                 "//TODO",
		Usage:                "//TODO",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "hello",
				Action: func(c *cli.Context) error {
					fmt.Println("Hello world!")
					return nil
				},
			},
		},
	}
	return app
}
