package main

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := http.DefaultClient

	req, err := http.NewRequest("GET", "https://www.turck.us/en", nil)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	// ...
	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading HTTP response body: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(string(responseBytes)))
	if err != nil {
		log.Fatal(err)
	}

	var links []string

	doc.Find("ul.navLevel2p li a").Each(func(i int, s *goquery.Selection) {
		cat, _ := s.Attr("title")
		s.Next().Find("ul.navLevel3p li a").Each(func(i int, s *goquery.Selection) {
			subcat, _ := s.Attr("title")
			//if emty cat or subcat, continue
			if cat == "" || cat == "Downloads" || subcat == "" {
				return
			}
			//link := cat + "/" + subcat
			link := "https://www.turck.us/en/productgroup/" + cat + "/" + subcat + "/pws?iwp[]=pwsPager-O-10000&test=test"

			links = append(links, link)
		})
	})

	for _, link := range links {
		log.Println(link)
	}
}
