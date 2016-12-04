package main

import (
	"github.com/dustinevan/euler/common/operations"
	"github.com/dustinevan/euler/common/primes"
	"fmt"
)

func main() {

	result := 2
	for i := []int{1,2,3}; len(i) < 10; i = append(i, i[len(i)-1]+1) {
		fmt.Println(i)
		pandigital := operations.Permutations(i)
		for _, x := range pandigital {
			if primes.IsPrime(x) && x > result {
				fmt.Println(x)
				result = x
			}
		}
	}

}
