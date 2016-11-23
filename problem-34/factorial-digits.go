package main

import (
	"fmt"
)

var factorialList = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}



func main() {
	//fmt.Println(factoring.CompareIntFractions(315, 3110310 , 843, 8323782 ))

	results := make([]int, 10)
	for i := 3; i < 10000000; i++ {
		sum := factorialSum(i)
		//fmt.Println(i, " ", sum)
		if sum == i {
			fmt.Println(i)
			results = append(results, i)
		}
	}
	//fmt.Println(factorialSum(10401020))
	//fmt.Println(results)
}
func factorialSum(i int) int {
	sum := 0
	x := i
	//fmt.Print(i, " ")
	for {
		if x == 0 {
			break
		}
		fac := factorialList[x%10]
		//fmt.Print(x, ":")

		sum = sum + fac
		x = x/10
		//fmt.Print(sum, " ")

	}

	return sum
}