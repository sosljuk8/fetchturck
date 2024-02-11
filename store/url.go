package store

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
)

type URLSet struct {
	XMLName string `xml:"urlset"`

	URLs []URL `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

// ParseURLs from file turck_sitemap.xml
func ParseURLs() ([]string, error) {
	var urls []string
	urlset := URLSet{}
	file, err := ioutil.ReadFile("store/turck_sitemap.xml")
	if err != nil {
		log.Fatal(err)
	}
	data := bytes.NewBufferString(string(file))

	// convert []byte to io.Reader

	if err = xml.NewDecoder(data).Decode(&urlset); err != nil {
		log.Fatalln(err)
	}
	for _, url := range urlset.URLs {
		urls = append(urls, url.Loc)
		//fmt.Println(url.Loc)
	}
	return urls, nil
}

//client := http.DefaultClient
//
//req, err := http.NewRequest("GET", "https://www.turck.us/en", nil)
//if err != nil {
//	log.Fatalf("error sending HTTP request: %v", err)
//}
//// ...
////req.Header.Add("If-None-Match", `W/"wyzzy"`)
//resp, err := client.Do(req)
//if err != nil {
//	log.Fatalf("error sending HTTP request: %v", err)
//}
//responseBytes, err := ioutil.ReadAll(resp.Body)
//if err != nil {
//	log.Fatalf("error reading HTTP response body: %v", err)
//}
//
//doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(string(responseBytes)))
//if err != nil {
//	log.Fatal(err)
//}
//
//var links []string
//
//doc.Find("ul.navLevel2p li a").Each(func(i int, s *goquery.Selection) {
//	cat, _ := s.Attr("title")
//	s.Next().Find("ul.navLevel3p li a").Each(func(i int, s *goquery.Selection) {
//		subcat, _ := s.Attr("title")
//		//if emty cat or subcat, continue
//		if cat == "" || cat == "Downloads" || subcat == "" {
//			return
//		}
//		//link := cat + "/" + subcat
//		link := "https://www.turck.us/en/productgroup/" + cat + "/" + subcat + "/pws?iwp[]=pwsPager-O-10000&test=test"
//
//		links = append(links, link)
//	})
//})
//
//for _, link := range links {
//	log.Println(link)
//}
