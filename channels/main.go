package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
		"http://hackersoon.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// }

	//infinity loop
	// for {
	// 	go checkLink(<-ch, ch)
	// }
	//alternet code for upper loop

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	ch <- link
}
