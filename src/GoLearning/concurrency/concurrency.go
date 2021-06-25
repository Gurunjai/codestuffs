package main

import (
	"log"
	"sync"
	"time"
)

var cMap sync.Map

func sBR(url string, ch chan<- uint64) {
	go func(url string, ch chan<- uint64) {
	lpBr:
		for {
			select {
			case <-time.After(100 * time.Millisecond):
				ch <- uint64(4969832)
				break lpBr
			}
		}
	}(url, ch)
}

func gIR(url string, done chan<- bool) {
	go func(url string, final chan<- bool) {
		for {
			select {
			case <-time.After(2 * time.Second):
				final <- true
				break
			}
		}
	}(url, done)
}

func processLReq(url string, wg *sync.WaitGroup) {
	go func(url string, wg *sync.WaitGroup) {
		ch := make(chan uint64)
		go func(url string, ch chan uint64) {
			br, ok := cMap.Load(url)
			if ok {
				log.Printf("Into the lookup of map")
				ch <- br.(uint64)
				return
			} else {
				sBR(url, ch)
			}

			go func(url string) {
				done := make(chan bool)
				go gIR(url, done)
			lpRes:
				for {
					select {
					case <-done:
						log.Printf("Done with gIR API\n")
						wg.Done()
						break lpRes
					}
				}
			}(url)
		}(url, ch)

		go func(ch <-chan uint64) {
			start := time.Now()
			rate := <-ch
			cMap.Store(url, rate)
			log.Printf("received %v in %0.2f seconds", rate, time.Since(start).Seconds())

			log.Printf("Respond\n")
		}(ch)
	}(url, wg)
}
