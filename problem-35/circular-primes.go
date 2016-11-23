package main

import (
	"fmt"
	"github.com/dustinevan/euler/common/operations"
	"regexp"
	"strconv"
	"sort"
)

//var factorialList = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

var nonPrimeNumMatches = regexp.MustCompile(`[0468]`)

func main() {
	results := make([]int, 0)
	sum := 0
	primes := operations.GetAllPrimesBelow(1000000)
	m := make(map[int]int)
	fmt.Println("primes done")
	for _, x := range primes {
		if !nonPrimeNumMatches.MatchString(strconv.Itoa(x)) {
			m[x] = 0
		}
	}
	for k, _ := range m {
		p := operations.GetUniqueIntRotations(k)
		fmt.Println("Checking:", k, p)
		if ContainsAll(m, p) {
			fmt.Println(k, "is a circular prime.")
			results = append(results, k)
			sum += k
		}
	}
	sort.Ints(results)
	fmt.Println(results, "Number of circular Primes:", len(results))
}

func ContainsAll(m map[int]int, p []int) bool {
	for _, p := range p {
		if _, exists := m[p]; exists == false {
			fmt.Println(p, "is not prime.")
			return false
		}
	}
	return true
}

func Sum(a []int) int {
	sum := 0
	for _, x := range a {
		sum += x
	}
	return sum
}
