package main

import (
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
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

	//bodyClean := strings.Trim(string(body), "[")
	//fmt.Println(bodyClean)
	//bodyClean = strings.Trim(bodyClean, "]")
	//fmt.Println(bodyClean)
	//bodyClean = strings.Trim(bodyClean, "\"")
	//fmt.Println(bodyClean)

	//bodySlice := strings.Split(string(body), "\",\"")

	domainSlice := []string{"@cevipsa.com", "@cpav3.com", "@cpav3.com", "@steveix.com", "@tenvil.com", "@tgvis.com", "@amozix.com", "@maxric.com"}
	bodySliceLenAdjusted := len(domainSlice) - 1
	randNum := rand.Intn(bodySliceLenAdjusted)
	domain := domainSlice[randNum]
	name := RandStringRunes(10)
	email := name + domain
	fmt.Println(email)
	return email

	/*
		// Launch a new browser with default options, and connect to it.
		//browser := rod.New().MustConnect()

		// Even you forget to close, rod will close it after main process ends.
		//defer browser.MustClose()

		// Create a new page
		//page := browser.MustPage("https://10minemail.com/en/")

		// We use css selector to get the search input element and input "git"
		//page.MustElement("#tenmin > div.section-top-qr > div > div > div.col-xs-12.col-sm-12.col-md-12.col-lg-12.col-xl-6 > div.temp-emailbox > form > div.input-box-col.hidden-xs-sm > button").MustClick()
		//text, _ := clipboard.ReadAll()
		//return text */
}

func main() {
	message := "[{\"_id\":{\"oid\":\"633f20b6f3bfe100218a233c\"},\"createdAt\":{\"milliseconds\":1665081526316},\"mail_id\":\"b81bd4ee4040573ee0f410ed93add2b3\",\"mail_address_id\":\"6a0ec9c3ee32466de9c196ed06d817b6\",\"mail_from\":\"Tinify <s\nupport@tinify.com>\",\"mail_subject\":\"Get your Tinify API key\",\"mail_preview\":\"...\",\"mail_text_only\":\"<!DOCTYPE html><html lang=\\\"en\\\"><head><meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-\n8\\\"><meta http-equiv=\\\"viewport\\\" content=\\\"width=device-width, initial-scale=1\\\"><meta http-equiv=\\\"X-UA-Compatible\\\" content=\\\"IE=edge\\\"></head><body style=\\\"-ms-text-size-adjust: 100%; -webkit-text-size-\nadjust: 100%; background: #f5f9fa; color: #40444F; font-family: Arial, sans-serif; font-size: 15px; line-height: 22px; margin: 0; min-width: 100%; padding: 0; width: 100% !important\\\" bgcolor=\\\"#f5f9fa\\\"><s\ntyle type=\\\"text/css\\\">body { width: 100% !important; min-width: 100%; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; margin: 0; padding: 0; }\\n.ExternalClass { width: 100%; }\\n.ExternalClass {\n line-height: 100%; }\\n#backgroundTable { margin: 0; padding: 0; width: 100% !important; line-height: 100% !important; }\\nimg { outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; display\n: block; }\\nbody { margin: 0; padding: 0; color: #40444F; background-color: #f5f9fa; font-size: 15px; font-family: Arial, sans-serif; line-height: 22px; -webkit-text-size-adjust: 100%; }\\nimg { display: blo\nck; margin: 0; padding: 0; }\\n@media only screen and (max-width: 620px) {\\n  html body table td { width: 260px !important; }\\n  html body table td.column { width: 0 !important; }\\n  html body table.content \n{ width: 320px !important; }\\n  html body td.header { width: 320px !important; height: 154px; background: url(https://tinypng.com/images/mail/header-api-mobile.png) no-repeat bottom left; background-size: 1\n00%; }\\n  html body td.header img { display: none !important; }\\n  html body td.header.cdn { height: 183px; }\\n  html body td.content { padding: 25px 30px 15px 30px; }\\n  html body td.icon-cell { display: n\none !important; }\\n}\\n</style><table class=\\\"wrapper\\\" width=\\\"100%\\\" style=\\\"background: #f5f9fa; border-collapse: collapse; border-spacing: 0; width: 100%\\\" bgcolor=\\\"#f5f9fa\\\"><tr style=\\\"border-collapse\n: collapse; border-spacing: 0\\\"><td class=\\\"wrapper\\\" align=\\\"center\\\" width=\\\"100%\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; width: 100%\\\" bgcolor=\\\"#f5f9fa\\\">\n<table class=\\\"content\\\" cellpadding=\\\"0\\\" cellspacing=\\\"0\\\" style=\\\"border-collapse: collapse; border-spacing: 0; border: 0; margin-left: auto; margin-right: auto; overflow-wrap: break-word; table-layout: \nfixed; width: 540px; word-break: break-word; word-wrap: break-word\\\"><tbody><tr style=\\\"border-collapse: collapse; border-spacing: 0\\\"><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-c\nollapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td><td class=\\\"spacer\\\" height=\\\"30\\\" width=\\\"540\\\" style=\\\"border-collapse: collapse !important; border-spacin\nr email address on the website.</div></td><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\">\n</td></tr></tbody></table></td></tr></table></body></html>\",\"mail_timestamp\":1665081526.312,\"mail_attachments_count\":0,\"mail_attachments\":{\"attachment\":[]}}]\n"
	rxStrict := xurls.Strict()
	resultSlice := rxStrict.FindAllString(message, -1)
	fmt.Println(resultSlice)
	substring := "dashboard"
	var result string
	for _, s := range resultSlice {
		if strings.Contains(s, substring) {
			result = s
			break
		}
	}
	fmt.Println(result)
	result = strings.Replace(result, "&amp;", "&", 2)
	fmt.Println(result)
	/*
		fmt.Println(result)
		email := registerEmail()
		fmt.Println(email)
		md5 := md5.Sum([]byte(email))
		md5String := hex.EncodeToString(md5[:])
		fmt.Println("the md5 hash is: " + md5String)
		name := getName()
		fmt.Println(name)
		signUp(name, email)
		//time.Sleep(10 * time.Second)
		getLink(md5String)
		//getKey()
		//string := RandStringRunes(20)
		//fmt.Println(string)
	*/
}

func signUp(name string, email string) {

	// Launch a new browser with default options, and connect to it.
	browser := rod.New().MustConnect()

	// Even you forget to close, rod will close it after main process ends.
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://tinypng.com/developers")

	// We use css selector to get the search input element and input "git"
	page.MustElement("#top > div > div > form > input:nth-child(1)").MustInput(name)
	//!!! need new name every time
	page.MustElement("#top > div > div > form > input[type=email]:nth-child(2)").MustInput(email)

	page.MustElement("#top > div > div > form > input[type=submit]:nth-child(3)").MustClick()

	//if page.MustElement("#top > div > div > form > section").MustContainsElement(page.MustElement("#top > div > div > form > section")) {
	//	fmt.Println("need to switch ip")
	//} else {
	//
	//}

	fmt.Println("finished")
	time.Sleep(2 * time.Second)
	page.MustWaitLoad().MustScreenshot("a.png")

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

func getLink(md5 string) {

	url := "https://privatix-temp-mail-v1.p.rapidapi.com/request/mail/id/" + md5 + "/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "792b0e438fmshfb9971aecd4728bp16d235jsn45bfd672c1be")
	req.Header.Add("X-RapidAPI-Host", "privatix-temp-mail-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))
	message := string(body)

	//message := "[{\"_id\":{\"oid\":\"633f1cae4e82d0003a14836b\"},\"createdAt\":{\"milliseconds\":1665080494945},\"mail_id\":\"36c06cb922a5e2a2bf485a6c9ce883b0\",\"mail_address_id\":\"d8fade24d723a87a66af6866354aa5cc\",\"mail_from\":\"Tinify <s\nupport@tinify.com>\",\"mail_subject\":\"Get your Tinify API key\",\"mail_preview\":\"...\",\"mail_text_only\":\"<!DOCTYPE html><html lang=\\\"en\\\"><head><meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-\n8\\\"><meta http-equiv=\\\"viewport\\\" content=\\\"width=device-width, initial-scale=1\\\"><meta http-equiv=\\\"X-UA-Compatible\\\" content=\\\"IE=edge\\\"></head><body style=\\\"-ms-text-size-adjust: 100%; -webkit-text-size-\nadjust: 100%; background: #f5f9fa; color: #40444F; font-family: Arial, sans-serif; font-size: 15px; line-height: 22px; margin: 0; min-width: 100%; padding: 0; width: 100% !important\\\" bgcolor=\\\"#f5f9fa\\\"><s\ntyle type=\\\"text/css\\\">body { width: 100% !important; min-width: 100%; -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; margin: 0; padding: 0; }\\n.ExternalClass { width: 100%; }\\n.ExternalClass {\n line-height: 100%; }\\n#backgroundTable { margin: 0; padding: 0; width: 100% !important; line-height: 100% !important; }\\nimg { outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; display\n: block; }\\nbody { margin: 0; padding: 0; color: #40444F; background-color: #f5f9fa; font-size: 15px; font-family: Arial, sans-serif; line-height: 22px; -webkit-text-size-adjust: 100%; }\\nimg { display: blo\nck; margin: 0; padding: 0; }\\n@media only screen and (max-width: 620px) {\\n  html body table td { width: 260px !important; }\\n  html body table td.column { width: 0 !important; }\\n  html body table.content \n{ width: 320px !important; }\\n  html body td.header { width: 320px !important; height: 154px; background: url(https://tinypng.com/images/mail/header-api-mobile.png) no-repeat bottom left; background-size: 1\n00%; }\\n  html body td.header img { display: none !important; }\\n  html body td.header.cdn { height: 183px; }\\n  html body td.content { padding: 25px 30px 15px 30px; }\\n  html body td.icon-cell { display: n\none !important; }\\n}\\n</style><table class=\\\"wrapper\\\" width=\\\"100%\\\" style=\\\"background: #f5f9fa; border-collapse: collapse; border-spacing: 0; width: 100%\\\" bgcolor=\\\"#f5f9fa\\\"><tr style=\\\"border-collapse\n: collapse; border-spacing: 0\\\"><td class=\\\"wrapper\\\" align=\\\"center\\\" width=\\\"100%\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; width: 100%\\\" bgcolor=\\\"#f5f9fa\\\">\n<table class=\\\"content\\\" cellpadding=\\\"0\\\" cellspacing=\\\"0\\\" style=\\\"border-collapse: collapse; border-spacing: 0; border: 0; margin-left: auto; margin-right: auto; overflow-wrap: break-word; table-layout: \nfixed; width: 540px; word-break: break-word; word-wrap: break-word\\\"><tbody><tr style=\\\"border-collapse: collapse; border-spacing: 0\\\"><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-c\nollapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td><td class=\\\"spacer\\\" height=\\\"30\\\" width=\\\"540\\\" style=\\\"border-collapse: collapse !important; border-spacin\ng: 0; margin: 0; padding: 0\\\"></td><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td></\ntr><tr style=\\\"border-collapse: collapse; border-spacing: 0\\\"><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" \nbgcolor=\\\"#f5f9fa\\\"></td><td class=\\\"header api\\\" width=\\\"540\\\" style=\\\"border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\"><img src=\\\"https://tinypng.com/images/mail/header-api\n.png\\\" style=\\\"-ms-interpolation-mode: bicubic; display: block; margin: 0; outline: none; padding: 0; text-decoration: none\\\" alt=\\\"api-header\\\"></td><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f\n5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td></tr><tr style=\\\"border-collapse: collapse; border-spacing: 0\\\"><td class=\\\"column\\\" width=\\\"\n30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td><td class=\\\"content\\\" width=\\\"460\\\" style=\\\"background: #fff; bord\ner-bottom-left-radius: 3px; border-bottom-right-radius: 3px; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 30px 40px\\\" bgcolor=\\\"#fff\\\"><p style=\\\"-webkit-text-size-adjust: 10\n0%; color: #40444F; font-family: Arial, sans-serif; line-height: 22px; margin: 0 0 15px\\\">Hi weathered water,</p><p style=\\\"-webkit-text-size-adjust: 100%; color: #40444F; font-family: Arial, sans-serif; li\nne-height: 22px; margin: 0 0 15px\\\">Here is the link to your dashboard. Go grab your API key!</p><p class=\\\"button\\\" style=\\\"-webkit-text-size-adjust: 100%; color: #40444F; font-family: Arial, sans-serif; l\nine-height: 22px; margin: 0 0 15px; text-align: center\\\" align=\\\"center\\\"><a href=\\\"https://tinify.com/login?token=QLjbzk4wNkQlmHZHXqPGlcp0Yxc3NsXw/xb54RzUhybGNYpj77d3x8&amp;new=true&amp;redirect=/dashboard\n/api\\\" style=\\\"background: #3abbd0; border-color: #3abbd0; border-radius: 2px; border-style: solid; border-width: 15px 15px 13px 20px; box-shadow: 0px 4px 0px 0px #269eb1; color: white; display: inline-bloc\nk; font-family: Arial, sans-serif; font-size: 16px; font-weight: bold; letter-spacing: 0.03em; line-height: 17px; margin: 10px auto 15px; text-align: center; text-decoration: none; text-shadow: 1px 1px 1px \n#269eb1\\\">Visit your dashboard</a></p><p style=\\\"-webkit-text-size-adjust: 100%; color: #40444F; font-family: Arial, sans-serif; line-height: 22px; margin: 0 0 15px\\\">You can always access your dashboard ag\nain through the <a href=\\\"https://tinify.com/developers\\\" style=\\\"color: #3abbd0; text-decoration: none\\\">developer section</a> on the website. If you have any questions or feedback, do not hesitate to repl\ny to this email or send us an email at<a href=\\\"mailto:support@tinify.com\\\" style=\\\"color: #3abbd0; text-decoration: none\\\"> support@tinify.com</a>.</p><p style=\\\"-webkit-text-size-adjust: 100%; color: #404\n44F; font-family: Arial, sans-serif; line-height: 22px; margin: 0 0 15px\\\">Have a magnificent day!</p><p style=\\\"-webkit-text-size-adjust: 100%; color: #40444F; font-family: Arial, sans-serif; line-height: \n22px; margin: 0 0 15px\\\"><em>Team Tinify, on behalf of George the Panda</em><br></p></td><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: \n0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td></tr><tr style=\\\"border-collapse: collapse; border-spacing: 0\\\"><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !\nimportant; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td><td width=\\\"540\\\" style=\\\"border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\"><div class=\\\"footer\\\n\" style=\\\"color: #B5C5C9; font-family: Arial, sans-serif; font-size: 11px; line-height: 14px; margin: 20px auto 30px; padding: 0 5%; text-align: center\\\" align=\\\"center\\\">This email was sent to <a href=\\\"ma\nilto:9bnjp4u2t4@cpav3.com\\\" style=\\\"color: #B5C5C9;\\\">9bnjp4u2t4@cpav3.com</a>.<br>You received this message because someone entered your email address on the website.</div></td><td class=\\\"column\\\" width=\\\n\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td></tr></tbody></table></td></tr></table></body></html>\",\"mail_text\n\":\"api-header [https://tinypng.com/images/mail/header-api.png]Hi weathered water,\\n\\nHere is the link to your dashboard. Go grab your API key!\\n\\nVisit your dashboard\\n[https://tinify.com/login?token=QLjbzk\n4wNkQlmHZHXqPGlcp0Yxc3NsXw/xb54RzUhybGNYpj77d3x8&new=true&redirect=/dashboard/api]\\n\\nYou can always access your dashboard again through the developer section\\n[https://tinify.com/developers] on the website\n. If you have any questions or\\nfeedback, do not hesitate to reply to this email or send us an email at \\nsupport@tinify.com [support@tinify.com].\\n\\nHave a magnificent day!\\n\\nTeam Tinify, on behalf of Geo\nrge the Panda\\n\\n\\nThis email was sent to 9bnjp4u2t4@cpav3.com [9bnjp4u2t4@cpav3.com].\\nYou received this message because someone entered your email address on the\\nwebsite.\",\"mail_html\":\"<!DOCTYPE html><ht\none\\\"> support@tinify.com</a>.</p><p style=\\\"-webkit-text-size-adjust: 100%; color: #40444F; font-family: Arial, sans-serif; line-height: 22px; margin: 0 0 15px\\\">Have a magnificent day!</p><p style=\\\"-webk\nit-text-size-adjust: 100%; color: #40444F; font-family: Arial, sans-serif; line-height: 22px; margin: 0 0 15px\\\"><em>Team Tinify, on behalf of George the Panda</em><br></p></td><td class=\\\"column\\\" width=\\\"\n30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td></tr><tr style=\\\"border-collapse: collapse; border-spacing: 0\\\"><t\nd class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bgcolor=\\\"#f5f9fa\\\"></td><td width=\\\"540\\\" style=\\\"border-collap\nse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\"><div class=\\\"footer\\\" style=\\\"color: #B5C5C9; font-family: Arial, sans-serif; font-size: 11px; line-height: 14px; margin: 20px auto 30px; \npadding: 0 5%; text-align: center\\\" align=\\\"center\\\">This email was sent to <a href=\\\"mailto:9bnjp4u2t4@cpav3.com\\\" style=\\\"color: #B5C5C9;\\\">9bnjp4u2t4@cpav3.com</a>.<br>You received this message because s\nomeone entered your email address on the website.</div></td><td class=\\\"column\\\" width=\\\"30\\\" style=\\\"background: #f5f9fa; border-collapse: collapse !important; border-spacing: 0; margin: 0; padding: 0\\\" bg\ncolor=\\\"#f5f9fa\\\"></td></tr></tbody></table></td></tr></table></body></html>\",\"mail_timestamp\":1665080494.938,\"mail_attachments_count\":0,\"mail_attachments\":{\"attachment\":[]}}]\n"
	rxStrict := xurls.Strict()
	resultSlice := rxStrict.FindAllString(message, -1)
	substring := "dashboard"
	var result string
	for _, s := range resultSlice {
		if strings.Contains(s, substring) {
			result = s
			break
		}
	}

	fmt.Println(result)
}
