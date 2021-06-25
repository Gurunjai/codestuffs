package main

import "testing"

func TestMedian(t *testing.T) {
	n1 := []int {1, 3}
	n2 := []int {2}
	want := float64(2)
	got := findMedianSortedArrays(n1, n2)
	if got != want {
		t.Fatalf("Failed\n got: %v\t want: %v\n", got, want)
	}
}