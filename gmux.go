package main

import (
	"fmt"
	"os"

	gmux "github.com/davinche/gmux/cli"
	"github.com/urfave/cli"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "GMux"
	app.Usage = "a tmux sessions manager"
	app.Version = VERSION

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable debug logging",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "new",
			Usage:     "create a new gmux config",
			Action:    gmux.New,
			ArgsUsage: "config_name",
		},
		{
			Name:         "edit",
			Usage:        "edit a gmux config",
			ArgsUsage:    "config_name",
			Action:       gmux.Edit,
			BashComplete: gmux.BashCompleteList,
		},
		{
			Name:         "delete",
			Aliases:      []string{"remove"},
			Usage:        "delete a gmux config",
			ArgsUsage:    "config_name",
			Action:       gmux.Delete,
			BashComplete: gmux.BashCompleteList,
		},
		{
			Name:         "start",
			Usage:        "start a tmux session using a gmux config",
			Action:       gmux.Start,
			ArgsUsage:    "config_name",
			BashComplete: gmux.BashCompleteList,
		},
		{
			Name:         "stop",
			Usage:        "stops a tmux session",
			Description:  "Removes a tmux session by running `tmux kill-session -t sessionname`.",
			ArgsUsage:    "session_name",
			Action:       gmux.Stop,
			BashComplete: gmux.BashCompleteList,
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "lists all available gmux configs",
			Action:  gmux.List,
		},
	}

	// Default action to show the help menu
	app.Action = func(c *cli.Context) error {
		configName := c.Args().First()
		if configName != "" {
			return gmux.Start(c)
		}
		return gmux.ShowHelp(c)
	}
	if err := app.Run(os.Args); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s\n",err.Error()))
	}	
}
