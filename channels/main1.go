package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
