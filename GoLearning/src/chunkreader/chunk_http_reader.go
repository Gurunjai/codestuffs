package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func readFromUpstream(w http.ResponseWriter, r *http.Request) {
	url := "http://gopl.io"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error in schedulig http GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	st := time.Now()
	n, err := io.Copy(w, resp.Body)
	if err != nil {
		log.Fatalf("Error in reading from response: %v\n", err)
	}
	log.Printf("Wrote %v bytes in %.2f seconds\n", n, time.Since(st).Seconds())
}

func main() {
	http.HandleFunc("/", readFromUpstream)

	http.ListenAndServe(":8080", nil)
}
