package main

import (
	"fmt"
	"time"
	"math/rand"
)

func bsearch(X []int, y int) bool {
	// fmt.Println(X)
	mid := (len(X) / 2)
	if mid < 0 || mid >= len(X) {
		// fmt.Println("NOT FOUND!!!!")
		return false
	}

	if X[mid] == y {
		// fmt.Println("Found!!!!!!!!!!!!!!!!")
		return true
	}

	if X[mid] < y {
		bsearch(X[mid+1:], y)
	} else {
		bsearch(X[:mid], y)
	}

	return false
}

func main() {
	min := 10
	max := 90
	in1 := []int{}
	for i := 0; i < 70; i++ {
		in1 = append(in1, rand.Intn(max - min) + min)
	}
	
	// fmt.Printf("in1: %+v\n", in1)
	b := make(chan bool)
	c := make(chan bool)
	timer := time.NewTimer(1 * time.Second)
	go func() {
		c <- true
		lLoop: for {
			if bsearch(in1, rand.Intn(max-min) + min) {
				break lLoop
			}
		}

		b <- true
	} ()

	<- c
	
	cLoop: for {
		select {
		case <- b:
			fmt.Printf("FOUND!!!!!\n")
			break cLoop
		case <-timer.C:
			fmt.Println("NOTFOUND!!!!")
			break cLoop
		}
	}
}