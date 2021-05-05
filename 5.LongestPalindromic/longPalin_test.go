package longpalin

import "testing"

func TestPalindromeSubstring(t *testing.T) {
	in := map[string]string{
		"babad": "bab",
		"cbbd":  "bb",
		"a":     "a",
		"ac":    "a",
	}

	for s, want := range in {
		got := longestPalindrome(s)

		if got != want {
			t.Fatalf("Failed for '%s':\n \t\t Got: '%s' \t Want: '%s'\n", s, got, want)
		}

		got = longestPalindrome2(s)
		if got != want {
			t.Fatalf("Failed for '%s' in longestPalindrome2:\n \t\t Got: '%s' \t Want: '%s'\n", s, got, want)
		}

	}
}
