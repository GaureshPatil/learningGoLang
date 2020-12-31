package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

// to format it in string 
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	// underscore if you dont intend to use the variable and to ignore the error
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
	
}