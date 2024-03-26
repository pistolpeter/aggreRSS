package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeedData(url string) (RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return RSS{}, err
	}

	// var reader []byte
	var data RSS

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(url)
		return RSS{}, errors.New("bad decode")
	}

	return data, nil
}
