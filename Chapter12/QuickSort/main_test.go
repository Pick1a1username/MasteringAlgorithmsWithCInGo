package main

import (
	"reflect"
	"testing"
)

func Test_test(t *testing.T) {
	testCases := []struct {
		description string
		arr         []int
		answer      []int
	}{
		{
			description: "TestCase 1",
			arr:         []int{24, 52, 11, 94, 28, 36, 14, 80},
			answer:      []int{11, 14, 24, 28, 36, 52, 80, 94},
		},
		// {
		// 	description: "TestCase 2",
		// 	arr:         []int{10, 5, 3, 6, 4, 9, 2, 1, 8, 7},
		// 	answer:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// },
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result := append([]int{}, testCase.arr...)
			qkSort(&result, 0, len(result)-1, compareInt)
			if !reflect.DeepEqual(result, testCase.answer) {
				t.Errorf("result: %d answer: %d", result, testCase.answer)
			}
		})
	}
}
