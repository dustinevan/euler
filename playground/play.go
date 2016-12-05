package main

import (
	"fmt"
	//"github.com/dustinevan/euler/common/operations"
	"strconv"
)


func main() {
	s := "1234567890"
	i, err := strconv.Atoi(s[1:4])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}


