package main

import (
	"fmt"
)


func main() {
	nums := [4]int8{1,2,3,4}
	comb := make([][]int, 0)
	for x := 0; x < 3; x++ {

		for y := 0; y < len(nums); y++ {

			i := x*len(nums)+y
			fmt.Println(i)
			if len(comb) < i+1 {
				comb = append(comb, make([]int8, 0))
			}
			comb[i] = append(comb[i], nums[y])


		}
	}
	fmt.Println(comb, nums)


}


