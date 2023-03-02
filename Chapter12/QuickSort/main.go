package main

import (
	"golang.org/x/exp/rand"
)

func main() {

}

// Compare two integers (used during median-of-three partitioning).
func compareInt(int1, int2 *int) int {
	if *int1 > *int2 {
		return 1
	} else if *int1 < *int2 {
		return -1
	} else {
		return 0
	}
}

func partition(data *[]int, i int, k int, compare func(key1, key2 int) int) {

	// Allocate storage for the partition value and swapping.
	pval := make([]int, len(*data))
	temp := make([]int, len(*data))

	// Use the median-of-three method to find the partition value.
	r := make([]int, 3)
	r[0] = (rand.Int() % (k - i + 1)) + i
	r[1] = (rand.Int() % (k - i + 1)) + i
	r[2] = (rand.Int() % (k - i + 1)) + i
}
