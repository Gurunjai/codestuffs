package sample

import (
	// "fmt" 
	"math"
)

func GetMod(n int) int {
	return (n - pow2(n))
}


func pow2(n int) int {
	x := n
	for n > 0 {
		x = n
		n = (n & (n - 1))
	}

	return x
}

func Parity(n uint64) int {
	n ^= n >> 32
	n ^= n >> 16
	n ^= n >> 8
	n ^= n >> 4
	n ^= n >> 2
	n ^= n >> 1
	return int(n & 0x1)
}

func IsPow2(n int) bool {
	return ((n & (n-1)) == 0)
}

func GetRemainder(num1, num2 int) int {
	if num1 == 0 || num2 == 0 {
		return -1
	}

	if !IsPow2(num2) {
		return -1
	}

	return (num1 & (num2 - 1))
}

func msbSet(n int) int {
	x := pow2(n)
	return int(math.Log2(float64(x)))
}