package main

import (
	"fmt"
	"testing"
)

func TestGetMem(t *testing.T) {
	fmt.Printf("Read Mem Stat is: %v\n", getMemStatsFloat())
}

func TestGetNetStats(t *testing.T) {
	fmt.Printf("Read Network Stat is: %v\n", getNetStats())
}