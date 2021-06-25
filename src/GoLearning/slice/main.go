package main

import (
	"fmt"
	"runtime"
)

func main() {
	manyBigSlices := make([][]int, 10)
	for i := 0; i < len(manyBigSlices); i++ {
		bigSlice := make([]int, 1e6)
		// Fill with data 1...1e6
		for j := 0; j < len(bigSlice); j++ {
			bigSlice[j] = j + 1
		}
		// Pop the last element
		bigSlice = bigSlice[:len(bigSlice)-1]
		// The length of bigSlice is 1e6 - 1, buts its capacity remains at 1e6
		// Also note that all 1e6 elements can still be accessed by reslicing

		// Now cut of the begining of the slice
		bigSlice = bigSlice[(len(bigSlice) - 1):]
		// The length of bigSlice is 1, but its capacity is 2
		// Note: There is now no way to reslice and access the first 999,998 elements
		manyBigSlices[i] = bigSlice
	}
	// Notice that the GC did not garbage collect the slices
	// even though we only have access to 1 element total
	runtime.GC()
	PrintMemUsage()
	fmt.Printf("manyBigSlices: %+v, len: %v, cap: %v\n", manyBigSlices, len(manyBigSlices), cap(manyBigSlices))
	// Add one more element to the slice (capacity will be full)
	for i := 0; i < len(manyBigSlices); i++ {
		manyBigSlices[i] = append(manyBigSlices[i], 123)
	}
	runtime.GC()
	PrintMemUsage()
	fmt.Printf("manyBigSlices: %+v, len: %v, cap: %v\n", manyBigSlices, len(manyBigSlices), cap(manyBigSlices))
	// Add one more element to the slice (each slice will be reallocated)
	for i := 0; i < len(manyBigSlices); i++ {
		manyBigSlices[i] = append(manyBigSlices[i], 456)
	}
	runtime.GC()
	// Now we notice that memory was cleaned up
	PrintMemUsage()
	fmt.Printf("manyBigSlices: %+v, len: %v, cap: %v\n", manyBigSlices, len(manyBigSlices), cap(manyBigSlices))
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc (cumulative) = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
