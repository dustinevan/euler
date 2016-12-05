package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func main() {
	filename := "/Users/dustincurrie/code/go/src/github.com/dustinevan/euler/common/resources/p042_words.txt"
	result := 0
	s := Stringify(filename)
	words := strings.Split(s, ",")

	for i, w := range words {
		 words[i] = strings.Replace(w, "\"", "", -1)
	}
	isTriangular(words[1])

	for _, w := range words {
		if isTriangular(w) {
			result++

		}
	}
	fmt.Println(result)
}


func Stringify(filename string) string {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Couldn't open/read file,", err)
		return ""
	}

	return string(data)
}

func isTriangular(s string) bool {
	result := 0
	for _, s := range s {
		result += int(s) - 64
	}

	t := 0
	for i := 1; t <= result; i++ {
		t = i * (i+1) / 2
		if t == result {
			fmt.Println(s, t, result)
			return true
		}
	}
	return false
}
