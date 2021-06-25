package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1Size := len(nums1)
    n2Size := len(nums2)
    
    midIndex := (n1Size + n2Size) / 2
    isEven := ((n1Size + n2Size) & 0x1 == 0)
    var index, midBefore, midVal, in1, in2 int
        
    for ; index <= midIndex && in1 < n1Size && in2 < n2Size; index++ {
        if nums1[in1] < nums2[in2] {
            midBefore = midVal
            midVal = nums1[in1]
            in1 += 1
        } else {
            midBefore = midVal
            midVal = nums2[in2]
            in2 += 1
        }
    }
    
    for ; index <= midIndex && in1 < n1Size; index++ {
        midBefore = midVal
        midVal = nums1[in1]
        in1 += 1
    }

    for ; index <= midIndex && in2 < n2Size; index++ {
        midBefore = midVal
        midVal = nums2[in2]
        in2 += 1
    }

    if isEven {
        return float64(midBefore + midVal) / float64(2.0)
    } else {
        return float64(midVal)
    }   
}

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	n2 := []int{0, 6}
	fmt.Printf("N1: %+v, N2: %+v\n Got: %v\n", n1, n2, findMedianSortedArrays(n1, n2))
}