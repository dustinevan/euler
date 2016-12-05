package main

import (
	"fmt"
	"github.com/dustinevan/euler/common/operations"
	"strconv"
)

func main() {
	d := []int{0,1,2,3,4,5,6,7,8,9}
	p := operations.Permutations(d)
	sum := 0
	for _, x := range p {
		if x < 1000000000 {
			continue
		}
		if isSubstringDivisible(x) {
			sum += x
		}
	}
	fmt.Println(sum)
}

func isSubstringDivisible(x int) bool {

	s := strconv.Itoa(x)
	//fmt.Println(s)
	i, err := strconv.Atoi(s[1:4])
	//fmt.Println(i)
	if i % 2 != 0 || err != nil {
		return false
	}

	i, err = strconv.Atoi(s[2:5])
	//fmt.Println(i)
	if i % 3 != 0 || err != nil {
		return false
	}

	i, err = strconv.Atoi(s[3:6])
	//fmt.Println(i)
	if i % 5 != 0 || err != nil {
		return false
	}

	i, err = strconv.Atoi(s[4:7])
	//fmt.Println(i)
	if i % 7 != 0 || err != nil {
		return false
	}

	i, err = strconv.Atoi(s[5:8])
	//fmt.Println(i)
	if i % 11 != 0 || err != nil {
		return false
	}

	i, err = strconv.Atoi(s[6:9])
	//fmt.Println(i)
	if i % 13 != 0 || err != nil {
		return false
	}

	i, err = strconv.Atoi(s[7:10])
	//fmt.Println(i)
	if i % 17 != 0 || err != nil {
		return false
	}
	return true
}