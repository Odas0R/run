package cmd

import (
	"log"

	"github.com/odas0r/cmd/pkg/config"
	"github.com/odas0r/cmd/pkg/editor"
	"github.com/odas0r/cmd/pkg/fs"
	"github.com/odas0r/cmd/pkg/shell"
	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
)

var (
	conf = config.Conf{
		Id:   "run-cmd",
		Dir:  "/home/odas0r/github.com/odas0r/configs",
		File: "config.json",
	}
)

func init() {
	if err := conf.Init(); err != nil {
		log.Fatal(err)
	}
}

func App() *cli.App {
	app := &cli.App{
		Name:                 "Run a command on a given path",
		EnableBashCompletion: true,
		Action: func(_ *cli.Context) error {
			dirs := fs.FindDir("/home/odas0r/github.com", 1)

			// filter the directories that have a .git folder in them and print them
			var gitDirs []string
			for _, dir := range dirs {
				if fs.Exists(dir + "/.git") {
					gitDirs = append(gitDirs, dir)
				}
			}

			// spawn the fzf menu
			repoPath := editor.Fzf((gitDirs), "Repos > ")
			if repoPath == "" {
				return nil
			}

			// query the config json for the history key
			history := conf.QueryVal("history")
			if history == nil {
				history = []interface{}{}
			}

			// convert the history to a string slice
			var historyStr []string
			for _, h := range history.([]interface{}) {
				historyStr = append(historyStr, h.(string))
			}

			// spawn the fzf menu with the history of executed commands
			output := editor.FzfPrintQuery(historyStr, "Command > ")

			// append string to the historyStr slice but only if it doesn't already
			// contain the string
			if !lo.Contains(historyStr, output) && len(output) > 0 {
				historyStr = append(historyStr, output)
				if err := conf.Set("history", historyStr); err != nil {
					return err
				}
			}

			// execute a bash script on a specific path
			shell.ExecInteractiveWithPath(output, repoPath)

			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "config",
				Action: func(c *cli.Context) error {
					return editor.Edit(conf.Path())
				},
			},
		},
	}
	return app
}
