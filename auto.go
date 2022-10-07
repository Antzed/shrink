package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/goombaio/namegenerator"
	"io/ioutil"
	"math/rand"
	"mvdan.cc/xurls/v2"
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
	//body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	domainSlice := []string{"@cevipsa.com", "@cpav3.com", "@cpav3.com", "@steveix.com", "@tenvil.com", "@tgvis.com", "@amozix.com", "@maxric.com"}
	domainSliceLenAdjusted := len(domainSlice) - 1
	randNum := rand.Intn(domainSliceLenAdjusted)
	domain := domainSlice[randNum]
	name := RandStringRunes(10)
	email := name + domain
	return email
}

func APIFetch() string {

	email := registerEmail()
	fmt.Println("registered temperory email: " + email)
	md5 := md5.Sum([]byte(email))
	md5String := hex.EncodeToString(md5[:])
	fmt.Println("the md5 hash of the email is: " + md5String)
	name := getName()
	fmt.Println("The random name generated is: " + name)
	signUp(name, email)
	link := getLink(md5String)
	key := getKey(link)
	fmt.Println("The api key retrived is: " + key)
	return key

}

func signUp(name string, email string) {

	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://tinypng.com/developers")

	page.MustElement("#top > div > div > form > input:nth-child(1)").MustInput(name)
	//!!! need new name every time
	page.MustElement("#top > div > div > form > input[type=email]:nth-child(2)").MustInput(email)

	page.MustElement("#top > div > div > form > input[type=submit]:nth-child(3)").MustClick()

	//fmt.Println("finished")
	time.Sleep(2 * time.Second)
	page.MustWaitLoad().MustScreenshot("a.png")

}

func getKey(link string) string {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage(link)
	//!!! need the original link from the email
	//fmt.Println("got page")

	key := page.MustElement("#__next > main > section > div > div > section > div.css-1vdlxxv > div.css-1usd5i5 > div > div.key").MustText()

	return key
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func getName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()
	name = strings.ReplaceAll(name, "-", " ")
	return name
}

func getLink(md5 string) string {

	url := "https://privatix-temp-mail-v1.p.rapidapi.com/request/mail/id/" + md5 + "/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "792b0e438fmshfb9971aecd4728bp16d235jsn45bfd672c1be")
	req.Header.Add("X-RapidAPI-Host", "privatix-temp-mail-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))
	message := string(body)

	rxStrict := xurls.Strict()
	resultSlice := rxStrict.FindAllString(message, -1)

	substring := "dashboard/api"
	var result string
	for _, s := range resultSlice {
		if strings.Contains(s, substring) {
			result = s
			break
		}
	}
	result = strings.Replace(result, "&amp;", "&", 2)
	fmt.Println("The api dashboard link is: " + result)

	return result

}
