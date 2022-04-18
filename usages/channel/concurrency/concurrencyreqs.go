package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

// main Before, our program had to request pages one by one at a time. Goroutines let us start processing the next request,
// while we’re waiting for a website to respond.
// The program completes in as little as one-third of the time!
func main() {
	pages := make(chan Page) // Make a channel of int values.
	urls := []string{"https://example.com", "https://golang.org/", "https://google.com"}

	for _, url := range urls {
		go responseSize(url, pages) // Pass the channel with each call to responseSize()
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages // There will be three sends on the channel, so do three receives.
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println(err1)
		}
	}(response.Body) // Release the network connection once the main func exits.

	body, err := ioutil.ReadAll(response.Body) // Read all the data in the response.
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body)) // Convert the data to a string and print it.
	channel <- Page{URL: url, Size: len(body)}
}
