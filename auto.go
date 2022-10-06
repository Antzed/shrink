package main

import (
	"fmt"
	"github.com/go-rod/rod"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func registerEmail() string {
	url := "https://privatix-temp-mail-v1.p.rapidapi.com/request/domains/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "792b0e438fmshfb9971aecd4728bp16d235jsn45bfd672c1be")
	req.Header.Add("X-RapidAPI-Host", "privatix-temp-mail-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	bodyClean := strings.Trim(string(body), "[")
	fmt.Println(bodyClean)
	bodyClean = strings.Trim(bodyClean, "]")
	fmt.Println(bodyClean)
	bodyClean = strings.Trim(bodyClean, "\"")
	fmt.Println(bodyClean)

	bodySlice := strings.Split(string(body), "\",\"")

	bodySliceLenAdjusted := len(bodySlice) - 1
	randNum := rand.Intn(bodySliceLenAdjusted)
	domain := bodySlice[randNum]
	name := RandStringRunes(5)
	email := name + domain
	return email

}

func main() {

	//email := registerEmail()
	//md5 := md52.Sum([]byte(email))
	//signUp()
	getKey()

}

func signUp(email string) {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://tinypng.com/developers")

	// We use css selector to get the search input element and input "git"
	page.MustElement("##top > div > div > form > input:nth-child(1)").MustInput("lao ting")
	//!!! need new name every time
	page.MustElement("#top > div > div > form > input[type=email]:nth-child(2)").MustInput(email)

	page.MustElement("#top > div > div > form > input[type=submit]:nth-child(3)").MustClick()

}

func getKey() {
	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://tinify.com/login?token=DM1qPL3nx4R9J5xQj8nsmLc8KgqFhsp6/KGi2zGvtlf__&new=true&redirect=/dashboard/api")
	//!!! need the original link from the email
	fmt.Println("got page")
	// We use css selector to get the search input element and input "git"
	//page.MustElement("#__next > main > section > div > div > section > div.css-1vdlxxv > div.table-footer > div.css-17b63ve > button").MustClick()
	//fmt.Println("got click")
	el := page.MustElement("#__next > main > section > div > div > section > div.css-1vdlxxv > div.css-1usd5i5 > div > div.key").MustText()
	fmt.Println(el)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
