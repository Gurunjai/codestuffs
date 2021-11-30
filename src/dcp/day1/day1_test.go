package day1

import "testing"

func TestTwoSum(t *testing.T) {
	in := []int{10, 5, 6, 8}
	target, want := 17, true
	if got := TwoSum(in, target); got != want {
		t.Fatalf("Failed\n \t\t Got: %v, \t\t Want: %v\n", got, want)
	}
}