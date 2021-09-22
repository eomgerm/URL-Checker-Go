package main

import (
	"fmt"
	"net/http"
	"errors"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
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
	
	for _, url := range urls {
		hitUrl(url)
	}
}

func hitUrl (url string) error {
	fmt.Println("Checking: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}