package main

import "testing"

/**
	ref: https://www.hackerrank.com/challenges/repeated-string/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
 */

func TestRepeatedStrings(t *testing.T) {
	scenarios := []struct {
		s string
		n int64
		expected int64
	}{
		{
			s:        "aba",
			n:        10,
			expected: 7,
		}, {
			s:        "a",
			n:        1000000000000,
			expected: 1000000000000,
		},
	}

	for _, scenario := range scenarios {
		actual := repeatedString(scenario.s, scenario.n)
		expected := scenario.expected

		if actual != expected {
			t.Errorf("TestSockMerchant expected %d actual %d", expected, actual)
			t.FailNow()
		}
	}
}

func repeatedString(s string, n int64) int64 {
	length := int64(len(s))

	numberOccurance := n/length
	offside := n - (length * numberOccurance)

	var numberOfA int64 = 0;
	var numberOfNotOffsideA int64 = 0;
	var i int64 = 0;
	for ; i < int64(len(s)); i++ {
		if s[i] == 'a' {
			numberOfA++;
		}

		if s[i] == 'a' && int64(i) < offside {
			numberOfNotOffsideA++;
		}
	}

	result := numberOccurance * int64(numberOfA) + numberOfNotOffsideA;
	return result;
}
