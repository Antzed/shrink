package main

import (
	"bufio"
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

				if _, err := os.Stat("credentials.txt"); errors.Is(err, os.ErrNotExist) {
					APIFetch()
				}

				apiKey := readKey("credentials.txt")

				tinyScript := writeRuby(apiKey)

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

				data = strings.TrimRight(data, "\n")

				dataInt, er := strconv.Atoi(data)
				checkErr(er)

				countLeft := 500 - dataInt
				countLeftString := strconv.Itoa(countLeft)
				fmt.Println("data is: " + data)
				fmt.Println("Used count: " + data + " Count left for this month: " + countLeftString)
				//monthCompressionCount = dataInt

				checkCount(dataInt)

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

func readKey(fileName string) string {
	file, err := os.Open(fileName)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var key string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "key:") {
			line = strings.ReplaceAll(line, "key:", "")
			key = line
		}
	}
	err2 := scanner.Err()
	checkErr(err2)

	return key
}

func checkCount(data int) {
	if data == 500 {
		fmt.Println("quota fully used, fetching new quota...")
		APIFetch()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

	compressions_this_month = Tinify.compression_count
	File.write('./count.txt', compressions_this_month)`

	return tinyScript
}
