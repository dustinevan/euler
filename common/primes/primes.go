package primes


var numList []bool
var primeLimst []int
var topPrime = 2
var nonPrimeFactorList [][][]int
var limit = 10000000
var topLimit = 10000000000

func init() {
	numList = make([]bool, limit)
	/*
	Non-Primes are true in the numList. This is because this algo finds
	primes by sifting out all the non primes, and the default bool in golang is false
	 */
	numList[0] = true
	numList[1] = true
	nonPrimeFactorList = make([][][]int, 1)
	nonPrimeFactorList[0] = make([][]int, 1)
	primeList := make([]int, 1)
	primeList[0] = 2
}

func IsPrime(x int) bool {
	if x < topPrime {
		return !numList[x]
	}
	if x > limit && x < topLimit{
		increaseLimit(x)
	}
	findPrimes(x)
	return !numList[x]
}

func increaseLimit(x int) {
	if x + 10000 < topLimit {
		numList = append(numList, make([]bool, x - limit + 10000)...)
		limit = x + 10000
	} else {
		numList = append(numList, make([]bool, topLimit - limit)...)
		limit = topLimit
	}
}

func findPrimes(x int) {
	calcNonPrimes(x)
}

func calcNonPrimes(x int) {
	topLevel := topLevelNeeded(x)
	for len(nonPrimeFactorList)-1 < topLevel {
		nonPrimeFactorList[len(nonPrimeFactorList)-1] = make([][]int, 1)
	}

	for x > topPrime {
		for level := 2; level < len(nonPrimeFactorList)-1; level++ {
			advanceMultiples(x, level)
		}
		collectPrimes()
	}
}

func advanceMulitples()

func topLevelNeeded(x int) int {
	power_of_two := 4
	count := 2
	for x > power_of_two {
		power_of_two *= 2
		count++
	}
	return count
}

func advanceMultiples(x, level int) {

}