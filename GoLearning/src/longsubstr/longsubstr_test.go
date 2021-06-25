package main

import "testing"

func TestLongSubStr(t *testing.T) {
	iList := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}

	for s, want := range iList {
		got := lengthOfLongestSubstring(s)
		if got != want {
			t.Errorf("Mismatch, Got: %v \t Want: %v\n", got, want)
		}
	}
}
