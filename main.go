package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

func main() {
	app := &cli.App{
		Name:  "shrink",
		Usage: "introduction shrink",
		Action: func(path *cli.Context) error {
			if path.Args().Get(0) == "" {
				fmt.Println("this is shrink, a cli for shrinking images")
			} else {
				if _, err := os.Stat("/path/to/whatever"); err == nil {
					fmt.Println("output dir exists")
				} else if errors.Is(err, os.ErrNotExist) {
					os.Mkdir("output", os.ModePerm)
				} else {
					fmt.Println("bro what?")
				}
				arguments := path.Args().Slice()
				for _, s := range arguments {
					cmd := exec.Command("ruby", "tiny.rb", s)
					fmt.Printf("running...%q", cmd)
					fmt.Println()
					err := cmd.Run()
					if err != nil {
						log.Fatal(err)
					}
				}
				fmt.Println("optimization complete")
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
