package main

/* #include <time.h> */
import "C"

import (
	"fmt"
	"log"
	"runtime"
	"syscall"
	"time"
)

const MAX_COUNT = 10

var startTime = time.Now()
var startTicks = C.clock()

func getCpuUsage() string {
	clockSec := float64(C.clock()-startTicks) / float64(c.CLOCKS_PER_SEC)
	goSecs := time.Since(startTime).Seconds()

	return fmt.Sprintf("CPU Usage: %.2f", clockSec/goSecs*100)
}

func bToMb(b uint64) uint32 {
	return uint32(b >> 20)
}

func getMemStats() string {
	var mStats runtime.MemStats
	runtime.ReadMemStats(&mStats)

	return fmt.Sprintf("Memory Stats: Total Alloc: %d MB, Sys: %d MB, NumGC: %d\n", bToMb(mStats.TotalAlloc), bToMb(mStats.Sys), mStats.NumGC)
}

func main() {
	var cnt int
	start := make(chan bool)
	done := make(chan bool)

	go func() {
		var b bool
		for {
			if !b {
				b = true
				start <- b
			}
		}
	}()

	go func() {
		<-start
	ll:
		for {
			select {
			case <-time.After(1 * time.Second):
				cnt++
				var ru syscall.Rusage
				if err := syscall.Getrusage(syscall.RUSAGE_SELF, &ru); err != nil {
					log.Fatalf("Error-%v\n", err)
				}

				log.Printf("\n\t\t\t\tSample - %d\n", cnt)

				// Dump Rusage
				log.Printf("Rusage: %+v\n", ru)

				// CPU Usage
				log.Println(getCpuUsage())

				// Mem Usage
				log.Println(getMemStats())

				if MAX_COUNT <= cnt {
					break ll
				}
			}
		}

		done <- true
	}()

	fmt.Println(fmt.Sprintf("\n\t\t\t\t ***** Done ***** "))
}
