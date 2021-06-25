package main

import (
	"log"
	"strconv"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9
	want := []int{0, 1}

	got := twoSum(input, target)

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("Mistmatch at index: %v\n\t Got: %v\n\t Want: %v\n", i, got[i], want[i])
		}
	}
}

func TestSubStr(t *testing.T) {
	s := "POBJACB24F44E2D15EEBE0838881B8C2CACB7A6A9D88AD8786B7ACB6ABB7ACB6ABB7ACB6ABB7ACAAB9B75349544B"
	rate, err := strconv.ParseUint(s[44:52], 16, 32)
	if err != nil {
		t.Errorf("error converting: %v\n", err)
	}

	rate = rate ^ 0xACB6ABB7
	log.Printf("Rate parsed: %v\n", rate)
}
