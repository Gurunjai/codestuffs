package intques

import (
	// "sync"
	"time"
)

const MAXSIZE = 10

var counter [MAXSIZE]uint64
var index int

func triggerTask() uint64 {
	counter[index]++

	var sum uint64
	for _, v := range(counter) {
		sum += v
	}

	return sum
}

func init() {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		done := make(chan bool)

		go func() {
			time.Sleep(30 * time.Second)
			done <- true
		}()

		lloop: for {
			select {
			case <-done:
				break lloop
			case <-ticker.C:
				index = ((index + 1) % MAXSIZE)
				counter[index] = 0
			}
		}
	}()
}