package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func GetUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print("Failed to fetch url")
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal("Failed to read from body")
	}
	//fmt.Println(resp.Body)
}

func main() {
	GetUrl("https://www.google.com/")
}
