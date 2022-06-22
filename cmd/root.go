package cmd

import (
	"bufio"
	"errors"
	"os"

	"github.com/sbvalois/gojira/issues"
	"github.com/sbvalois/gojira/setup"
	"github.com/sbvalois/gojira/transitions"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// Contains all cli commands
func CreateCliApp(filename string) *cli.App {
	return &cli.App{
		Name:  "Gojira",
		Usage: "Get your jira tasks, right in your terminal",
		Commands: []cli.Command{
			{
				Name:  "setup",
				Usage: "Setup your jira integration",
				Action: func(c *cli.Context) error {
					setup.RunBasicSetup(filename, bufio.NewReader(os.Stdin))
					return nil
				},
			},
			{
				Name:    "issues",
				Usage:   "Get issues",
				Aliases: []string{"i"},
				Subcommands: []cli.Command{
					{
						Name:  "open",
						Usage: "Gets all issues with status todo",
						Action: func(c *cli.Context) error {
							issues.GetIssues("open")
							return nil
						},
					},
					{
						Name:  "inprogress",
						Usage: "Gets all issues with status in progress",
						Action: func(c *cli.Context) error {
							issues.GetIssues("inprogress")
							return nil
						},
					},
				},
			},
			{
				Name:  "issue",
				Usage: "Get one issue",
				Action: func(c *cli.Context) error {
					if c.NArg() == 0 {
						return errors.New("No issue key provided")
					}
					issues.GetIssue(c.Args().First())
					return nil
				},
				Subcommands: []cli.Command{
					{
						Name:  "start",
						Usage: "sets an issue to in progress",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							transitions.Set(c.Args().First(), viper.GetInt("transitions.inProgress"), "IN PROGRESS")
							return nil
						},
					},
					{
						Name:  "open",
						Usage: "sets an issue to open",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							transitions.Set(c.Args().First(), viper.GetInt("transitions.open"), "OPEN")
							return nil
						},
					},
					{
						Name:  "review",
						Usage: "Sets an issue to review. This will only work if the issue is already in progress",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							transitions.Set(c.Args().First(), viper.GetInt("transitions.review"), "REVIEW")
							return nil
						},
					},
					{
						Name:  "done",
						Usage: "sets an issue to done. This will only work if the issue is already in progress",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							transitions.Set(c.Args().First(), viper.GetInt("transitions.done"), "DONE")
							return nil
						},
					},
					{
						Name:  "browser",
						Usage: "Opens the jira task in your default browser",
						Action: func(c *cli.Context) error {
							if c.NArg() == 0 {
								return errors.New("No issue key provided")
							}
							OpenBrowser(c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}
}
