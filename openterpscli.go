package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

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
WEBSITE: https://openterps.countryfriedcoders.me
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
				Usage: "looks up cannabis strain by name",
				Action: func(c *cli.Context) error {
					fmt.Println("look up strain name: ", c.Args().First())
					fmt.Printf("getting info on strain: %q", c.Args().First())
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
