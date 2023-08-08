package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	websiteList := []string{
		"https://golang.org",
		"https://nodejs.org",
		"https://google.com",
		"https://github.com",
		"https://youtube.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
