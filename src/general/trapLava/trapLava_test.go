package trapLava

import "testing"

type Formatter struct {
	nums []int
	want int
}

func TestTrapLava(t *testing.T) {
	testVal := make(map[int]Formatter)

	testVal[1] = Formatter {
		nums: []int{0, 1, 2, 1, 2, 0, 3},
		want: 3,
	}

	testVal[2] = Formatter {
		nums: []int{3, 0, 1, 2, 1, 2, 0, 1, 3},
		want: 14,
	}

	testVal[3] = Formatter {
		nums: []int{3, 0, 1, 2, 1, 2, 0, 1},
		want: 5,
	}

	testVal[4] = Formatter {
		nums: []int{5, 0, 1, 2, 1, 2, 0},
		want: 4,
	}

	for key, vals := range testVal {
		if got := trapLava(vals.nums); got != vals.want {
			t.Errorf("Failed for Key: %v\n, \t\t Got: %v, \t Want: %v\n", key, got, vals.want)
		}
	}

}