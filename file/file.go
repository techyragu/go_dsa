package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// b, err := os.ReadFile("a.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(b))

	fd, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, fd)

}
