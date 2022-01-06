package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var url = []string{
	"https://www.google.com/",
	"https://www.airasia.co.in/home",
	"https://raft.github.io/",
	"https://en.wikipedia.org/wiki/Merkle_tree",
}

func main() {
	fmt.Println("Start Main Thread")
	openWebsite()
	fmt.Println("End Main Thread")
}

func openWebsite() {
	http.HandleFunc("/", sendRequest)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func sendRequest(wr http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	fmt.Println("Open URL")
	for _, url := range url {
		wg.Add(1)
		go openUrl(url, &wg)
	}
	wg.Wait()
	fmt.Println("Done opening. Receiving response")
	fmt.Println("All URL response received!")
	fmt.Fprintf(wr, "Res")

}

func openUrl(url string, wg *sync.WaitGroup) {
	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Error:", error)
	}
	wg.Done()
	fmt.Println(response.Status, "", url)
}
