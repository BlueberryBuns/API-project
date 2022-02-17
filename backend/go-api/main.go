package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Testing golang application with docker 🔥")
	})

	fmt.Printf("Started working!\n")
	http.ListenAndServe(":8080", nil)
	fmt.Print("Stopped working!\n")
}
