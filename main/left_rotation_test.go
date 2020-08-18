package main

import (
	"reflect"
	"testing"
)

func TestLeftRotation(t *testing.T) {
	scenarios := []struct {
		a []int32
		d int32
		expected []int32
	}{
		{
			a:        []int32{1, 2, 3, 4, 5},
			d:        4,
			expected: []int32{5, 1, 2, 3, 4},
		}, {
			a:        []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
			d:        10,
			expected: []int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
		}, {
			a:        []int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97},
			d:        13,
			expected: []int32{87, 97, 33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60},
		},
	}

	for _, scenario := range scenarios {
		actual := rotLeft(scenario.a, scenario.d)
		expected := scenario.expected

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("TestSockMerchant expected %v actual %v", expected, actual)
			t.FailNow()
		}
	}
}

func rotLeft(a []int32, d int32) []int32 {
	result := make([]int32, len(a))
	diff := int(d)

	for i := 0; i < len(a); i++{
		newIndex := (i + (len(a) - diff)) % len(a)
		result[newIndex] = a[i]
	}
	return result;
}
