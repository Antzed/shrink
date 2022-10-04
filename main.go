package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
)

//go:embed tiny.rb
var tinyScript string

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

				f, err := os.Create("tiny.rb")
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				_, err2 := f.WriteString(tinyScript)
				if err2 != nil {
					log.Fatal(err2)
				}
				fmt.Println("done")
				//dir, err := os.Getwd()
				//if err != nil {
				//	log.Println(err)
				//}
				//fmt.Printf("current directory is %q", dir)
				//fmt.Println()
				//tinyPath := dir + "/tiny.rb"

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
				e := os.Remove("tiny.rb")
				if e != nil {
					log.Fatal(e)
				}
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
