package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

//go:embed tiny.rb
var tinyScript string

var monthCompressionCount int

func main() {
	app := &cli.App{
		//Commands: []*cli.Command{
		//	{
		//		Name: "",
		//	},
		//:wq:wq},
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

				arguments := path.Args().Slice()

				for _, s := range arguments {
					cmd := exec.Command("ruby", "tiny.rb", s)
					fmt.Printf("running...%q", cmd)
					fmt.Println()
					err := cmd.Run()
					checkErr(err)
				}
				fmt.Println("optimization complete")

				content, err := os.ReadFile("count.txt")
				checkErr(err)

				data := string(content)

				fmt.Println("data is: " + data)

				dataInt, er := strconv.Atoi(data)
				checkErr(er)
				checkCount(dataInt)
				countLeft := 500 - dataInt
				countLeftString := strconv.Itoa(countLeft)

				fmt.Println("Used count: " + data + " Count left for this month: " + countLeftString)
				monthCompressionCount = dataInt

				e := os.Remove("tiny.rb")
				checkErr(e)

				e2 := os.Remove("count.txt")
				checkErr(e2)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func checkCount(data int) {
	if data == 500 {
		email := RegisterEmail()
		fmt.Println(email)
	}
}

func RegisterEmail() string {
	url := "https://privatix-temp-mail-v1.p.rapidapi.com/request/mail/id/null/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "792b0e438fmshfb9971aecd4728bp16d235jsn45bfd672c1be")
	req.Header.Add("X-RapidAPI-Host", "privatix-temp-mail-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	return string(body)

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
