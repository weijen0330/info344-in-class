package main

import (
	"fmt"
	"os"
	"time"
)

const usage = `
usage:
	crawler <starting-url>
`

func worker(linkq chan string, resultsq chan []string) {
	// channel can be looped
	for link := range linkq {
		plinks, err := getPageLinks(link)
		// if we get a bad link, report it and go to the next item in the channel
		if err != nil {
			fmt.Printf("Error fetching %s: %v", link, err)
			continue
		}

		fmt.Printf("%s (%d links)\n", link, len(plinks.Links))
		time.Sleep(time.Millisecond * 500)

		// put all the links we get back into the link queue
		if len(plinks.Links) > 0 {

			// creating a go routine whose only job is write to the result's
			// queue and unblock the main's goroutine
			go func(links []string) {
				resultsq <- links
			}(plinks.Links)

		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	nWorkers := 10
	linkq := make(chan string, 100)
	resultsq := make(chan []string)
	for i := 0; i < nWorkers; i++ {
		// spin up 10 workers
		go worker(linkq, resultsq)
	}

	// set starting URL
	linkq <- os.Args[1]

	// tracking if the URL is seen
	seen := map[string]bool{}

	for links := range resultsq {

		// - = index of element, which we don't care about
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				linkq <- link
			}
		}
	}
}
