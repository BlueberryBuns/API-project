package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// log.Println("Testing golang application with docker 🔥")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "An arror occured", http.StatusBadRequest)
			return
			//Alternative version
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("An error occured"))
			// return
		}

		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "I have read data %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World :3")
	})

	fmt.Printf("Started working!\n")
	http.ListenAndServe(":8080", nil)
	fmt.Print("Stopped working!\n")
}
