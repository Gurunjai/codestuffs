package intques

import (
	"fmt"
	"time"
	"testing"
)

func TestTriggerTask(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond)
	validater := time.NewTicker(time.Second)
	defer ticker.Stop()
	defer validater.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(30 * time.Second)
		done <- true
	} ()

	i := 1
	lloop: for {
		select {
		case <-done:
			break lloop
		case <-ticker.C:
			triggerTask()			
		case <-validater.C:
			fmt.Printf("Trigger Tasked (Sample-%v): %v\n", i, triggerTask())
			i++
		default:
			triggerTask()
		}
	}
}