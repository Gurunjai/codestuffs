package multiplesof3

import "testing"

func TestSumMultiplesOf3and5( t *testing.T ) {
	want := 233168
	if got := sumOfMultiplesof3And5(1000); got != want {
		t.Fatalf("Mismatched:\n \t\t Want: %v, \t Got: %v\n", want, got)
	}
}