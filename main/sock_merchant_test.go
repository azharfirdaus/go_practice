package main

import "testing"

/**
	reference: https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
 */

func TestSockMerchant(t *testing.T) {
	scenarios := []struct {
		n int32
		ar []int32
		expected int32
	}{
		{
			n:        9,
			ar:       []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			expected: 3,
		},
	}

	for _, scenario := range scenarios {
		actual := sockMerchant(scenario.n, scenario.ar)
		expected := scenario.expected

		if actual != expected {
			t.Errorf("TestSockMerchant expected %d actual %d", expected, actual)
			t.FailNow()
		}
	}
}

func sockMerchant(n int32, ar []int32) int32 {

	record := map[int32] *int{}
	var result int32 = 0;

	for i := 0; i < len(ar); i++{
		if record[ar[i]] == nil{
			value := 1;
			record[ar[i]] = &value;
		} else {
			temp := *record[ar[i]];
			temp++;
			record[ar[i]] = &temp

			if temp % 2 == 0 {
				result++
			}
		}
	}

	return result;
}
