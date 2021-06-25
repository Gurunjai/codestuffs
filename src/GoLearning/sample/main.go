package main

import (
	"fmt"
	"time"
)

func main() {
	var t *time.Timer
	t = time.AfterFunc(1*time.Second, func() {
		fmt.Println("what!!!!")
	})
ll:
	for {
		select {
		case <-time.After(10 * time.Second):
			if t.Stop() {
				<-t.C
			}
			break ll
		}
	}
}
