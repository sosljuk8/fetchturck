package parse

import (
	"bytes"
	"fetchturck/store"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"log"
	"log/slog"
	"strings"
)

func OnLink(*colly.HTMLElement) (bool, error) {
	return true, nil

}

func OnPage(r *colly.Response) error {

	url := r.Request.URL.String()

	// if HTML contains div with class "product-page" then this is the page
	if !bytes.Contains(r.Body, []byte(`class="prodDetail clearfix"`)) {
		// just skip this url, no errors triggered
		return nil
	}
	fmt.Println("ON PAGE:", url)

	_, err := store.SaveInFile(string(r.Body), url)
	if err != nil {
		fmt.Println("Error saving in file:", err)
		return err
	}

	return nil
}

func GetUrls() ([]string, error) {
	urls, err := store.ParseURLs()
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func LoadableURLs(c *colly.Collector) bool {

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		10, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 1000000}, // Use default queue storage
	)

	c.OnScraped(func(r *colly.Response) {
		//fmt.Println(string(r.Body), " scraped!")
		loadedURLs, err := ParseLoaded(r)
		if err != nil {
			fmt.Println(err)
		}
		// Add URLs to the queue
		for _, u := range loadedURLs {
			q.AddURL(u)
		}
		slog.Info("ON SCRAP", slog.String("url", r.Request.URL.String()))
	})
	//urls, err := GetUrls()
	//if err != nil {
	//	fmt.Println(err)
	//}

	// Add URLs to the queue
	//for _, url := range urls {
	//	q.AddURL(url)
	//}

	q.AddURL("https://www.turck.us/en/productgroup/Sensors/Condition%20Monitoring%20Sensors/pws?iwp[]=pwsPager-O-10000&test=test")

	q.Run(c)

	c.OnResponse(func(r *colly.Response) {
		url := r.Request.URL.String()

		slog.Info("ON PAGE", slog.String("url", url))
	})

	return false
}

func ParseLoaded(r *colly.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(string(r.Body)))
	if err != nil {
		log.Fatal(err)
	}

	// parse href attribute from each div.pwresult div.pwimg first child a trim space and append to urls
	var urls []string
	doc.Find("div.pwresult div.pwimg a").Each(func(i int, s *goquery.Selection) {
		// "https://www.turck.us"
		href, _ := s.Attr("href")
		urls = append(urls, "https://www.turck.us"+strings.TrimSpace(href))
	})

	return urls, nil
}
