// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/urfave/cli/v2"
// )

// func main() {
// 	app := &cli.App{
// 		Name:  "hello",
// 		Usage: "greet the user",
// 		Action: func(c *cli.Context) error {
// 			fmt.Println(`
//                  |
//                 |.|
//                |\./|
//                |\./|
// .              |\./|              .
// \^.\          |\\.//|          /.^/
//  \--.|\       |\\.//|       /|.--/
//    \--.| \    |\\.//|    / |.--/
//     \---.|\    |\./|    /|.---/
//        \--.|\  |\./|  /|.--/
//           \ .\  |.|  /. /
// _ -_^_^_^_-  \ \\ // /  -_^_^_^_- _
//   - -/_/_/- ^ ^  |  ^ ^ -\_\_\- -
//             /-./ | \.-\
//            /-/   |   \-\
// 	  ;|'         '|;
// -=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
// 	  OpenTerpsCLI
// 			`)
// 			return nil
// 		},
// 	}

// 	err := app.Run(os.Args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "info",
				Usage: "greet the user",
				Action: func(c *cli.Context) error {
					fmt.Println(`
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
	  OpenTerpsCLI
			`)
					return nil
				},
			},
			{
				Name:  "strain",
				Usage: "looks up cannabis strain by name",
				Action: func(c *cli.Context) error {
					fmt.Println("look up strain name: ", c.Args().First())
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
