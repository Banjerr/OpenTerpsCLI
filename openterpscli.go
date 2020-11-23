package main

import (
	"fmt"
	"log"
	"openterpscli/strain"
	"os"

	"github.com/urfave/cli/v2"
)

// BaseURL - the main / of the API
const BaseURL = `https://open-terps.herokuapp.com/`

func main() {
	cli.AppHelpTemplate = `
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
                 |
                |.|
               |\./|
               |\./|
.              |\./|              .
\^.\          |\\.//|          /.^/
 \--.|\       |\\.//|       /|.--/
   \--.| \    |\\.//|    / |.--/
    \---.|\    |\./|    /|.---/
       \--.|\  |\./|  /|.--/
          \ .\  |.|  /. /
_ -_^_^_^_-  \ \\ // /  -_^_^_^_- _
  - -/_/_/- ^ ^  |  ^ ^ -\_\_\- -
            /-./ | \.-\
           /-/   |   \-\
	        ;|'         '|;
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
 _____  ____  ____  _  _  ____  ____  ____  ____  ___     ___  __    ____ 
(  _  )(  _ \( ___)( \( )(_  _)( ___)(  _ \(  _ \/ __)   / __)(  )  (_  _)
 )(_)(  )___/ )__)  )  (   )(   )__)  )   / )___/\__ \  ( (__  )(__  _)(_ 
(_____)(__)  (____)(_)\_) (__) (____)(_)\_)(__)  (___/   \___)(____)(____)
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
WEBSITE: https://open-terps.herokuapp.com/
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
SUPPORT: benjamminredden@gmail.com
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "API-User",
				Aliases: []string{"u"},
				Value:   "your-open-terps-api-user",
				Usage:   "API User for the OpenTeprs API",
				EnvVars: []string{"APIUser"},
			},
			&cli.StringFlag{
				Name:    "API-Key",
				Aliases: []string{"p"},
				Value:   "your-open-terps-api-key",
				Usage:   "API Key for the OpenTeprs API",
				EnvVars: []string{"APIKey"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "info",
				Usage: "provides useful informational about the OpenTerps CLI 🍁",
				Action: func(c *cli.Context) error {
					fmt.Println(``)
					return nil
				},
			},
			{
				Name:  "strain",
				Usage: "options for strains",
				Subcommands: []*cli.Command{
					{
						Name:  "get",
						Usage: "get a(ll) strain(s)",
						Action: func(c *cli.Context) error {
							strain.GetStrains(BaseURL)
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "add a new strain",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing strain",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "update an existing strain",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
