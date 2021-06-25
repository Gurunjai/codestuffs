package main

import (
	"sync"
	"testing"
)

func TestProcessLReq(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(10)
	processLReq("abcistheurl", &wg)
	processLReq("abcistheurl", &wg)
	processLReq("abcistheurl", &wg)
	processLReq("abcistheurl", &wg)
	processLReq("abcistheurl", &wg)

	processLReq("xyzisnottheurl", &wg)
	processLReq("xyzisnottheurl", &wg)
	processLReq("xyzisnottheurl", &wg)
	processLReq("xyzisnottheurl", &wg)
	processLReq("xyzisnottheurl", &wg)
	wg.Wait()
}
