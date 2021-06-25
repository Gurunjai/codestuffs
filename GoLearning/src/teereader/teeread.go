package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

func main() {
	printMemUsage()
	res, err := http.Get("https://golang.org/doc/")
	if err != nil {
		log.Fatalf("failed to retrieve the data: %v\n", err)
	}
	defer res.Body.Close()

	printMemUsage()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("failed to read the response body: %v\n", err)
	}
	fmt.Fprintf(ioutil.Discard, string(buf))

	printMemUsage()

	runtime.GC()
	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	log.Printf("Alloc = %v MB\n", bToMb(m.Alloc))
	log.Printf("\tTotalAlloc = %v MB", bToMb(m.TotalAlloc))
	log.Printf("\tSys = %v MB", bToMb(m.Sys))
	log.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b >> 20
}
