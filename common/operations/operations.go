package operations

import (
	"strconv"
	"fmt"
	"github.com/dustinevan/euler/common/operations/sub"
)

var factorialList = []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
var primeList []int
var primeMap map[int]int

func init() {
	primeList = make([]int, 4)
	primeMap = make(map[int]int)
	primeList[0] = 2
	primeMap[2] = 1
	primeList[1] = 3
	primeMap[3] = 1
	primeList[2] = 5
	primeMap[5] = 1
	primeList[3] = 7
	primeMap[7] = 1
}

func Test() bool {
	return sub.Test()
}

func IsPrime(x int) bool {

	if primeList[len(primeList)-1] < x {
		getMorePrimes(x)
	}
	if primeMap[x] == 1 {
		return true
	} else {
		return false
	}
}


func GetAllPrimesBelow(limit int) []int {
	getMorePrimes(limit)
	return primeList
}

func getMorePrimes(limit int) {
	if primeList[len(primeList)-1] < limit {
		for {
			if primeList[len(primeList)-1] < limit {
				getNextPrime()
			} else {
				break
			}
		}
	}
}

func getNextPrime() {
	next := primeList[len(primeList)-1] + 2

	nextNum:
	for {
		for i := 1; i < len(primeList); i++ {
			if next%primeList[i] == 0 {
				next = next + 2
				continue nextNum
			}
		}
		//fmt.Println(next)
		primeList = append(primeList, next)
		primeMap[next] = 1
		return
	}
}

func GetPrimeIndex(x int) int {

	for i:= 0; i < len(primeList); i++ {
		if primeList[i] == x {
			return i
		}
	}
	return -1
}

func Factor(num int) []int {
	factors := make([]int, 0)
	for {
		for _, x := range primeList {
			//factor out each prime n number of times
			for {
				if num%x == 0 {
					num = num / x
					factors = append(factors, x)
				} else {
					break
				}
			}
			if num == 1 {
				return factors
			} else {
				getNextPrime()
			}
		}
	}
}

func ReduceIntFraction(numerator, denominator int) (num int, den int) {
	numFactors := Factor(numerator)
	denFactors := Factor(denominator)
	loop:
	for {
		for i, x := range numFactors {
			for j, y := range denFactors {
				if x == y {
					numFactors = append(numFactors[:i], numFactors[i+1:]...)
					denFactors = append(denFactors[:j], denFactors[j+1:]...)
					goto loop
				}
			}
		}
		break
	}

	return Product(numFactors), Product(denFactors)

}

func Product(s []int) int {
	x := 1
	for _, i := range s {
		x *= i
	}
	return x
}

func CompareIntFractions(num0, den0, num1, den1 int) bool {
	a, b := ReduceIntFraction(num0, den0)
	c, d := ReduceIntFraction(num1, den1)
	if a == c && b == d {
		return true
	}
	return false
}





func Factorial(i int) int {
	product := 1
	if i < 10 {
		return factorialList[i]
	}
	if i >= 10 {
		for x := 2; x <= i; x++ {
			product *= x
		}
		return product
	}
	return product
}

// takes in a list of digits like 1, 2, 3. returns all permutations of those digits
// 123, 132, 213, 231, 312, 321
func Permutations(digits []int ) []int {

	if len(digits) == 1 {
		p := make([]int, 1)
		p[0] = digits[0]
		return p
	}
	result := make([]int, 0)
	for i := 0; i < len(digits); i++ {
		if i == 0 {
			p := Permutations(digits[1:])
			for _, x := range p {
				x += IntPow10(digits[i], len(digits)-1)
				result = append(result, x)
			}
		} else if i < len(digits)-1 {
			t := make([]int, 0)
			t = append(t, digits[:i]...)
			t = append(t, digits[(i+1):]...)
			p := Permutations(t)
			for _, x := range p {
				x += IntPow10(digits[i], len(digits)-1)
				result = append(result, x)
			}
		} else {
			p := Permutations(digits[:len(digits)-1])
			for _, x := range p {
				x += IntPow10(digits[i], len(digits)-1)
				result = append(result, x)
			}

		}
	}
	return result
}

func Combinations(digits []int, choose int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < choose; i++ {
		for j := 0; j < len(digits); j++ {
			result[i][j] := append(result[i][j], digits)
		}
	}
}

func IntPow10(num, powOfTen int) int {

	result := num
	for i := 0; i < powOfTen; i++ {
		result *= 10
	}
	return result
}


func GetDigits(num int) []int {

	x := num
	result := make([]int, 0)
	for x > 10 {
		result = append(result, (x%10))
		x /= 10
	}
	result = append(result, x)
	return result
}


func RotateInt(x int) int {

	if x < 10 {
		return x
	}

	powers := len(strconv.Itoa(x))-1
	onesDigit := x%10
	return IntPow10(onesDigit, powers) + x/10
}


func GetIntRotations(x int) []int {

	results := make([]int, 0)
	numOfRotations := len(strconv.Itoa(x))
	next := x
	for i := 0; i < numOfRotations; i++ {
		results = append(results, next)
		next = RotateInt(next)
	}
	return results
}

func GetUniqueIntRotations(x int) []int {

	results := make(map[int]int)
	numOfRotations := len(strconv.Itoa(x))
	next := x
	for i := 0; i < numOfRotations; i++ {
		results[next] = 0
		next = RotateInt(next)
	}
	out := make([]int, 0)
	for k, _ := range results {
		out = append(out,k)
	}
	return out
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

func IsPalindrome(s string) bool {

	for i := 0; i < (len(s)/2 + 1); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}