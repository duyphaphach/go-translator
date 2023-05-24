package main

import "log"
import "net/http"

func main() {
	var constant = "Hello World"

	http.HandleFunc("/translate", translateHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
