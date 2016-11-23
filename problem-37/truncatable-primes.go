package main

import (
	//"strings"
	"strconv"
	"github.com/dustinevan/euler/common/operations"
	"fmt"
)

func main() {
	numbers := []string{"1", "3", "7", "9"}
	results := make([]int, 0)

	current := map[string]int{ "3" : 0, "7" : 0 }


	for {
		next := make(map[string]int)
		for k, _ := range current {
			for _, n := range numbers {
				newNum, _ := strconv.Atoi(k+n)
				fmt.Print("Checking ", newNum)
				if operations.IsPrime(newNum) {
					fmt.Print(" +Prime")
					next[k+n] = 1
					if isReverseTruncatablePrime(k+n) {
						fmt.Println(" !!!FOUND:", newNum, "!!!")
						results = append(results, newNum)
					}
				} else {
					fmt.Println(" X Not Prime")
				}

			}
 		}
		fmt.Println(next)
		fmt.Println(results)
		if len(next) == 0 {
			break
		} else {
			current = next
		}
	}
	fmt.Println(results)
}


func isReverseTruncatablePrime(num string) bool {
	length := len(num)
	for i := 0; i < length-1; i++ {
		x, _ := strconv.Atoi( num[length-1-i:] )
		if !operations.IsPrime(x) {
			fmt.Println(" X Not Reverse Truncatable Prime becuase", x, "is not prime")
			return false
		}
	}
	fmt.Print(" +Reverse Truncatable Prime")
	return true
}