package main

import (
	"strconv"
	"fmt"
	"github.com/dustinevan/euler/common/operations"
)

func main() {
	sum := 0
	for i := 0; i < 1000000; i++ {
		s := strconv.Itoa(i)
		if !operations.IsPalindrome(s) {
			continue
		}
		s1 := strconv.FormatInt(int64(i), 2)
		if !operations.IsPalindrome(s1) {
			continue
		}
		fmt.Println(s, s1)
		sum += i
	}

	fmt.Println(sum)
}
