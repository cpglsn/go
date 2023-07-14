package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// start := time.Now()

	for _, site := range sites {
		wg.Add(1)

		go checkLink(site)

	}

	wg.Wait()

	// fmt.Printf("took %v\n", time.Since(start))
}

func checkLink(link string) {

	defer wg.Done()

	for {
		_, err := http.Get(link)
		if err != nil {
			fmt.Println("Error: ", link, " not working")
		}

		fmt.Println(link, " is up and responsing")

		time.Sleep(time.Second * 5)
	}

}
