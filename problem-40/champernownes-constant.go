package main

import (
	"strconv"
	"fmt"
)

func main() {

	var s string
	i := 1
	for len(s) <= 1000000 {
		s += strconv.Itoa(i)
		i++
	}
	j := 1
	result := 1
	for j <= 1000000 {
		x := string(s[j-1])
		if y, err := strconv.Atoi(x); err != nil {
			break;
		} else {
			fmt.Println(j, y)
			result *= y
		}
		j *= 10
	}
	fmt.Println(result)
}
