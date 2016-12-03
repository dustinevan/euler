package main

import "fmt"

func main() {

	numberOfSolutions := 0
	result := 0

	for i := 1; i <= 1000; i++ {
		numOfSolutions := getTriangles(i)
		if numOfSolutions > numberOfSolutions {
			numberOfSolutions = numOfSolutions
			result = i
		}
	}
	fmt.Println(result)
}

func getTriangles(p int) int {
	result := 0
	for a := 1; a < p/2; a++ {
		for b := 1; b <= a && b < p/2; b++ {
			if (a*a + b*b) == (p-a-b)*(p-a-b) {
				result++
				fmt.Println(a, b, p-a-b, p, result)
			}
		}
	}
	return result
}