package main

import "testing"

func TestFindMountainPeak(t *testing.T) {
	in := []int{1, 2, 0}
	want := 1
	got := findMountainPeak(in)
	if got != want {
		t.Fatalf("Failed\n\t\t Got: %v\t Want: %v\n", got, want)
	}
}