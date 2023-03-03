package main

import (
	"sort"

	"golang.org/x/exp/rand"
)

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

func partition(data *[]int, i int, k int, compare func(key1, key2 *int) int) int {
	// Use the median-of-three method to find the partition value.
	r := make([]int, 3)
	r[0] = (rand.Int() % (k - i + 1)) + i
	r[1] = (rand.Int() % (k - i + 1)) + i
	r[2] = (rand.Int() % (k - i + 1)) + i
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	pval := (*data)[r[1]]

	// Create two partitions around the partition value.
	for {
		// Move left until an element is found in the wrong partition.
		for compare(&(*data)[k], &pval) > 0 {
			k--
		}
		// Move right until an element is found in the wrong partition.
		for compare(&(*data)[i], &pval) < 0 {
			i++
		}

		if i >= k {
			// Stop partitioning when the left and right counters cross.
			break
		} else {
			// Swap the elements now under the left and right counters.
			(*data)[i], (*data)[k] = (*data)[k], (*data)[i]
		}
	}
	return k
}

func qkSort(data *[]int, i int, k int, compare func(key1, key2 *int) int) int {
	// Stop the recursion when it is not possible to partition further.
	for i < k {
		// Determine where to partition the elements.
		j := partition(data, i, k, compare)
		// Recursively sort the left partition.
		qkSort(data, i, j, compare)
		i = j + 1
	}
	return 0
}
