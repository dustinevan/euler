package main

import (
	"sort"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	results := make([]int, 0)
	for i := 99999; i > 0; i-- {
		p, flag := hasPandigital(i)
		if flag == true {
			results = append(results, p)
		}
	}
	sort.Ints(results)
	fmt.Println(results)
}

func hasPandigital(x int) (int, bool) {
	var productString string
	i := 1
	for len(productString) < 9 {
		//fmt.Println(x, "X", i)
		productString += strconv.Itoa(i*x)
		i++
	}

	if len(productString) == 9 && isPandigital(productString) {
		result, err := strconv.Atoi(productString)
		if err != nil {
			fmt.Println("Huh?")
			return 0, false
		}
		return result, true
	}
	return 0, false
}

func isPandigital(s string) bool{
	if strings.Contains(s, "0") {
		return false
	}
	if !strings.Contains(s, "1") {
		return false
	}
	if !strings.Contains(s, "2") {
		return false
	}
	if !strings.Contains(s, "3") {
		return false
	}
	if !strings.Contains(s, "4") {
		return false
	}
	if !strings.Contains(s, "5") {
		return false
	}
	if !strings.Contains(s, "6") {
		return false
	}
	if !strings.Contains(s, "7") {
		return false
	}
	if !strings.Contains(s, "8") {
		return false
	}
	if !strings.Contains(s, "9") {
		return false
	}
	return true
}