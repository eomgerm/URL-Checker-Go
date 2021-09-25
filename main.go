package main

import (
	"fmt"
	"net/http"
	"errors"
)

var errRequestFailed = errors.New("Request Failed")

type result struct {
	url string
	status string
}

func main() {
	results := make(map[string]string)
	urls := []string{
	"https://www.airbnb.com/",
	"https://www.google.com/",
	"https://www.amazon.com/",
	"https://www.reddit.com/",
	"https://www.google.com/",
	"https://soundcloud.com/",
	"https://www.facebook.com/",
	"https://www.instagram.com/",
	"https://academy.nomadcoders.co/",
	}
	c := make(chan result)
	
	for _, url := range urls {
		go hitUrl(url, c)
	}
	
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	
	for url, status := range(results) {
		fmt.Println("URL: ", url, " STATUS: ", status)
	}
}

func hitUrl (url string, c chan<- result) {
	status := "OK"
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}