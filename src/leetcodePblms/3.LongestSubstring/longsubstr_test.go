package longsubstr

import "testing"

func TestLengthOfLongSubstr(t *testing.T) {
	in := map[string]int {
		"abcabcbb" : 3,
		"bbbbb" : 1,
		"pwwkew" : 3,
		"" : 0,
	}

	for s, want := range in {
		got := lengthOfLongestSubstring(s)

		if got != want {
			t.Fatalf("Failed for string '%s'\n\t\t Got: %v \t Want: %v\n", s, got, want)
		}
	}
}