package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	result := make(map[int]string)

	for x := 1; x < 100; x++ {
		for y := 10; y < 10000; y++ {
			pandigital := strconv.Itoa(x) + strconv.Itoa(y) + strconv.Itoa(x*y)
			if isPandigital(pandigital) {
				result[x*y] = strconv.Itoa(x) + ":" + strconv.Itoa(y)
			}
		}
	}

	//fmt.Println(result)
	sum := 0
	for k, _ := range result {
		sum += k
	}
	fmt.Println(sum)
}

func isPandigital(s string) bool {
	//fmt.Println(s)
	if strings.Contains(s, "0") {
		return false
	}
	for i := 1; i < 10; i++ {
		if strings.Count(s, strconv.Itoa(i)) != 1 {
			return false
		}
	}
	return true
}
