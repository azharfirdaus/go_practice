package main3

import (
	"fmt"
)

func main() {
	var n int32 = 9;
	ar := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	result := sockMerchant(n, ar);
	fmt.Printf("sockMerchant: %d\n", result)

	c := []int32{0, 0, 1, 0, 0, 1, 0}
	result2 := jumpingOnClouds(c)
	fmt.Printf("jumpingOnClouds: %d\n", result2)

	result3 := repeatedString("a", 1000000000000)
	fmt.Printf("repeatedString: %d\n", result3)

	a := []int32{1, 2, 3, 4, 5}
	var d int32 = 4
	result4 := rotLeft(a, d)
	fmt.Print("rofLeft: %v\n", result4)

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

// Complete the repeatedString function below.
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

func rotLeft(a []int32, d int32) []int32 {

	result := make([]int32, len(a))
	diff := int(d)

	for i := 0; i < len(a); i++{
		newIndex := (i + (len(a) - diff)) % len(a)
		result[newIndex] = a[i]
	}
	return result;
}




