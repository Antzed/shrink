package main

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//var monthCompressionCount int

//var apiKey = []string{"3Jkz9YskWc7gb4qgj6shsh86MHMTKPTK"}

func main() {
	cliFunc()
}

func cliFunc() {
	app := &cli.App{
		Name:  "shrink",
		Usage: "introduction shrink",
		Action: func(path *cli.Context) error {
			if path.Args().Get(0) == "" {
				fmt.Println("this is shrink, a cli for shrinking images")
			} else {
				if _, err := os.Stat("/output"); err == nil {
					fmt.Println("output dir exists")
				} else if errors.Is(err, os.ErrNotExist) {
					os.Mkdir("output", os.ModePerm)
				} else {
					fmt.Println("bro what?")
				}

				if _, err := os.Stat("key.txt"); errors.Is(err, os.ErrNotExist) {
					createKeyFile()
				}

				apiKey, err := os.ReadFile("key.txt")
				checkErr(err)

				apiKeyString := string(apiKey)
				tinyScript := writeRuby(apiKeyString)

				f, err := os.Create("tiny.rb")
				checkErr(err)
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
				data = strings.TrimRight(data, "\n")

				dataInt, er := strconv.Atoi(data)
				checkErr(er)
				checkCount(dataInt)
				countLeft := 500 - dataInt
				countLeftString := strconv.Itoa(countLeft)

				fmt.Println("Used count: " + data + " Count left for this month: " + countLeftString)
				//monthCompressionCount = dataInt

				e := os.Remove("tiny.rb")
				checkErr(e)

				//e2 := os.Remove("count.txt")
				//checkErr(e2)
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
		key := APIFetch()
		updateKeyFile(key)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createKeyFile() {
	f, err := os.Create("key.txt")
	checkErr(err)
	defer f.Close()
	key := APIFetch()
	_, err2 := f.WriteString(key)
	checkErr(err2)
}

func updateKeyFile(key string) {
	f, err := os.Open("key.txt")
	checkErr(err)
	defer f.Close()
	f.Truncate(0)
	_, err2 := f.WriteString(key)
	checkErr(err2)
}

func writeRuby(key string) string {

	var tinyScript string = `gem "tinify"

	require "tinify"
	Tinify.key = "` + key + `"
	puts "verified"

	for arg in ARGV
	source = Tinify.from_file(arg)
	source.to_file("./output/"+arg.chop.chop.chop.chop+"_optimized.png")
	end

	`

	return tinyScript
}

//compressions_this_month = Tinify.compression_count
//File.write('./count.txt', compressions_this_month)
