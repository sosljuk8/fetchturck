package main

import (
	"fetchturck/parse"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"log/slog"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link == "" {
			return
		}
		//slog.Info("LINK FOUND", slog.String("url", link))

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("Page visited: ", r.Request.URL)
		slog.Info("LINK VISITED", slog.String("url", r.Request.URL.String()))

		err := parse.OnPage(r)
		if err != nil {
			fmt.Println("OnResponse failed", err.Error())
			return
		}
	})

	// downloading the target HTML page
	//c.Visit("https://www.hoodoo.digital/blog/how-to-make-a-web-crawler-using-go-and-colly-tutorial")

	//urls, err := parse.GetUrls()
	//if err != nil {
	//	fmt.Println(err)
	//}

	// Add URLs to the queue
	//for _, url := range urls {
	//	q.AddURL(url)
	//}

	if parse.LoadableURLs(c) {
		fmt.Println("URLs loaded")
	}

	// Consume URLs
	//q.Run(c)

	//u, err := store.ParseURLs()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, url := range u {
	//	fmt.Println(url)
	//}

	//os.Stdout.Write(output)

}
