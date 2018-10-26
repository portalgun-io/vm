package main

import (
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli-framework"
	log "github.com/multiverse-os/cli-framework/log"
	color "github.com/multiverse-os/cli-framework/text/color"
	fs "github.com/portalgun-io/vm/fs"
)

func main() {
	// Initializing Flags
	var locale string
	var configPath string

	cmd := cli.New(&cli.CLI{
		Name:    "portalgun",
		Usage:   "A command-line interface for the portalgun hypervisor manager",
		Version: cli.Version{Major: 0, Minor: 1, Patch: 0},
		Commands: []cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "Start the specified virtual machine",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "stop",
				Aliases: []string{"t"},
				Usage:   "Stop the specified virtual machine",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "restart",
				Aliases: []string{"r"},
				Usage:   "Restart the specified virtual machine",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Create a new virtual machine",
				Action: func(c *cli.Context) error {
					c.CLI.Logger.Log(log.INFO, "Creating a new virtual machine...")
					if configPath == "" {
						log.Print(log.INFO, "No configuration defined, using default settings...")
					} else {
						fmt.Println("Configuration specified: " + configPath)
						fmt.Println("First need to check if it exists, then need to check if its parsable/valid YAML.")
						if fs.PathExists(configPath) {
							fmt.Println("File exists, now its time to check if it is valid and contains the ocrrect/necessary data.")
						} else {
							fmt.Println(color.Red("[Error] ") + "Nothing found at specified config path.")
						}
					}
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all virtual machines",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "config",
				Value:       "",
				Usage:       "Path to virtual machine configuration file",
				Destination: &configPath,
			},
			cli.StringFlag{
				Name:        "locale",
				Value:       "en",
				Usage:       "Define localization setting",
				Destination: &locale,
			},
		},
		BashCompletion: true,
		BashComplete: func(c *cli.Context) {
			if c.NArg() > 0 {
				return
			}
			// TODO: here we can iterate over the virtaul machines so we can autocomplete them
		},
	})

	cmd.Run(os.Args)
}
