package main

import "testing"

/**
	ref: https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup
 */

func TestJumpingOnCloud(t *testing.T) {
	scenarios := []struct {
		c []int32
		expected int32
	}{
		{
			c:       []int32{0, 0, 1, 0, 0, 1, 0},
			expected: 4,
		}, {
			c:        []int32{0, 0, 0, 0, 1, 0},
			expected: 3,
		},
	}

	for _, scenario := range scenarios {
		actual := jumpingOnClouds(scenario.c)
		expected := scenario.expected

		if actual != expected {
			t.Errorf("TestSockMerchant expected %d actual %d", expected, actual)
			t.FailNow()
		}
	}
}

func jumpingOnClouds(c []int32) int32 {
	index := 0;
	var numberJump int32 = 0;
	for index < len(c) - 1 {
		nextJump := index + 2;
		if nextJump >= len(c) {
			nextJump = index + 1;
		}

		if c[nextJump] == 0{
			numberJump++;
			index = nextJump;
		} else if c[nextJump - 1] == 0{
			numberJump++;
			index = nextJump - 1
		} else {
			break;
		}
	}

	return numberJump;
}